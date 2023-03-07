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
		if err := txn.Cancel(); err != nil {
			return nil, err
		}
		return nil, err
	}

	err = txn.CreateSchema(schema)
	if err != nil {
		if err := txn.Cancel(); err != nil {
			return nil, err
		}
		return nil, err
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
	}

	err = txn.Commit()
	if err != nil {
		return nil, err
	}

	return mysql.NewResult(), nil
}

// Update should handle a UPDATE statement.
func (service *Service) Update(ctx context.Context, conn *mysql.Conn, stmt *query.Update) (*mysql.Result, error) {
	return mysql.NewResult(), nil
}

// Delete should handle a DELETE statement.
func (service *Service) Delete(ctx context.Context, conn *mysql.Conn, stmt *query.Delete) (*mysql.Result, error) {
	return mysql.NewResult(), nil
}

// Select should handle a SELECT statement.
func (service *Service) Select(ctx context.Context, conn *mysql.Conn, stmt *query.Select) (*mysql.Result, error) {
	return mysql.NewResult(), nil
}

// ShowDatabases should handle a SHOW DATABASES statement.
func (service *Service) ShowDatabases(ctx context.Context, conn *mysql.Conn) (*mysql.Result, error) {
	return nil, newQueryNotSupportedError("ShowDatabases")
}

// ShowTables should handle a SHOW TABLES statement.
func (service *Service) ShowTables(ctx context.Context, conn *mysql.Conn, database string) (*mysql.Result, error) {
	return nil, newQueryNotSupportedError("ShowTables")
}
