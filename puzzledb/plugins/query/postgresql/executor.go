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

//nolint:staticcheck
package postgresql

import (
	"time"

	"github.com/cybergarage/go-postgresql/postgresql"
	"github.com/cybergarage/go-postgresql/postgresql/protocol"
	"github.com/cybergarage/go-postgresql/postgresql/query"
	stmt "github.com/cybergarage/go-sqlparser/sql"
	"github.com/cybergarage/puzzledb-go/puzzledb/context"
	sql "github.com/cybergarage/puzzledb-go/puzzledb/plugins/query/sql"
)

// Begin handles a BEGIN query.
func (service *Service) Begin(conn postgresql.Conn, stmt stmt.Begin) (protocol.Responses, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("Begin")
	defer ctx.FinishSpan()

	err := service.Service.Begin(conn, stmt)
	if err != nil {
		return nil, err
	}
	return protocol.NewCommandCompleteResponsesWith(stmt.String())
}

// Commit handles a COMMIT query.
func (service *Service) Commit(conn postgresql.Conn, stmt stmt.Commit) (protocol.Responses, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("Commit")
	defer ctx.FinishSpan()

	err := service.Service.Commit(conn, stmt)
	if err != nil {
		return nil, err
	}
	return protocol.NewCommandCompleteResponsesWith(stmt.String())
}

// Rollback handles a ROLLBACK query.
func (service *Service) Rollback(conn postgresql.Conn, stmt stmt.Rollback) (protocol.Responses, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("Rollback")
	defer ctx.FinishSpan()

	err := service.Service.Rollback(conn, stmt)
	if err != nil {
		return nil, err
	}
	return protocol.NewCommandCompleteResponsesWith(stmt.String())
}

// CreateDatabase handles a CREATE DATABASE query.
func (service *Service) CreateDatabase(conn postgresql.Conn, stmt stmt.CreateDatabase) (protocol.Responses, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("CreateDatabase")
	defer ctx.FinishSpan()

	err := service.Service.CreateDatabase(conn, stmt)
	if err != nil {
		return nil, err
	}
	return protocol.NewCommandCompleteResponsesWith(stmt.String())
}

// CreateTable handles a CREATE TABLE query.
func (service *Service) CreateTable(conn postgresql.Conn, stmt stmt.CreateTable) (protocol.Responses, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("CreateTable")
	defer ctx.FinishSpan()

	err := service.Service.CreateTable(conn, stmt)
	if err != nil {
		return nil, err
	}
	return protocol.NewCommandCompleteResponsesWith(stmt.String())
}

// AlterDatabase handles a ALTER DATABASE query.
func (service *Service) AlterDatabase(conn postgresql.Conn, stmt stmt.AlterDatabase) (protocol.Responses, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("AlterTable")
	defer ctx.FinishSpan()

	err := service.Service.AlterDatabase(conn, stmt)
	if err != nil {
		return nil, err
	}
	return protocol.NewCommandCompleteResponsesWith(stmt.String())
}

// AlterTable handles a ALTER TABLE query.
func (service *Service) AlterTable(conn postgresql.Conn, stmt stmt.AlterTable) (protocol.Responses, error) {
	err := service.Service.AlterTable(conn, stmt)
	if err != nil {
		return nil, err
	}
	return protocol.NewCommandCompleteResponsesWith(stmt.String())
}

// DropDatabase handles a DROP DATABASE query.
func (service *Service) DropDatabase(conn postgresql.Conn, stmt stmt.DropDatabase) (protocol.Responses, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("DropDatabase")
	defer ctx.FinishSpan()

	err := service.Service.DropDatabase(conn, stmt)
	if err != nil {
		return nil, err
	}
	return protocol.NewCommandCompleteResponsesWith(stmt.String())
}

// DropIndex handles a DROP INDEX query.
func (service *Service) DropTable(conn postgresql.Conn, stmt stmt.DropTable) (protocol.Responses, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("DropTable")
	defer ctx.FinishSpan()

	err := service.Service.DropTable(conn, stmt)
	if err != nil {
		return nil, err
	}
	return protocol.NewCommandCompleteResponsesWith(stmt.String())
}

