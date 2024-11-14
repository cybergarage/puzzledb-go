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

// nolint:gocognit
package sql

import (
	"errors"

	"github.com/cybergarage/go-logger/log"
	"github.com/cybergarage/go-sqlparser/sql"
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
)

// Use handles a USE query.
func (service *Service) Use(conn Conn, stmt sql.Use) error {
	stmt.DatabaseName()
	conn.SetDatabase(stmt.DatabaseName())
	return nil
}

// Begin handles a BEGIN query.
func (service *Service) Begin(conn Conn, stmt sql.Begin) error {
	ctx := conn.SpanContext()

	dbName := conn.Database()
	db, err := service.Store().GetDatabase(ctx, dbName)
	if err != nil {
		return err
	}

	// Check if the transaction is already started.

	txn, err := service.GetTransaction(conn, db)
	if err == nil {
		err := service.CancelTransactionWithError(ctx, conn, db, txn, err)
		if err != nil {
			return err
		}
		err = service.RemoveTransaction(conn, db)
		if err != nil {
			return err
		}
	}

	// Start a new transaction.

	txn, err = db.Transact(true)
	if err != nil {
		return err
	}

	txn.SetAutoCommit(false)
	err = service.SetTransaction(conn, db, txn)
	if err != nil {
		return service.CancelTransactionWithError(ctx, conn, db, txn, err)
	}

	return nil
}

// Commit handles a COMMIT query.
func (service *Service) Commit(conn Conn, stmt sql.Commit) error {
	ctx := conn.SpanContext()

	dbName := conn.Database()
	db, err := service.Store().GetDatabase(ctx, dbName)
	if err != nil {
		return err
	}

	// Check if the transaction is already started.

	txn, err := service.GetTransaction(conn, db)
	if err != nil {
		return err
	}

	// Commit the transaction.

	err = service.CommitTransaction(ctx, conn, db, txn)
	if err != nil {
		return err
	}

	return nil
}

// Rollback handles a ROLLBACK query.
func (service *Service) Rollback(conn Conn, stmt sql.Rollback) error {
	ctx := conn.SpanContext()

	dbName := conn.Database()
	db, err := service.Store().GetDatabase(ctx, dbName)
	if err != nil {
		return err
	}

	// Check if the transaction is already started.

	txn, err := service.GetTransaction(conn, db)
	if err != nil {
		return err
	}

	// Cancel the transaction.

	err = service.CancelTransaction(ctx, conn, db, txn)
	if err != nil {
		return err
	}

	return nil
}

// CreateDatabase handles a CREATE DATABASE query.
func (service *Service) CreateDatabase(conn Conn, stmt sql.CreateDatabase) error {
	ctx := conn.SpanContext()

	dbName := stmt.DatabaseName()

	_, err := service.Store().GetDatabase(ctx, dbName)
	if err == nil {
		if stmt.IfNotExists() {
			return nil
		}
		return newErrDatabaseExist(dbName)
	}

	err = service.Store().CreateDatabase(ctx, dbName)
	if err != nil {
		return err
	}

	// Post a event message to the coordinator.

	err = service.PostDatabaseCreateMessage(dbName)
	if err != nil {
		log.Error(err)
	}

	return nil
}

// CreateTable handles a CREATE TABLE query.
func (service *Service) CreateTable(conn Conn, stmt sql.CreateTable) error {
	ctx := conn.SpanContext()

	dbName := conn.Database()

	// Get the collection definition from the schema.

	col, err := NewDocumentSchemaFrom(stmt)
	if err != nil {
		return err
	}

	// Check if the database exists.

	db, err := service.Store().GetDatabase(ctx, dbName)
	if err != nil {
		return err
	}

	// Create a new table.

	txn, err := db.Transact(true)
	if err != nil {
		return err
	}

	tblName := stmt.TableName()
	_, err = txn.GetCollection(ctx, tblName)
	if err == nil {
		if err := txn.Cancel(ctx); err != nil {
			return err
		}
		if stmt.IfNotExists() {
			return nil
		}
		return newErrSchemaExist(stmt.TableName())
	}

	err = txn.CreateCollection(ctx, col)
	if err != nil {
		return service.CancelTransactionWithError(ctx, conn, db, txn, err)
	}

	err = txn.Commit(ctx)
	if err != nil {
		return err
	}

	// Post a event message to the coordinator.

	err = service.PostCollectionCreateMessage(dbName, tblName)
	if err != nil {
		log.Error(err)
	}

	return nil
}

