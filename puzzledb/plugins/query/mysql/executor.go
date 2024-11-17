// Copyright (C) 2024 The go-mysql Authors. All rights reserved.
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
	"time"

	"github.com/cybergarage/go-mysql/mysql/net"
	"github.com/cybergarage/go-mysql/mysql/protocol"
	"github.com/cybergarage/go-sqlparser/sql"
	"github.com/cybergarage/puzzledb-go/puzzledb/context"
)

// Conn represents a connection.
type Conn = net.Conn

// Response represents a response.
type Response = protocol.Response

// Begin handles a BEGIN query.
func (service *Service) Begin(conn Conn, stmt sql.Begin) (Response, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("Begin")
	defer ctx.FinishSpan()

	return protocol.NewResponseWithError(service.Service.Begin(conn, stmt))
}

// Commit handles a COMMIT query.
func (service *Service) Commit(conn Conn, stmt sql.Commit) (Response, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("Commit")
	defer ctx.FinishSpan()

	return protocol.NewResponseWithError(service.Service.Commit(conn, stmt))
}

// Rollback handles a ROLLBACK query.
func (service *Service) Rollback(conn Conn, stmt sql.Rollback) (Response, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("Rollback")
	defer ctx.FinishSpan()

	return protocol.NewResponseWithError(service.Service.Rollback(conn, stmt))
}

// CreateDatabase handles a CREATE DATABASE query.
func (service *Service) CreateDatabase(conn Conn, stmt sql.CreateDatabase) (Response, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("CreateDatabase")
	defer ctx.FinishSpan()

	return protocol.NewResponseWithError(service.Service.CreateDatabase(conn, stmt))
}

// CreateTable handles a CREATE TABLE query.
func (service *Service) CreateTable(conn Conn, stmt sql.CreateTable) (Response, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("CreateTable")
	defer ctx.FinishSpan()

	return protocol.NewResponseWithError(service.Service.CreateTable(conn, stmt))
}

// AlterDatabase handles a ALTER DATABASE query.
func (service *Service) AlterDatabase(conn Conn, stmt sql.AlterDatabase) (Response, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("AlterDatabase")
	defer ctx.FinishSpan()

	return protocol.NewResponseWithError(service.Service.AlterDatabase(conn, stmt))
}

// AlterTable handles a ALTER TABLE query.
func (service *Service) AlterTable(conn Conn, stmt sql.AlterTable) (Response, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("AlterTable")
	defer ctx.FinishSpan()

	return protocol.NewResponseWithError(service.Service.AlterTable(conn, stmt))
}

// DropDatabase handles a DROP DATABASE query.
func (service *Service) DropDatabase(conn Conn, stmt sql.DropDatabase) (Response, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("DropDatabase")
	defer ctx.FinishSpan()

	return protocol.NewResponseWithError(service.Service.DropDatabase(conn, stmt))
}

// DropIndex handles a DROP INDEX query.
func (service *Service) DropTable(conn Conn, stmt sql.DropTable) (Response, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("DropTable")
	defer ctx.FinishSpan()

	return protocol.NewResponseWithError(service.Service.DropTable(conn, stmt))
}

// Insert handles a INSERT query.
func (service *Service) Insert(conn Conn, stmt sql.Insert) (Response, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("Insert")
	defer ctx.FinishSpan()

	now := time.Now()
	err := service.Service.Insert(conn, stmt)
	if err != nil {
		return nil, err
	}
	mInsertLatency.Observe(float64(time.Since(now).Milliseconds()))

	return protocol.NewResponseWithError(err)
}

// Select handles a SELECT query.
func (service *Service) Select(conn Conn, stmt sql.Select) (Response, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("Select")
	defer ctx.FinishSpan()

	now := time.Now()
	rs, err := service.Service.Select(conn, stmt)
	if err != nil {
		return nil, err
	}
	mSelectLatency.Observe(float64(time.Since(now).Milliseconds()))

	return protocol.NewTextResultSetFromResultSet(rs)
}

// Update handles a UPDATE query.
func (service *Service) Update(conn Conn, stmt sql.Update) (Response, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("Update")
	defer ctx.FinishSpan()

	now := time.Now()
	rs, err := service.Service.Update(conn, stmt)
	if err != nil {
		return protocol.NewResponseWithError(err)
	}
	mUpdateLatency.Observe(float64(time.Since(now).Milliseconds()))

	return protocol.NewOK(
		protocol.WithOKAffectedRows(uint64(rs.RowsAffected())),
	)
}

// Delete handles a DELETE query.
func (service *Service) Delete(conn Conn, stmt sql.Delete) (Response, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("Delete")
	defer ctx.FinishSpan()

	now := time.Now()
	rs, err := service.Service.Delete(conn, stmt)
	if err != nil {
		return protocol.NewResponseWithError(err)
	}
	mDeleteLatency.Observe(float64(time.Since(now).Milliseconds()))

	return protocol.NewOK(
		protocol.WithOKAffectedRows(uint64(rs.RowsAffected())),
	)
}

// Use handles a USE query.
func (service *Service) Use(conn Conn, stmt sql.Use) (Response, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("Use")
	defer ctx.FinishSpan()

	return protocol.NewResponseWithError(service.Service.Use(conn, stmt))
}