// Insert handles a INSERT query.
func (service *Service) Insert(conn postgresql.Conn, stmt stmt.Insert) (protocol.Responses, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("Insert")
	defer ctx.FinishSpan()

	now := time.Now()
	err := service.Service.Insert(conn, stmt)
	if err != nil {
		return nil, err
	}
	mInsertLatency.Observe(float64(time.Since(now).Milliseconds()))
	return protocol.NewInsertCompleteResponsesWith(1)
}

// Select handles a SELECT query.
func (service *Service) Select(conn postgresql.Conn, stmt stmt.Select) (protocol.Responses, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("Select")
	defer ctx.FinishSpan()

	now := time.Now()

	rs, err := service.Service.Select(conn, stmt)
	if err != nil {
		return nil, err
	}

	mSelectLatency.Observe(float64(time.Since(now).Milliseconds()))

	return query.NewResponseFromResultSet(rs)
}

// Update handles a UPDATE query.
func (service *Service) Update(conn postgresql.Conn, stmt stmt.Update) (protocol.Responses, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("Update")
	defer ctx.FinishSpan()

	now := time.Now()
	rs, err := service.Service.Update(conn, stmt)
	if err != nil {
		return nil, err
	}
	mUpdateLatency.Observe(float64(time.Since(now).Milliseconds()))

	return protocol.NewUpdateCompleteResponsesWith(int(rs.RowsAffected()))
}

// Delete handles a DELETE query.
func (service *Service) Delete(conn postgresql.Conn, stmt stmt.Delete) (protocol.Responses, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("Delete")
	defer ctx.FinishSpan()

	now := time.Now()
	rs, err := service.Service.Delete(conn, stmt)
	if err != nil {
		return nil, err
	}
	mDeleteLatency.Observe(float64(time.Since(now).Milliseconds()))
	return protocol.NewDeleteCompleteResponsesWith(int(rs.RowsAffected()))
}

// Copy handles a COPY query.
func (service *Service) Copy(conn postgresql.Conn, stmt stmt.Copy) (protocol.Responses, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("Copy")
	defer ctx.FinishSpan()

	store := service.Store()

	dbName := conn.Database()
	db, err := store.LookupDatabase(ctx, dbName)
	if err != nil {
		return nil, err
	}

	txn, err := service.Transact(conn, db, true)
	if err != nil {
		return nil, err
	}

	err = txn.SetTimeout(0)
	if err != nil {
		return nil, service.CancelTransactionWithError(ctx, conn, db, txn, err)
	}

	col, err := txn.LookupCollection(ctx, stmt.TableName())
	if err != nil {
		return nil, service.CancelTransactionWithError(ctx, conn, db, txn, err)
	}

	if txn.IsAutoCommit() {
		err := service.CommitTransaction(ctx, conn, db, txn)
		if err != nil {
			return nil, err
		}
	}

	schema, err := sql.NewQuerySchemaFrom(col)
	if err != nil {
		return nil, err
	}

	return postgresql.NewCopyInResponsesFrom(stmt, schema)
}

// Copy handles a COPY DATA protocol.
func (service *Service) CopyData(conn postgresql.Conn, stmt stmt.Copy, stream *postgresql.CopyStream) (protocol.Responses, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("CopyData")
	defer ctx.FinishSpan()

	store := service.Store()

	dbName := conn.Database()
	db, err := store.LookupDatabase(ctx, dbName)
	if err != nil {
		return nil, err
	}

	txn, err := service.Transact(conn, db, true)
	if err != nil {
		return nil, err
	}

	err = txn.SetTimeout(0)
	if err != nil {
		return nil, service.CancelTransactionWithError(ctx, conn, db, txn, err)
	}

	col, err := txn.LookupCollection(ctx, stmt.TableName())
	if err != nil {
		return nil, service.CancelTransactionWithError(ctx, conn, db, txn, err)
	}

	if txn.IsAutoCommit() {
		err := service.CommitTransaction(ctx, conn, db, txn)
		if err != nil {
			return nil, err
		}
	}

	schema, err := sql.NewQuerySchemaFrom(col)
	if err != nil {
		return nil, err
	}

	return postgresql.NewCopyCompleteResponsesFrom(stmt, stream, conn, schema, service)
}
