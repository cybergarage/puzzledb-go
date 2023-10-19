// Copyright (C) 2022 The PuzzleDB Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mysql

import (
	"errors"

	"github.com/cybergarage/go-logger/log"
	"github.com/cybergarage/go-mysql/mysql"
	"github.com/cybergarage/go-mysql/mysql/query"
	sql "github.com/cybergarage/go-sqlparser/sql/query"
	"github.com/cybergarage/puzzledb-go/puzzledb/context"
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
	sqlc "github.com/cybergarage/puzzledb-go/puzzledb/plugins/query/sql"
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
)

// Begin should handle a BEGIN statement.
func (service *Service) Begin(conn *mysql.Conn, stmt *query.Begin) (*mysql.Result, error) {
	err := service.Service.Begin(conn)
	if err != nil {
		return nil, err
	}
	return mysql.NewResult(), nil
}

// Commit should handle a COMMIT statement.
func (service *Service) Commit(conn *mysql.Conn, stmt *query.Commit) (*mysql.Result, error) {
	err := service.Service.Commit(conn)
	if err != nil {
		return nil, err
	}
	return mysql.NewResult(), nil
}

// Rollback should handle a ROLLBACK statement.
func (service *Service) Rollback(conn *mysql.Conn, stmt *query.Rollback) (*mysql.Result, error) {
	err := service.Service.Rollback(conn)
	if err != nil {
		return nil, err
	}
	return mysql.NewResult(), nil
}

// CreateDatabase should handle a CREATE database statement.
func (service *Service) CreateDatabase(conn *mysql.Conn, stmt *query.Database) (*mysql.Result, error) {
	q := sql.NewCreateDatabaseWith(
		stmt.Name(),
		sql.NewIfNotExistsWith(stmt.IfNotExists()),
	)

	err := service.Service.CreateDatabase(conn, q)
	if err != nil {
		return mysql.NewResult(), err
	}

	return mysql.NewResult(), nil
}

// AlterDatabase should handle a ALTER database statement.
func (service *Service) AlterDatabase(conn *mysql.Conn, stmt *query.Database) (*mysql.Result, error) {
	log.Debugf("%v", stmt)

	dbName := stmt.Name()

	// Post a event message to the coordinator.

	err := service.PostDatabaseUpdateMessage(dbName)
	if err != nil {
		log.Error(err)
	}

	return nil, newQueryNotSupportedError("AlterTable")
}

// DropDatabase should handle a DROP database statement.
func (service *Service) DropDatabase(conn *mysql.Conn, stmt *query.Database) (*mysql.Result, error) {
	q := sql.NewDropDatabaseWith(
		stmt.Name(),
		sql.NewIfExistsWith(stmt.IfExists()),
	)

	err := service.Service.DropDatabase(conn, q)
	if err != nil {
		return mysql.NewResult(), err
	}

	return mysql.NewResult(), nil
}

// CreateTable should handle a CREATE table statement.
func (service *Service) CreateTable(conn *mysql.Conn, stmt *query.Schema) (*mysql.Result, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("CreateTable")
	defer ctx.FinishSpan()

	dbName := conn.Database()

	// Get the collection definition from the schema.

	col, err := NewCollectionWith(stmt)
	if err != nil {
		return nil, err
	}

	// Check if the database exists.

	db, err := service.Store().GetDatabase(ctx, dbName)
	if err != nil {
		return nil, err
	}

	// Create a new table.

	txn, err := db.Transact(true)
	if err != nil {
		return nil, err
	}

	tblName := stmt.TableName()
	_, err = txn.GetCollection(ctx, tblName)
	if err == nil {
		if err := txn.Cancel(ctx); err != nil {
			return nil, err
		}
		if stmt.GetIfNotExists() {
			return mysql.NewResult(), nil
		}
		return mysql.NewResult(), newSchemaExistError(stmt.TableName())
	}

	err = txn.CreateCollection(ctx, col)
	if err != nil {
		return nil, service.CancelTransactionWithError(ctx, txn, err)
	}

	err = txn.Commit(ctx)
	if err != nil {
		return nil, err
	}

	// Post a event message to the coordinator.

	err = service.PostCollectionCreateMessage(dbName, tblName)
	if err != nil {
		log.Error(err)
	}

	return mysql.NewResult(), nil
}

