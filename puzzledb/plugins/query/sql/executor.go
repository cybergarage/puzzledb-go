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
	"github.com/cybergarage/go-postgresql/postgresql"
	"github.com/cybergarage/go-postgresql/postgresql/protocol/message"
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

	col, err := NewCollectionWith(stmt)
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
		return service.cancelTransactionWithError(ctx, txn, err)
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

// CreateIndex handles a CREATE INDEX query.
func (service *Service) CreateIndex(conn Conn, stmt *query.CreateIndex) (message.Responses, error) {
	return nil, postgresql.NewErrNotImplemented("CREATE INDEX")
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
func (service *Service) DropTable(conn Conn, stmt *query.DropTable) (message.Responses, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("DropTable")
	defer ctx.FinishSpan()

	store := service.Store()
	dbName := conn.Database()

	// Check if the database exists.

	db, err := store.GetDatabase(ctx, dbName)
	if err != nil {
		if stmt.IfExists() {
			return message.NewCommandCompleteResponsesWith(stmt.String())
		}
		return nil, err
	}

	// Drop the specified tables.

	txn, err := db.Transact(true)
	if err != nil {
		return nil, err
	}

	tables := stmt.Tables()
	for _, table := range tables {
		tblName := table.TableName()
		_, err = txn.GetCollection(ctx, tblName)
		if err != nil {
			if stmt.IfExists() {
				continue
			}
			return nil, service.cancelTransactionWithError(ctx, txn, err)
		}
		err = txn.RemoveCollection(ctx, tblName)
		if err != nil {
			return nil, service.cancelTransactionWithError(ctx, txn, err)
		}
	}

	err = txn.Commit(ctx)
	if err != nil {
		return nil, err
	}

	// Post a event message to the coordinator.

	for _, table := range tables {
		tblName := table.TableName()
		err := service.PostCollectionDropMessage(dbName, tblName)
		if err != nil {
			log.Error(err)
		}
	}

	return message.NewCommandCompleteResponsesWith(stmt.String())
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
		return service.cancelTransactionWithError(ctx, txn, err)
	}

	// Inserts the object using the primary key

	docKey, docObj, err := NewObjectFromInsert(dbName, col, stmt)
	if err != nil {
		return service.cancelTransactionWithError(ctx, txn, err)
	}

	err = txn.InsertDocument(ctx, docKey, docObj)
	if err != nil {
		return service.cancelTransactionWithError(ctx, txn, err)
	}

	// Inserts the secondary indexes.

	err = service.insertSecondaryIndexes(ctx, conn, txn, col, docObj, docKey)
	if err != nil {
		return service.cancelTransactionWithError(ctx, txn, err)
	}

	err = txn.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

// Select handles a SELECT query.
func (service *Service) Select(conn Conn, stmt *query.Select) (context.Context, store.Transaction, document.Collection, store.ResultSet, error) {
	selectDocumentObjects := func(ctx context.Context, conn Conn, txn store.Transaction, schema document.Schema, cond *query.Condition, orderby *query.OrderBy, limit *query.Limit) (store.ResultSet, error) {
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
		return nil, newErrIndexTypeNotSupported(docKeyType)
	}

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
		return ctx, nil, nil, nil, service.cancelTransactionWithError(ctx, txn, err)
	}

	rs, err := selectDocumentObjects(ctx, conn, txn, col, stmt.Where(), stmt.OrderBy(), stmt.Limit())
	if err != nil {
		err = service.cancelTransactionWithError(ctx, txn, err)
	}

	return ctx, txn, col, rs, err
}

// Update handles a UPDATE query.
func (service *Service) Update(conn Conn, stmt *query.Update) (message.Responses, error) {
	return nil, postgresql.NewErrNotImplemented("UPDATE")
}

// Delete handles a DELETE query.
func (service *Service) Delete(conn Conn, stmt *query.Delete) (message.Responses, error) {
	return nil, postgresql.NewErrNotImplemented("DELETE")
}
