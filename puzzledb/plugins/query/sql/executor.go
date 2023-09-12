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
	"github.com/cybergarage/go-logger/log"
	"github.com/cybergarage/go-sqlparser/sql/query"
	"github.com/cybergarage/puzzledb-go/puzzledb/context"
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
)

// CreateDatabase handles a CREATE DATABASE query.
func (service *Service) CreateDatabase(conn Conn, stmt *query.CreateDatabase) error {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("CreateDatabase")
	defer ctx.FinishSpan()

	dbName := stmt.DatabaseName()

	store := service.Store()
	_, err := store.GetDatabase(ctx, dbName)
	if err == nil {
		if stmt.IfNotExists() {
			return nil
		}
		return newErrDatabaseExist(dbName)
	}

	err = store.CreateDatabase(ctx, dbName)
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

	store := service.Store()
	dbName := conn.Database()

	// Get the collection definition from the schema.

	col, err := NewDocumentSchemaFrom(stmt)
	if err != nil {
		return err
	}

	// Check if the database exists.

	db, err := store.GetDatabase(ctx, dbName)
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

// DropDatabase handles a DROP DATABASE query.
func (service *Service) DropDatabase(conn Conn, stmt *query.DropDatabase) error {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("DropDatabase")
	defer ctx.FinishSpan()

	store := service.Store()
	dbName := stmt.DatabaseName()

	// Check if the database exists.

	_, err := store.GetDatabase(ctx, dbName)
	if err != nil {
		if stmt.IfExists() {
			return nil
		}
		return err
	}

	// Drop the specified database.

	err = store.RemoveDatabase(ctx, dbName)
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

	store := service.Store()
	dbName := conn.Database()

	// Check if the database exists.

	db, err := store.GetDatabase(ctx, dbName)
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

	store := service.Store()

	dbName := conn.Database()
	db, err := store.GetDatabase(ctx, dbName)
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

	store := service.Store()

	dbName := conn.Database()
	db, err := store.GetDatabase(ctx, dbName)
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

	store := service.Store()

	dbName := conn.Database()
	db, err := store.GetDatabase(ctx, dbName)
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

	store := service.Store()

	dbName := conn.Database()
	db, err := store.GetDatabase(ctx, dbName)
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

	docKey, docKeyType, err := NewKeyFromCond(dbName, col, stmt.Where())
	if err != nil {
		return 0, service.CancelTransactionWithError(ctx, txn, err)
	}

	nDeleted := 0

	switch docKeyType {
	case document.PrimaryIndex:
		err = service.DeleteDocument(ctx, conn, txn, col, docKey)
		if err != nil {
			return 0, service.CancelTransactionWithError(ctx, txn, err)
		}
		nDeleted = 1
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
			objKey, err := NewKeyFromIndex(dbName, col, prIdx, obj)
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
