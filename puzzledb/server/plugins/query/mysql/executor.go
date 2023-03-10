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
	"context"

	"github.com/cybergarage/go-logger/log"
	"github.com/cybergarage/go-mysql/mysql"
	"github.com/cybergarage/go-mysql/mysql/query"
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
)

// CreateDatabase should handle a CREATE database statement.
func (service *Service) CreateDatabase(ctx context.Context, conn *mysql.Conn, stmt *query.Database) (*mysql.Result, error) {
	log.Debugf("%v", stmt)
	dbName := stmt.Name()

	store := service.Store()
	_, err := store.GetDatabase(dbName)
	if err == nil {
		if stmt.IfNotExists() {
			return mysql.NewResult(), nil
		}
		return mysql.NewResult(), newDatabaseExistError(dbName)
	}

	err = store.CreateDatabase(dbName)
	if err != nil {
		return nil, err
	}

	return mysql.NewResult(), nil
}

// AlterDatabase should handle a ALTER database statement.
func (service *Service) AlterDatabase(ctx context.Context, conn *mysql.Conn, stmt *query.Database) (*mysql.Result, error) {
	log.Debugf("%v", stmt)
	return nil, newQueryNotSupportedError("AlterTable")
}

// DropDatabase should handle a DROP database statement.
func (service *Service) DropDatabase(ctx context.Context, conn *mysql.Conn, stmt *query.Database) (*mysql.Result, error) {
	dbName := stmt.Name()

	store := service.Store()
	err := store.RemoveDatabase(dbName)
	if err != nil {
		return nil, err
	}

	return mysql.NewResult(), nil
}

// CreateTable should handle a CREATE table statement.
func (service *Service) CreateTable(ctx context.Context, conn *mysql.Conn, stmt *query.Schema) (*mysql.Result, error) {
	log.Debugf("%v", stmt)
	store := service.Store()

	db, err := store.GetDatabase(conn.Database())
	if err != nil {
		return nil, err
	}

	txn, err := db.Transact(true)
	if err != nil {
		return nil, err
	}

	_, err = txn.GetSchema(stmt.TableName())
	if err == nil {
		if err := txn.Cancel(); err != nil {
			return nil, err
		}
		if stmt.GetIfNotExists() {
			return mysql.NewResult(), nil
		}
		return mysql.NewResult(), newSchemaExistError(stmt.TableName())
	}

	schema, err := NewSchemaWith(stmt)
	if err != nil {
		return nil, service.CancelTransactionWithError(txn, err)
	}

	err = txn.CreateSchema(schema)
	if err != nil {
		return nil, service.CancelTransactionWithError(txn, err)
	}

	err = txn.Commit()
	if err != nil {
		return nil, err
	}

	return mysql.NewResult(), nil
}

// AlterTable should handle a ALTER table statement.
func (service *Service) AlterTable(ctx context.Context, conn *mysql.Conn, stmt *query.Schema) (*mysql.Result, error) {
	log.Debugf("%v", stmt)
	return nil, newQueryNotSupportedError("AlterTable")
}

// DropTable should handle a DROP table statement.
func (service *Service) DropTable(ctx context.Context, conn *mysql.Conn, stmt *query.Schema) (*mysql.Result, error) {
	log.Debugf("%v", stmt)
	return nil, newQueryNotSupportedError("DropTable")
}

// RenameTable should handle a RENAME table statement.
func (service *Service) RenameTable(ctx context.Context, conn *mysql.Conn, stmt *query.Schema) (*mysql.Result, error) {
	log.Debugf("%v", stmt)
	return nil, newQueryNotSupportedError("RenameTable")
}

// TruncateTable should handle a TRUNCATE table statement.
func (service *Service) TruncateTable(ctx context.Context, conn *mysql.Conn, stmt *query.Schema) (*mysql.Result, error) {
	log.Debugf("%v", stmt)
	return nil, newQueryNotSupportedError("TruncateTable")
}

// Insert should handle a INSERT statement.
func (service *Service) Insert(ctx context.Context, conn *mysql.Conn, stmt *query.Insert) (*mysql.Result, error) {
	log.Debugf("%v", stmt)
	store := service.Store()

	dbName := conn.Database()
	db, err := store.GetDatabase(dbName)
	if err != nil {
		return nil, err
	}

	txn, err := db.Transact(true)
	if err != nil {
		return nil, err
	}

	schema, err := txn.GetSchema(stmt.TableName())
	if err != nil {
		return nil, service.CancelTransactionWithError(txn, err)
	}

	docKey, doc, err := NewObjectWith(dbName, schema, stmt)
	if err != nil {
		return nil, service.CancelTransactionWithError(txn, err)
	}

	err = txn.InsertDocument(docKey, doc)
	if err != nil {
		return nil, service.CancelTransactionWithError(txn, err)
	}

	err = txn.Commit()
	if err != nil {
		return nil, err
	}

	return mysql.NewResult(), nil
}