// AlterDatabase handles a ALTER DATABASE query.
func (service *Service) AlterDatabase(conn Conn, stmt sql.AlterDatabase) error { //nolint:staticcheck
	return newErrNotSupported(stmt.String())
}

// AlterTable handles a ALTER TABLE query.
func (service *Service) AlterTable(conn Conn, stmt sql.AlterTable) error {
	ctx := conn.SpanContext()

	dbName := conn.Database()
	tblName := stmt.TableName()

	// Check if the database exists.

	db, err := service.Store().GetDatabase(ctx, dbName)
	if err != nil {
		return err
	}

	// Alter the specified tables.

	txn, err := db.Transact(true)
	if err != nil {
		return err
	}

	schema, err := txn.GetCollection(ctx, tblName)
	if err != nil {
		return service.CancelTransactionWithError(ctx, conn, db, txn, err)
	}

	if column, ok := stmt.AddColumn(); ok {
		elem, err := NewDocumentElementFrom(column)
		if err != nil {
			return service.CancelTransactionWithError(ctx, conn, db, txn, err)
		}
		err = schema.AddElement(elem)
		if err != nil {
			return service.CancelTransactionWithError(ctx, conn, db, txn, err)
		}
	}

	if idx, ok := stmt.AddIndex(); ok {
		var err error
		schema, err = NewAlterAddIndexSchemaWith(schema, idx)
		if err != nil {
			return service.CancelTransactionWithError(ctx, conn, db, txn, err)
		}
	}

	if col, ok := stmt.DropColumn(); ok {
		err := schema.DropElement(col.Name())
		if err != nil {
			return service.CancelTransactionWithError(ctx, conn, db, txn, err)
		}
	}

	if _, ok := stmt.RenameTo(); ok {
		return newErrNotSupported(stmt.String())
	}

	if _, _, ok := stmt.RenameColumns(); ok {
		return newErrNotSupported(stmt.String())
	}

	// Update schema

	err = txn.UpdateCollection(ctx, schema)
	if err != nil {
		return service.CancelTransactionWithError(ctx, conn, db, txn, err)
	}

	err = txn.Commit(ctx)
	if err != nil {
		return err
	}

	// Post a event message to the coordinator.

	err = service.PostCollectionUpdateMessage(dbName, tblName)
	if err != nil {
		log.Error(err)
	}

	return nil
}

// DropDatabase handles a DROP DATABASE query.
func (service *Service) DropDatabase(conn Conn, stmt sql.DropDatabase) error {
	ctx := conn.SpanContext()

	dbName := stmt.DatabaseName()

	// Check if the database exists.

	_, err := service.Store().GetDatabase(ctx, dbName)
	if err != nil {
		if stmt.IfExists() {
			return nil
		}
		return err
	}

	// Drop the specified database.

	err = service.Store().RemoveDatabase(ctx, dbName)
	if err != nil {
		return err
	}

	// Post a event message to the coordinator.

	err = service.PostDatabaseDropMessage(dbName)
	if err != nil {
		log.Error(err)
	}

	return nil
}

// DropIndex handles a DROP INDEX query.
func (service *Service) DropTable(conn Conn, stmt sql.DropTable) error {
	ctx := conn.SpanContext()

	dbName := conn.Database()

	// Check if the database exists.

	db, err := service.Store().GetDatabase(ctx, dbName)
	if err != nil {
		if stmt.IfExists() {
			return nil
		}
		return err
	}

	// Drop the specified tables.

	txn, err := db.Transact(true)
	if err != nil {
		return err
	}

	tables := stmt.Tables()
	for _, table := range tables {
		tblName := table.TableName()
		_, err = txn.GetCollection(ctx, tblName)
		if err != nil {
			if stmt.IfExists() {
				continue
			}
			return service.CancelTransactionWithError(ctx, conn, db, txn, err)
		}
		err = txn.RemoveCollection(ctx, tblName)
		if err != nil {
			return service.CancelTransactionWithError(ctx, conn, db, txn, err)
		}
	}

	err = txn.Commit(ctx)
	if err != nil {
		return err
	}

	// Post a event message to the coordinator.

	for _, table := range tables {
		tblName := table.TableName()
		err := service.PostCollectionDropMessage(dbName, tblName)
		if err != nil {
			log.Error(err)
		}
	}

	return nil
}

