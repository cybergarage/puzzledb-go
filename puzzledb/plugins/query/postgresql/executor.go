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

package postgresql

import (
	"github.com/cybergarage/go-logger/log"
	"github.com/cybergarage/go-postgresql/postgresql"
	"github.com/cybergarage/go-postgresql/postgresql/protocol/message"
	"github.com/cybergarage/go-postgresql/postgresql/query"
	"github.com/cybergarage/puzzledb-go/puzzledb/context"
)

// CreateDatabase handles a CREATE DATABASE query.
func (service *Service) CreateDatabase(conn *postgresql.Conn, stmt *query.CreateDatabase) (message.Responses, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("CreateDatabase")
	defer ctx.FinishSpan()

	dbName := stmt.DatabaseName()

	store := service.Store()
	_, err := store.GetDatabase(ctx, dbName)
	if err == nil {
		if stmt.IfNotExists() {
			return message.NewCommandCompleteResponsesWith(stmt.String())
		}
		return nil, postgresql.NewErrDatabaseExist(dbName)
	}

	err = store.CreateDatabase(ctx, dbName)
	if err != nil {
		return nil, err
	}

	// Post a event message to the coordinator.

	err = service.PostDatabaseCreateMessage(dbName)
	if err != nil {
		log.Error(err)
	}

	return message.NewCommandCompleteResponsesWith(stmt.String())
}

// CreateTable handles a CREATE TABLE query.
func (service *Service) CreateTable(conn *postgresql.Conn, stmt *query.CreateTable) (message.Responses, error) {
	return nil, postgresql.NewErrNotImplemented("CREATE TABLE")
}

// CreateIndex handles a CREATE INDEX query.
func (service *Service) CreateIndex(conn *postgresql.Conn, stmt *query.CreateIndex) (message.Responses, error) {
	return nil, postgresql.NewErrNotImplemented("CREATE INDEX")
}

// DropDatabase handles a DROP DATABASE query.
func (service *Service) DropDatabase(conn *postgresql.Conn, stmt *query.DropDatabase) (message.Responses, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("DropDatabase")
	defer ctx.FinishSpan()

	store := service.Store()
	dbName := stmt.DatabaseName()

	// Check if the database exists.

	_, err := store.GetDatabase(ctx, dbName)
	if err != nil {
		if stmt.IfExists() {
			return message.NewCommandCompleteResponsesWith(stmt.String())
		}
		return nil, err
	}

	// Drop the specified database.

	err = store.RemoveDatabase(ctx, dbName)
	if err != nil {
		return nil, err
	}

	// Post a event message to the coordinator.

	err = service.PostDatabaseDropMessage(dbName)
	if err != nil {
		log.Error(err)
	}

	return message.NewCommandCompleteResponsesWith(stmt.String())
}

// DropIndex handles a DROP INDEX query.
func (service *Service) DropTable(conn *postgresql.Conn, stmt *query.DropTable) (message.Responses, error) {
	return nil, postgresql.NewErrNotImplemented("DROP TABLE")
}

// Insert handles a INSERT query.
func (service *Service) Insert(conn *postgresql.Conn, stmt *query.Insert) (message.Responses, error) {
	return nil, postgresql.NewErrNotImplemented("INSERT")
}

// Select handles a SELECT query.
func (service *Service) Select(conn *postgresql.Conn, stmt *query.Select) (message.Responses, error) {
	return nil, postgresql.NewErrNotImplemented("SELECT")
}

// Update handles a UPDATE query.
func (service *Service) Update(conn *postgresql.Conn, stmt *query.Update) (message.Responses, error) {
	return nil, postgresql.NewErrNotImplemented("UPDATE")
}

// Delete handles a DELETE query.
func (service *Service) Delete(conn *postgresql.Conn, stmt *query.Delete) (message.Responses, error) {
	return nil, postgresql.NewErrNotImplemented("DELETE")
}