// AlterTable should handle a ALTER table statement.
func (service *Service) AlterTable(conn *mysql.Conn, stmt *query.Schema) (*mysql.Result, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("AlterTable")
	defer ctx.FinishSpan()

	dbName := conn.Database()
	tblName := stmt.TableName()

	// Get the collection definition from the schema.

	_, err := NewAlterCollectionWith(stmt)
	if err != nil {
		return nil, err
	}

	// Check if the database exists.

	db, err := service.Store().GetDatabase(ctx, dbName)
	if err != nil {
		return nil, err
	}

	// Check if the collection exists.

	txn, err := db.Transact(false)
	if err != nil {
		return nil, err
	}

	col, err := txn.GetCollection(ctx, stmt.TableName())
	if err != nil {
		return nil, service.CancelTransactionWithError(ctx, txn, err)
	}

	err = txn.Commit(ctx)
	if err != nil {
		return nil, err
	}

	if col == nil {
		return nil, newSchemaNotExistError(stmt.TableName())
	}

	// Post a event message to the coordinator.

	err = service.PostCollectionCreateMessage(dbName, tblName)
	if err != nil {
		log.Error(err)
	}

	return nil, newQueryNotSupportedError("AlterTable")
}

// DropTable should handle a DROP table statement.
func (service *Service) DropTable(conn *mysql.Conn, stmt *query.Schema) (*mysql.Result, error) {
	tables := []*sql.Table{}
	for _, table := range stmt.GetFromTables() {
		tables = append(tables, sql.NewTableWith(table.Name.String()))
	}
	q := sql.NewDropTableWith(
		sql.NewTablesWith(tables...),
		sql.NewIfExistsWith(stmt.GetIfExists()),
	)
	err := service.Service.DropTable(conn, q)
	if err != nil {
		return mysql.NewResult(), err
	}

	return mysql.NewResult(), nil
}

// RenameTable should handle a RENAME table statement.
func (service *Service) RenameTable(conn *mysql.Conn, stmt *query.Schema) (*mysql.Result, error) {
	log.Debugf("%v", stmt)
	return nil, newQueryNotSupportedError("RenameTable")
}

// TruncateTable should handle a TRUNCATE table statement.
func (service *Service) TruncateTable(conn *mysql.Conn, stmt *query.Schema) (*mysql.Result, error) {
	log.Debugf("%v", stmt)
	return nil, newQueryNotSupportedError("TruncateTable")
}

// Insert should handle a INSERT statement.
func (service *Service) Insert(conn *mysql.Conn, stmt *query.Insert) (*mysql.Result, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("Insert")
	defer ctx.FinishSpan()

	dbName := conn.Database()
	db, err := service.Store().GetDatabase(ctx, dbName)
	if err != nil {
		return nil, err
	}

	txn, err := db.Transact(true)
	if err != nil {
		return nil, err
	}

	col, err := txn.GetCollection(ctx, stmt.TableName())
	if err != nil {
		return nil, service.CancelTransactionWithError(ctx, txn, err)
	}

	// Inserts the object using the primary key/

	docKey, docObj, err := NewObjectFromInsert(dbName, col, stmt)
	if err != nil {
		return nil, service.CancelTransactionWithError(ctx, txn, err)
	}

	err = txn.InsertDocument(ctx, docKey, docObj)
	if err != nil {
		return nil, service.CancelTransactionWithError(ctx, txn, err)
	}

	// Inserts the secondary indexes.

	err = service.insertSecondaryIndexes(ctx, conn, txn, col, docObj, docKey)
	if err != nil {
		return nil, service.CancelTransactionWithError(ctx, txn, err)
	}

	err = txn.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return mysql.NewResultWithRowsAffected(1), nil
}

func (service *Service) insertSecondaryIndexes(ctx context.Context, conn *mysql.Conn, txn store.Transaction, schema document.Schema, obj document.MapObject, prKey document.Key) error {
	idxes, err := schema.SecondaryIndexes()
	if err != nil {
		return err
	}
	for _, idx := range idxes {
		err := service.insertSecondaryIndex(ctx, conn, txn, schema, obj, idx, prKey)
		if err != nil {
			return err
		}
	}
	return nil
}