// Insert handles a INSERT query.
func (service *Service) Insert(conn Conn, stmt sql.Insert) error {
	ctx := conn.SpanContext()

	// Gets the specified database.

	dbName := conn.Database()
	db, err := service.Store().GetDatabase(ctx, dbName)
	if err != nil {
		return err
	}

	// Starts a new transaction.

	txn, err := service.Transact(conn, db, true)
	if err != nil {
		return err
	}

	// Gets the specified collection.

	col, err := txn.GetCollection(ctx, stmt.TableName())
	if err != nil {
		return service.CancelTransactionWithError(ctx, conn, db, txn, err)
	}

	// Inserts the object using the primary key

	docKey, docObj, err := NewDocumentObjectFromInsert(dbName, col, stmt)
	if err != nil {
		return service.CancelTransactionWithError(ctx, conn, db, txn, err)
	}

	err = txn.InsertDocument(ctx, docKey, docObj)
	if err != nil {
		return service.CancelTransactionWithError(ctx, conn, db, txn, err)
	}

	// Inserts the secondary indexes.

	err = service.InsertSecondaryIndexes(ctx, conn, txn, col, docObj, docKey)
	if err != nil {
		return service.CancelTransactionWithError(ctx, conn, db, txn, err)
	}

	// Commits the transaction if the transaction is auto commit.

	if txn.IsAutoCommit() {
		err := service.CommitTransaction(ctx, conn, db, txn)
		if err != nil {
			return err
		}
	}

	return nil
}

// Select handles a SELECT query.
func (service *Service) Select(conn Conn, stmt sql.Select) (sql.ResultSet, error) {
	ctx := conn.SpanContext()

	// Checks the table if the statement has only one table.

	tables := stmt.From()
	if len(tables) != 1 {
		return nil, newErrJoinQueryNotSupported(tables)
	}

	// Gets the specified database.

	dbName := conn.Database()
	db, err := service.Store().GetDatabase(ctx, dbName)
	if err != nil {
		return nil, err
	}

	// Starts a new transaction.

	txn, err := service.Transact(conn, db, false)
	if err != nil {
		return nil, err
	}

	// Gets the specified collection.

	table := tables[0]
	tableName := table.TableName()
	col, err := txn.GetCollection(ctx, tableName)
	if err != nil {
		return nil, service.CancelTransactionWithError(ctx, conn, db, txn, err)
	}

	// Selects the specified objects.

	rs, err := service.SelectDocumentObjects(ctx, conn, txn, col, stmt.Where(), stmt.OrderBy(), stmt.Limit())
	if err != nil {
		err = errors.Join(err, service.CancelTransactionWithError(ctx, conn, db, txn, err))
	}

	// Commits the transaction if the transaction is auto commit.

	if txn.IsAutoCommit() {
		err = errors.Join(err, service.CommitTransaction(ctx, conn, db, txn))
	}

	// Returs result set.

	if err != nil {
		return nil, err
	}

	return NewResultSetFrom(db, col, rs)
}

