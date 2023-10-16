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

package sql

import (
	"errors"

	"github.com/cybergarage/go-logger/log"
	"github.com/cybergarage/go-postgresql/postgresql"
	"github.com/cybergarage/go-sqlparser/sql/query"
	"github.com/cybergarage/puzzledb-go/puzzledb/context"
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
)

// Begin handles a BEGIN query.
func (service *Service) Begin(conn Conn, stmt *query.Begin) error {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("Begin")
	defer ctx.FinishSpan()

	dbName := conn.Database()
	db, err := service.Store().GetDatabase(ctx, dbName)
	if err != nil {
		return err
	}

	// Check if the transaction is already started.

	txn, err := service.GetTransaction(conn, db)
	if err == nil {
		err := service.CancelTransactionWithError(ctx, txn, err)
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

	err = service.SetTransaction(conn, db, txn)
	if err != nil {
		return service.CancelTransactionWithError(ctx, txn, err)
	}

	return nil
}

// Commit handles a COMMIT query.
func (service *Service) Commit(conn Conn, stmt *query.Commit) error {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("Commit")
	defer ctx.FinishSpan()

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

	// Remove the transaction.

	err = service.RemoveTransaction(conn, db)
	if err != nil {
		return err
	}

	// Commit the transaction.

	err = txn.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

// Rollback handles a ROLLBACK query.
func (service *Service) Rollback(conn Conn, stmt *query.Rollback) error {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("Commit")
	defer ctx.FinishSpan()

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

	// Remove the transaction.

	err = service.RemoveTransaction(conn, db)
	if err != nil {
		return err
	}

	// Cancel the transaction.

	err = txn.Cancel(ctx)
	if err != nil {
		return err
	}

	return nil
}

// CreateDatabase handles a CREATE DATABASE query.
func (service *Service) CreateDatabase(conn Conn, stmt *query.CreateDatabase) error {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("CreateDatabase")
	defer ctx.FinishSpan()

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
func (service *Service) CreateTable(conn Conn, stmt *query.CreateTable) error {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("CreateTable")
	defer ctx.FinishSpan()

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
		return service.CancelTransactionWithError(ctx, txn, err)
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
func (service *Service) AlterDatabase(conn *postgresql.Conn, stmt *query.AlterDatabase) error { //nolint:staticcheck
	return newErrNotSupported(stmt.String())
}

// AlterTable handles a ALTER TABLE query.
func (service *Service) AlterTable(conn *postgresql.Conn, stmt *query.AlterTable) error {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("AlterTable")
	defer ctx.FinishSpan()

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
		return service.CancelTransactionWithError(ctx, txn, err)
	}

	if column, ok := stmt.AddColumn(); ok {
		elem, err := NewDocumentElementFrom(column)
		if err != nil {
			return service.CancelTransactionWithError(ctx, txn, err)
		}
		err = schema.AddElement(elem)
		if err != nil {
			return service.CancelTransactionWithError(ctx, txn, err)
		}
	}

	if idx, ok := stmt.AddIndex(); ok {
		var err error
		schema, err = NewAlterAddIndexSchemaWith(schema, idx)
		if err != nil {
			return service.CancelTransactionWithError(ctx, txn, err)
		}
	}

	if col, ok := stmt.DropColumn(); ok {
		err := schema.DropElement(col.Name())
		if err != nil {
			return service.CancelTransactionWithError(ctx, txn, err)
		}
	}

	if _, ok := stmt.RenameTable(); ok {
		return newErrNotSupported(stmt.String())
	}

	if _, _, ok := stmt.RenameColumns(); ok {
		return newErrNotSupported(stmt.String())
	}

	// Update schema

	err = txn.UpdateCollection(ctx, schema)
	if err != nil {
		return service.CancelTransactionWithError(ctx, txn, err)
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
func (service *Service) DropDatabase(conn Conn, stmt *query.DropDatabase) error {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("DropDatabase")
	defer ctx.FinishSpan()

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
func (service *Service) DropTable(conn Conn, stmt *query.DropTable) error {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("DropTable")
	defer ctx.FinishSpan()

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
			return service.CancelTransactionWithError(ctx, txn, err)
		}
		err = txn.RemoveCollection(ctx, tblName)
		if err != nil {
			return service.CancelTransactionWithError(ctx, txn, err)
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
func (service *Service) Insert(conn Conn, stmt *query.Insert) error {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("Insert")
	defer ctx.FinishSpan()

	dbName := conn.Database()
	db, err := service.Store().GetDatabase(ctx, dbName)
	if err != nil {
		return err
	}

	txn, err := db.Transact(true)
	if err != nil {
		return err
	}

	col, err := txn.GetCollection(ctx, stmt.TableName())
	if err != nil {
		return service.CancelTransactionWithError(ctx, txn, err)
	}

	// Inserts the object using the primary key

	docKey, docObj, err := NewDocumentObjectFromInsert(dbName, col, stmt)
	if err != nil {
		return service.CancelTransactionWithError(ctx, txn, err)
	}

	err = txn.InsertDocument(ctx, docKey, docObj)
	if err != nil {
		return service.CancelTransactionWithError(ctx, txn, err)
	}

	// Inserts the secondary indexes.

	err = service.InsertSecondaryIndexes(ctx, conn, txn, col, docObj, docKey)
	if err != nil {
		return service.CancelTransactionWithError(ctx, txn, err)
	}

	err = txn.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

// Select handles a SELECT query.
func (service *Service) Select(conn Conn, stmt *query.Select) (context.Context, store.Transaction, document.Collection, store.ResultSet, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("Select")

	dbName := conn.Database()
	db, err := service.Store().GetDatabase(ctx, dbName)
	if err != nil {
		return ctx, nil, nil, nil, err
	}

	// TODO: Support multiple tables
	tables := stmt.From()
	if len(tables) != 1 {
		return ctx, nil, nil, nil, newErrJoinQueryNotSupported(tables)
	}

	txn, err := db.Transact(false)
	if err != nil {
		return ctx, nil, nil, nil, err
	}

	table := tables[0]
	tableName := table.Name()
	col, err := txn.GetCollection(ctx, tableName)
	if err != nil {
		return ctx, nil, nil, nil, service.CancelTransactionWithError(ctx, txn, err)
	}

	rs, err := service.SelectDocumentObjects(ctx, conn, txn, col, stmt.Where(), stmt.OrderBy(), stmt.Limit())
	if err != nil {
		err = service.CancelTransactionWithError(ctx, txn, err)
	}

	return ctx, txn, col, rs, err
}

// Update handles a UPDATE query.
func (service *Service) Update(conn Conn, stmt *query.Update) (int, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("Update")
	defer ctx.FinishSpan()

	dbName := conn.Database()
	db, err := service.Store().GetDatabase(ctx, dbName)
	if err != nil {
		return 0, err
	}

	txn, err := db.Transact(true)
	if err != nil {
		return 0, err
	}

	tableName := stmt.TableName()
	col, err := txn.GetCollection(ctx, tableName)
	if err != nil {
		return 0, service.CancelTransactionWithError(ctx, txn, err)
	}

	updateCols := stmt.Columns()
	rs, err := service.SelectDocumentObjects(ctx, conn, txn, col, stmt.Where(), nil, nil)
	if err != nil {
		return 0, service.CancelTransactionWithError(ctx, txn, err)
	}

	nUpdated := 0
	for rs.Next() {
		docObj := rs.Object()
		err := service.UpdateDocument(ctx, conn, txn, col, docObj, updateCols)
		if err != nil {
			return 0, service.CancelTransactionWithError(ctx, txn, err)
		}
		nUpdated++
	}

	err = txn.Commit(ctx)
	if err != nil {
		return 0, err
	}

	return nUpdated, nil
}

// Delete handles a DELETE query.
func (service *Service) Delete(conn Conn, stmt *query.Delete) (int, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("Delete")
	defer ctx.FinishSpan()

	dbName := conn.Database()
	db, err := service.Store().GetDatabase(ctx, dbName)
	if err != nil {
		return 0, err
	}

	txn, err := db.Transact(true)
	if err != nil {
		return 0, err
	}

	tableName := stmt.TableName()
	col, err := txn.GetCollection(ctx, tableName)
	if err != nil {
		return 0, service.CancelTransactionWithError(ctx, txn, err)
	}

	docKey, docKeyType, err := NewDocumentKeyFromCond(dbName, col, stmt.Where())
	if err != nil {
		return 0, service.CancelTransactionWithError(ctx, txn, err)
	}

	nDeleted := 0

	switch docKeyType {
	case document.PrimaryIndex:
		err = service.DeleteDocument(ctx, conn, txn, col, docKey)
		if err != nil {
			if stmt.Where() == nil && errors.Is(err, store.ErrNotExist) {
				nDeleted = 0
			} else {
				return 0, service.CancelTransactionWithError(ctx, txn, err)
			}
		} else {
			nDeleted = 1
		}
	case document.SecondaryIndex:
		rs, err := txn.FindDocumentsByIndex(ctx, docKey)
		if err != nil {
			return 0, service.CancelTransactionWithError(ctx, txn, err)
		}
		prIdx, err := col.PrimaryIndex()
		if err != nil {
			return 0, service.CancelTransactionWithError(ctx, txn, err)
		}
		for rs.Next() {
			docObj := rs.Object()
			obj, err := document.NewMapObjectFrom(docObj)
			if err != nil {
				return 0, err
			}
			objKey, err := NewDocumentKeyFromIndex(dbName, col, prIdx, obj)
			if err != nil {
				return 0, service.CancelTransactionWithError(ctx, txn, err)
			}
			err = service.DeleteDocument(ctx, conn, txn, col, objKey)
			if err != nil {
				return 0, service.CancelTransactionWithError(ctx, txn, err)
			}
			nDeleted++
		}
	}

	err = txn.Commit(ctx)
	if err != nil {
		return 0, err
	}

	return nDeleted, nil
}