func (service *Service) insertSecondaryIndex(ctx context.Context, conn *mysql.Conn, txn store.Transaction, schema document.Schema, obj document.MapObject, idx document.Index, prKey document.Key) error {
	dbName := conn.Database()
	secKey, err := sqlc.NewDocumentKeyFromIndex(dbName, schema, idx, obj)
	if err != nil {
		return err
	}
	return txn.InsertIndex(ctx, secKey, prKey)
}

// Select should handle a SELECT statement.
func (service *Service) Select(conn *mysql.Conn, stmt *query.Select) (*mysql.Result, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("Select")
	defer ctx.FinishSpan()

	dbName := conn.Database()
	db, err := service.Store().GetDatabase(ctx, dbName)
	if err != nil {
		return nil, err
	}

	txn, err := db.Transact(false)
	if err != nil {
		return nil, err
	}

	// TODO: Support multiple tables
	tables := stmt.From()
	if len(tables) != 1 {
		return nil, service.CancelTransactionWithError(ctx, txn, newJoinQueryNotSupportedError(tables))
	}

	table := tables[0]
	tableName, err := table.Name()
	if err != nil {
		return nil, service.CancelTransactionWithError(ctx, txn, err)
	}

	col, err := txn.GetCollection(ctx, tableName)
	if err != nil {
		return nil, service.CancelTransactionWithError(ctx, txn, err)
	}

	rs, err := service.selectDocumentObjects(ctx, conn, txn, col, stmt.Where, stmt.OrderBy, stmt.Limit)
	if err != nil {
		return nil, service.CancelTransactionWithError(ctx, txn, err)
	}

	res, err := NewResultFrom(dbName, col, rs.Objects())
	if err != nil {
		return nil, service.CancelTransactionWithError(ctx, txn, err)
	}

	err = txn.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (service *Service) selectDocumentObjects(ctx context.Context, conn *mysql.Conn, txn store.Transaction, schema document.Schema, cond *query.Condition, orderby query.OrderBy, limit *query.Limit) (store.ResultSet, error) {
	docKey, docKeyType, err := NewKeyFromCond(conn.Database(), schema, cond)
	if err != nil {
		return nil, err
	}

	opts := []store.Option{}
	opts = append(opts, NewLimitWith(limit)...)
	opts = append(opts, NewOrderWith(orderby)...)

	switch docKeyType {
	case document.PrimaryIndex:
		return txn.FindDocuments(ctx, docKey, opts...)
	case document.SecondaryIndex:
		return txn.FindDocumentsByIndex(ctx, docKey, opts...)
	}
	return nil, newIndexTypeNotSupportedError(docKeyType)
}

// Update should handle a UPDATE statement.
func (service *Service) Update(conn *mysql.Conn, stmt *query.Update) (*mysql.Result, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("Update")
	defer ctx.FinishSpan()

	dbName := conn.Database()
	db, err := service.Store().GetDatabase(ctx, dbName)
	if err != nil {
		return nil, err
	}

	txn, err := db.Transact(true)
	if err != nil {
		return nil, err
	}

	// TODO: Support multiple tables
	tables := stmt.Tables()
	if len(tables) != 1 {
		return nil, service.CancelTransactionWithError(ctx, txn, newJoinQueryNotSupportedError(tables))
	}

	table := tables[0]
	tableName, err := table.Name()
	if err != nil {
		return nil, service.CancelTransactionWithError(ctx, txn, err)
	}

	col, err := txn.GetCollection(ctx, tableName)
	if err != nil {
		return nil, service.CancelTransactionWithError(ctx, txn, err)
	}

	updateCols, err := stmt.Columns()
	if err != nil {
		return nil, service.CancelTransactionWithError(ctx, txn, err)
	}

	rs, err := service.selectDocumentObjects(ctx, conn, txn, col, stmt.Where, stmt.OrderBy, stmt.Limit)
	if err != nil {
		return nil, service.CancelTransactionWithError(ctx, txn, err)
	}

	for rs.Next() {
		docObj := rs.Object()
		err := service.updateDocument(ctx, conn, txn, col, docObj, updateCols)
		if err != nil {
			return nil, service.CancelTransactionWithError(ctx, txn, err)
		}
	}

	err = txn.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return mysql.NewResult(), nil
}

func (service *Service) updateDocument(ctx context.Context, conn *mysql.Conn, txn store.Transaction, schema document.Schema, obj any, updateCols *query.Columns) error {
	docObj, err := document.NewMapObjectFrom(obj)
	if err != nil {
		return err
	}

	// Removes current secondary indexes
	err = service.RemoveSecondaryIndexes(ctx, conn, txn, schema, docObj)
	if err != nil {
		return err
	}

	// Updates object
	dbName := conn.Database()
	for _, updateCol := range updateCols.Columns() {
		name := updateCol.Name()
		v, err := document.NewValueForSchema(schema, name, updateCol.Value())
		if err != nil {
			return err
		}
		docObj[name] = v
	}
	docKey, err := sqlc.NewDocumentKeyFromObject(dbName, schema, docObj)
	if err != nil {
		return err
	}

	err = txn.UpdateDocument(ctx, docKey, docObj)
	if err != nil {
		return err
	}

	// Inserts new secondary indexes.
	err = service.insertSecondaryIndexes(ctx, conn, txn, schema, docObj, docKey)
	if err != nil {
		return err
	}
	return nil
}

// Delete should handle a DELETE statement.
func (service *Service) Delete(conn *mysql.Conn, stmt *query.Delete) (*mysql.Result, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("Delete")
	defer ctx.FinishSpan()

	dbName := conn.Database()
	db, err := service.Store().GetDatabase(ctx, dbName)
	if err != nil {
		return nil, err
	}

	txn, err := db.Transact(true)
	if err != nil {
		return nil, err
	}

	// TODO: Support multiple tables
	tables := stmt.Tables()
	if len(tables) != 1 {
		return nil, service.CancelTransactionWithError(ctx, txn, newJoinQueryNotSupportedError(tables))
	}

	table := tables[0]
	tableName, err := table.Name()
	if err != nil {
		return nil, service.CancelTransactionWithError(ctx, txn, err)
	}

	col, err := txn.GetCollection(ctx, tableName)
	if err != nil {
		return nil, service.CancelTransactionWithError(ctx, txn, err)
	}

	docKey, docKeyType, err := NewKeyFromCond(dbName, col, stmt.Where)
	if err != nil {
		return nil, service.CancelTransactionWithError(ctx, txn, err)
	}

	switch docKeyType {
	case document.PrimaryIndex:
		err = service.DeleteDocument(ctx, conn, txn, col, docKey)
		if err != nil {
			if stmt.Where != nil || !errors.Is(err, store.ErrNotExist) {
				return nil, service.CancelTransactionWithError(ctx, txn, err)
			}
		}
	case document.SecondaryIndex:
		rs, err := txn.FindDocumentsByIndex(ctx, docKey)
		if err != nil {
			return nil, service.CancelTransactionWithError(ctx, txn, err)
		}
		prIdx, err := col.PrimaryIndex()
		if err != nil {
			return nil, service.CancelTransactionWithError(ctx, txn, err)
		}
		for rs.Next() {
			docObj := rs.Object()
			obj, err := document.NewMapObjectFrom(docObj)
			if err != nil {
				return nil, err
			}
			objKey, err := sqlc.NewDocumentKeyFromIndex(dbName, col, prIdx, obj)
			if err != nil {
				return nil, service.CancelTransactionWithError(ctx, txn, err)
			}
			err = service.DeleteDocument(ctx, conn, txn, col, objKey)
			if err != nil {
				return nil, service.CancelTransactionWithError(ctx, txn, err)
			}
		}
	}

	err = txn.Commit(ctx)
	if err != nil {
		return nil, err
	}
	return mysql.NewResult(), nil
}

// ShowDatabases should handle a SHOW DATABASES statement.
func (service *Service) ShowDatabases(conn *mysql.Conn) (*mysql.Result, error) {
	return nil, newQueryNotSupportedError("ShowDatabases")
}

// ShowTables should handle a SHOW TABLES statement.
func (service *Service) ShowTables(conn *mysql.Conn, database string) (*mysql.Result, error) {
	return nil, newQueryNotSupportedError("ShowTables")
}