// Update handles a UPDATE query.
func (service *Service) Update(conn Conn, stmt sql.Update) (sql.ResultSet, error) {
	ctx := conn.SpanContext()

	// Gets the specified database.

	dbName := conn.Database()
	db, err := service.Store().GetDatabase(ctx, dbName)
	if err != nil {
		return nil, err
	}

	// Starts a new transaction.

	txn, err := service.Transact(conn, db, true)
	if err != nil {
		return nil, err
	}

	// Gets the specified collection.

	tableName := stmt.TableName()
	col, err := txn.GetCollection(ctx, tableName)
	if err != nil {
		return nil, service.CancelTransactionWithError(ctx, conn, db, txn, err)
	}

	// Updates the specified objects.

	updateCols := stmt.Columns()
	rs, err := service.SelectDocumentObjects(ctx, conn, txn, col, stmt.Where(), nil, nil)
	if err != nil {
		return nil, service.CancelTransactionWithError(ctx, conn, db, txn, err)
	}

	nUpdated := 0
	for rs.Next() {
		docObj := rs.Object()
		err := service.UpdateDocument(ctx, conn, txn, col, docObj, updateCols)
		if err != nil {
			return nil, service.CancelTransactionWithError(ctx, conn, db, txn, err)
		}
		nUpdated++
	}

	// Commits the transaction if the transaction is auto commit.

	if txn.IsAutoCommit() {
		err := service.CommitTransaction(ctx, conn, db, txn)
		if err != nil {
			return nil, err
		}
	}

	return sql.NewResultSet(
		sql.WithResultSetRowsAffected(uint64(nUpdated)),
	), nil
}

// Delete handles a DELETE query.
func (service *Service) Delete(conn Conn, stmt sql.Delete) (sql.ResultSet, error) {
	ctx := conn.SpanContext()

	// Gets the specified database.

	dbName := conn.Database()
	db, err := service.Store().GetDatabase(ctx, dbName)
	if err != nil {
		return nil, err
	}

	// Starts a new transaction.

	txn, err := service.Transact(conn, db, true)
	if err != nil {
		return nil, err
	}

	// Gets the specified collection.

	tableName := stmt.TableName()
	col, err := txn.GetCollection(ctx, tableName)
	if err != nil {
		return nil, service.CancelTransactionWithError(ctx, conn, db, txn, err)
	}

	// Deletes the specified objects.

	docKey, docKeyType, err := NewDocumentKeyFromCond(dbName, col, stmt.Where())
	if err != nil {
		return nil, service.CancelTransactionWithError(ctx, conn, db, txn, err)
	}

	nDeleted := 0

	switch docKeyType {
	case document.PrimaryIndex:
		err = service.DeleteDocument(ctx, conn, txn, col, docKey)
		if err != nil {
			if stmt.Where() == nil && errors.Is(err, store.ErrNotExist) {
				nDeleted = 0
			} else {
				return nil, service.CancelTransactionWithError(ctx, conn, db, txn, err)
			}
		} else {
			nDeleted = 1
		}
	case document.SecondaryIndex:
		rs, err := txn.FindDocumentsByIndex(ctx, docKey)
		if err != nil {
			return nil, service.CancelTransactionWithError(ctx, conn, db, txn, err)
		}
		prIdx, err := col.PrimaryIndex()
		if err != nil {
			return nil, service.CancelTransactionWithError(ctx, conn, db, txn, err)
		}
		for rs.Next() {
			docObj := rs.Object()
			obj, err := document.NewMapObjectFrom(docObj)
			if err != nil {
				return nil, err
			}
			objKey, err := NewDocumentKeyFromIndex(dbName, col, prIdx, obj)
			if err != nil {
				return nil, service.CancelTransactionWithError(ctx, conn, db, txn, err)
			}
			err = service.DeleteDocument(ctx, conn, txn, col, objKey)
			if err != nil {
				return nil, service.CancelTransactionWithError(ctx, conn, db, txn, err)
			}
			nDeleted++
		}
	}

	// Commits the transaction if the transaction is auto commit.

	if txn.IsAutoCommit() {
		err := service.CommitTransaction(ctx, conn, db, txn)
		if err != nil {
			return nil, err
		}
	}

	return sql.NewResultSet(
		sql.WithResultSetRowsAffected(uint64(nDeleted)),
	), nil
}

// SystemSelect handles a system SELECT query.
func (service *Service) SystemSelect(conn Conn, stmt sql.Select) (sql.ResultSet, error) {
	return nil, newErrNotSupported(stmt.String())
}

// ParserError handles a parser error.
func (service *Service) ParserError(conn Conn, query string, err error) error {
	return err
}
