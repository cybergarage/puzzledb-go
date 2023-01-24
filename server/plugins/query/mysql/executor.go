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
	"fmt"

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
	if err != nil {
		return mysql.NewResult(), fmt.Errorf(errorDatabaseFound, dbName)
	}

	err = store.CreateDatabase(dbName)
	if err != nil {
		return mysql.NewResult(), err
	}

	return mysql.NewResult(), nil
}

// AlterDatabase should handle a ALTER database statement.
func (service *Service) AlterDatabase(ctx context.Context, conn *mysql.Conn, stmt *query.Database) (*mysql.Result, error) {
	log.Debugf("%v", stmt)
	return mysql.NewResult(), nil
}

// DropDatabase should handle a DROP database statement.
func (service *Service) DropDatabase(ctx context.Context, conn *mysql.Conn, stmt *query.Database) (*mysql.Result, error) {
	log.Debugf("%v", stmt)
	return mysql.NewResult(), nil
}

// CreateTable should handle a CREATE table statement.
func (service *Service) CreateTable(ctx context.Context, conn *mysql.Conn, stmt *query.Schema) (*mysql.Result, error) {
	log.Debugf("%v", stmt)
	return mysql.NewResult(), nil
}

// AlterTable should handle a ALTER table statement.
func (srv *Service) AlterTable(ctx context.Context, conn *mysql.Conn, stmt *query.Schema) (*mysql.Result, error) {
	log.Debugf("%v", stmt)
	return mysql.NewResult(), nil
}

// DropTable should handle a DROP table statement.
func (srv *Service) DropTable(ctx context.Context, conn *mysql.Conn, stmt *query.Schema) (*mysql.Result, error) {
	dbName := conn.Database
	db, ok := srv.GetDatabase(dbName)
	if !ok {
		return nil, fmt.Errorf(errorDatabaseNotFound, dbName)
	}
	tableName := stmt.TableName()
	table, ok := db.GetTable(tableName)
	if !ok {
		return mysql.NewResult(), nil
	}

	if !db.DropTable(table) {
		return nil, fmt.Errorf("%s could not deleted", table.TableName())
	}

	return mysql.NewResult(), nil
}

// RenameTable should handle a RENAME table statement.
func (srv *Service) RenameTable(ctx context.Context, conn *mysql.Conn, stmt *query.Schema) (*mysql.Result, error) {
	log.Debugf("%v", stmt)
	return mysql.NewResult(), nil
}

// TruncateTable should handle a TRUNCATE table statement.
func (srv *Service) TruncateTable(ctx context.Context, conn *mysql.Conn, stmt *query.Schema) (*mysql.Result, error) {
	log.Debugf("%v", stmt)
	return mysql.NewResult(), nil
}

// Insert should handle a INSERT statement.
func (srv *Service) Insert(ctx context.Context, conn *mysql.Conn, stmt *query.Insert) (*mysql.Result, error) {
	log.Debugf("%v", stmt)
	dbName := conn.Database
	tableName := stmt.TableName()
	table, ok := srv.GetTableWithDatabase(dbName, tableName)
	if !ok {
		return mysql.NewResult(), fmt.Errorf(errorTableNotFound, dbName, tableName)
	}

	row, err := query.NewRowWithInsert(stmt)
	if err != nil {
		return nil, err
	}

	if err := table.AddRow(row); err != nil {
		return nil, err
	}

	table.Dump()

	return mysql.NewResultWithRowsAffected(1), nil
}

// Update should handle a UPDATE statement.
func (srv *Service) Update(ctx context.Context, conn *mysql.Conn, stmt *query.Update) (*mysql.Result, error) {
	log.Debugf("%v", stmt)

	dbName := conn.Database
	cond := stmt.Where

	database, ok := srv.GetDatabase(dbName)
	if !ok {
		return nil, fmt.Errorf(errorDatabaseFound, dbName)
	}

	nEffectedRows := uint64(0)
	for _, table := range stmt.Tables() {
		tableName, err := table.Name()
		if err != nil {
			return nil, err
		}
		table, ok := database.GetTable(tableName)
		if !ok {
			return nil, fmt.Errorf(errorTableNotFound, dbName, tableName)
		}

		columns, err := stmt.Columns()
		if err != nil {
			return nil, err
		}

		nUpdatedRows, err := table.Update(columns, cond)
		if err != nil {
			return nil, err
		}
		nEffectedRows += uint64(nUpdatedRows)
	}

	return mysql.NewResultWithRowsAffected(nEffectedRows), nil
}

// Delete should handle a DELETE statement.
func (srv *Service) Delete(ctx context.Context, conn *mysql.Conn, stmt *query.Delete) (*mysql.Result, error) {
	dbName := conn.Database
	cond := stmt.Where

	database, ok := srv.GetDatabase(dbName)
	if !ok {
		return nil, fmt.Errorf(errorDatabaseFound, dbName)
	}

	nEffectedRows := uint64(0)
	for _, table := range stmt.Tables() {
		tableName, err := table.Name()
		if err != nil {
			return nil, err
		}
		table, ok := database.GetTable(tableName)
		if !ok {
			return nil, fmt.Errorf(errorTableNotFound, dbName, tableName)
		}

		nDeletedRows, err := table.Delete(cond)
		if err != nil {
			return nil, err
		}
		nEffectedRows += uint64(nDeletedRows)
	}

	return mysql.NewResultWithRowsAffected(nEffectedRows), nil
}

// Select should handle a SELECT statement.
func (srv *Service) Select(ctx context.Context, conn *mysql.Conn, stmt *query.Select) (*mysql.Result, error) {
	log.Debugf("%v", stmt)

	dbName := conn.Database
	database, ok := srv.GetDatabase(dbName)
	if !ok {
		return nil, fmt.Errorf(errorDatabaseFound, dbName)
	}

	// NOTE: Select scans only a first table

	tables := stmt.From()
	tableName, err := tables[0].Name()
	if err != nil {
		return nil, err
	}

	table, ok := database.GetTable(tableName)
	if !ok {
		// TODO: Support dummy dual table for MySQL connector 5.1.49
		if tableName == "dual" {
			return mysql.NewResult(), nil
		}
		return nil, fmt.Errorf(errorTableNotFound, dbName, tableName)
	}

	cond := stmt.Where
	matchedRows := table.FindMatchedRows(cond)

	return mysql.NewResultWithRows(database.Database, table.Schema, matchedRows)
}

// ShowDatabases should handle a SHOW DATABASES statement.
func (srv *Service) ShowDatabases(ctx context.Context, conn *mysql.Conn) (*mysql.Result, error) {
	return mysql.NewResult(), nil
}

// ShowTables should handle a SHOW TABLES statement.
func (srv *Service) ShowTables(ctx context.Context, conn *mysql.Conn, database string) (*mysql.Result, error) {
	return mysql.NewResult(), nil
}