// Update should handle a UPDATE statement.
func (service *Service) Update(ctx context.Context, conn *mysql.Conn, stmt *query.Update) (*mysql.Result, error) {
	store := service.Store()

	dbName := conn.Database()
	db, err := store.GetDatabase(dbName)
	if err != nil {
		return nil, err
	}

	txn, err := db.Transact(true)
	if err != nil {
		return nil, err
	}

	err = txn.Commit()
	if err != nil {
		return nil, err
	}

	return mysql.NewResult(), nil
}

// Delete should handle a DELETE statement.
func (service *Service) Delete(ctx context.Context, conn *mysql.Conn, stmt *query.Delete) (*mysql.Result, error) {
	store := service.Store()

	dbName := conn.Database()
	db, err := store.GetDatabase(dbName)
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
		return nil, service.CancelTransactionWithError(txn, newJoinQueryNotSupportedError(tables))
	}

	table := tables[0]
	tableName, err := table.Name()
	if err != nil {
		return nil, service.CancelTransactionWithError(txn, err)
	}

	schema, err := txn.GetSchema(tableName)
	if err != nil {
		return nil, service.CancelTransactionWithError(txn, err)
	}

	docKey, docKeyType, err := NewKeyWithCond(dbName, schema, stmt.Where)
	if err != nil {
		return nil, service.CancelTransactionWithError(txn, err)
	}

	switch docKeyType {
	case document.PrimaryIndex:
		err := txn.RemoveDocument(docKey)
		if err != nil {
			return nil, service.CancelTransactionWithError(txn, err)
		}
	case document.SecondaryIndex:
		prIdx, err := schema.PrimaryIndex()
		if err != nil {
			return nil, service.CancelTransactionWithError(txn, err)
		}
		rs, err := txn.FindDocumentsByIndex(docKey)
		if err != nil {
			return nil, service.CancelTransactionWithError(txn, err)
		}
		for rs.Next() {
			obj := rs.Object()
			docKey, err := NewKeyWithIndex(dbName, schema, prIdx, obj)
			if err != nil {
				return nil, service.CancelTransactionWithError(txn, err)
			}
			err = txn.RemoveDocument(docKey)
			if err != nil {
				return nil, service.CancelTransactionWithError(txn, err)
			}
		}
	}

	err = txn.Commit()
	if err != nil {
		return nil, err
	}
	return mysql.NewResult(), nil
}

// Select should handle a SELECT statement.
func (service *Service) Select(ctx context.Context, conn *mysql.Conn, stmt *query.Select) (*mysql.Result, error) {
	store := service.Store()

	dbName := conn.Database()
	db, err := store.GetDatabase(dbName)
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
		return nil, service.CancelTransactionWithError(txn, newJoinQueryNotSupportedError(tables))
	}

	table := tables[0]
	tableName, err := table.Name()
	if err != nil {
		return nil, service.CancelTransactionWithError(txn, err)
	}

	schema, err := txn.GetSchema(tableName)
	if err != nil {
		return nil, service.CancelTransactionWithError(txn, err)
	}

	docKey, docKeyType, err := NewKeyWithCond(dbName, schema, stmt.Where)
	if err != nil {
		return nil, service.CancelTransactionWithError(txn, err)
	}

	var objs []document.Object
	switch docKeyType {
	case document.PrimaryIndex:
		rs, err := txn.FindDocuments(docKey)
		if err != nil {
			return nil, service.CancelTransactionWithError(txn, err)
		}
		objs = rs.Objects()
	case document.SecondaryIndex:
		rs, err := txn.FindDocumentsByIndex(docKey)
		if err != nil {
			return nil, service.CancelTransactionWithError(txn, err)
		}
		objs = rs.Objects()
	}

	rs, err := NewResultFrom(schema, objs)
	if err != nil {
		return nil, service.CancelTransactionWithError(txn, err)
	}

	err = txn.Commit()
	if err != nil {
		return nil, err
	}

	return rs, nil
}

// ShowDatabases should handle a SHOW DATABASES statement.
func (service *Service) ShowDatabases(ctx context.Context, conn *mysql.Conn) (*mysql.Result, error) {
	return nil, newQueryNotSupportedError("ShowDatabases")
}

// ShowTables should handle a SHOW TABLES statement.
func (service *Service) ShowTables(ctx context.Context, conn *mysql.Conn, database string) (*mysql.Result, error) {
	return nil, newQueryNotSupportedError("ShowTables")
}

func (service *Service) selectDocumentObjects(ctx context.Context, conn *mysql.Conn, txn store.Transaction, tables []*query.Table, cond *query.Condition) (store.ResultSet, error) {
	dbName := conn.Database()

	// TODO: Support multiple tables
	if len(tables) != 1 {
		return nil, service.CancelTransactionWithError(txn, newJoinQueryNotSupportedError(tables))
	}

	table := tables[0]
	tableName, err := table.Name()
	if err != nil {
		return nil, err
	}

	schema, err := txn.GetSchema(tableName)
	if err != nil {
		return nil, err
	}

	docKey, docKeyType, err := NewKeyWithCond(dbName, schema, cond)
	if err != nil {
		return nil, err
	}

	switch docKeyType {
	case document.PrimaryIndex:
		return txn.FindDocuments(docKey)
	case document.SecondaryIndex:
		return txn.FindDocumentsByIndex(docKey)
	}
	return nil, nil
}
