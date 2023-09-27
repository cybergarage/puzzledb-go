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
	"github.com/cybergarage/go-postgresql/postgresql"
	"github.com/cybergarage/go-postgresql/postgresql/protocol/message"
	"github.com/cybergarage/go-postgresql/postgresql/query"
	"github.com/cybergarage/puzzledb-go/puzzledb/context"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/query/sql"
)

// CreateDatabase handles a CREATE DATABASE query.
func (service *Service) CreateDatabase(conn *postgresql.Conn, stmt *query.CreateDatabase) (message.Responses, error) {
	err := service.Service.CreateDatabase(conn, stmt)
	if err != nil {
		return nil, err
	}
	return message.NewCommandCompleteResponsesWith(stmt.String())
}

// CreateTable handles a CREATE TABLE query.
func (service *Service) CreateTable(conn *postgresql.Conn, stmt *query.CreateTable) (message.Responses, error) {
	err := service.Service.CreateTable(conn, stmt)
	if err != nil {
		return nil, err
	}
	return message.NewCommandCompleteResponsesWith(stmt.String())
}

// AlterDatabase handles a ALTER DATABASE query.
func (service *Service) AlterDatabase(conn *postgresql.Conn, stmt *query.AlterDatabase) (message.Responses, error) {
	return nil, query.NewErrNotImplemented("ALTER TABLE")
}

// AlterTable handles a ALTER TABLE query.
func (service *Service) AlterTable(conn *postgresql.Conn, stmt *query.AlterTable) (message.Responses, error) {
	return nil, query.NewErrNotImplemented("ALTER TABLE")
}

// DropDatabase handles a DROP DATABASE query.
func (service *Service) DropDatabase(conn *postgresql.Conn, stmt *query.DropDatabase) (message.Responses, error) {
	err := service.Service.DropDatabase(conn, stmt)
	if err != nil {
		return nil, err
	}
	return message.NewCommandCompleteResponsesWith(stmt.String())
}

// DropIndex handles a DROP INDEX query.
func (service *Service) DropTable(conn *postgresql.Conn, stmt *query.DropTable) (message.Responses, error) {
	err := service.Service.DropTable(conn, stmt)
	if err != nil {
		return nil, err
	}
	return message.NewCommandCompleteResponsesWith(stmt.String())
}

// Insert handles a INSERT query.
func (service *Service) Insert(conn *postgresql.Conn, stmt *query.Insert) (message.Responses, error) {
	err := service.Service.Insert(conn, stmt)
	if err != nil {
		return nil, err
	}
	return message.NewInsertCompleteResponsesWith(1)
}

// Select handles a SELECT query.
func (service *Service) Select(conn *postgresql.Conn, stmt *query.Select) (message.Responses, error) { //nolint:gocognit
	ctx, txn, col, rs, err := service.Service.Select(conn, stmt)
	defer ctx.FinishSpan()
	if err != nil {
		return nil, err
	}

	// Schema

	schema, err := sql.NewQuerySchemaFrom(col)
	if err != nil {
		return nil, err
	}

	// Responses

	res := message.NewResponses()

	// Row description response

	selectors := stmt.Selectors()
	if selectors.IsSelectAll() {
		selectors = schema.Selectors()
	}

	rowDesc := message.NewRowDescription()
	for n, selector := range selectors {
		field, err := query.NewRowFieldFrom(schema, selector, n)
		if err != nil {
			return nil, err
		}
		rowDesc.AppendField(field)
	}
	res = res.Append(rowDesc)

	// Data row response

	nRows := 0
	if !selectors.HasAggregateFunction() {
		offset := stmt.Limit().Offset()
		limit := stmt.Limit().Limit()
		rowNo := 0
		for rs.Next() {
			rowNo++
			if 0 < offset && rowNo <= offset {
				continue
			}
			obj := rs.Object()
			row, err := sql.NewRowFrom(obj)
			if err != nil {
				return nil, err
			}
			dataRow, err := query.NewDataRowForSelectors(schema, rowDesc, selectors, row)
			if err != nil {
				return nil, err
			}
			res = res.Append(dataRow)
			nRows++
			if 0 < limit && limit <= nRows {
				break
			}
		}
	} else {
		groupBy := stmt.GroupBy().Column()
		queryRows := []query.Row{}
		for rs.Next() {
			obj := rs.Object()
			row, err := sql.NewRowFrom(obj)
			if err != nil {
				return nil, err
			}
			queryRows = append(queryRows, row)
		}
		dataRows, err := query.NewDataRowsForAggregateFunction(schema, rowDesc, selectors, queryRows, groupBy)
		if err != nil {
			return nil, err
		}
		for _, dataRow := range dataRows {
			res = res.Append(dataRow)
			nRows++
		}
	}

	cmpRes, err := message.NewSelectCompleteWith(nRows)
	if err != nil {
		return nil, err
	}
	res = res.Append(cmpRes)

	// Commit the transaction

	err = txn.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Update handles a UPDATE query.
func (service *Service) Update(conn *postgresql.Conn, stmt *query.Update) (message.Responses, error) {
	n, err := service.Service.Update(conn, stmt)
	if err != nil {
		return nil, err
	}
	return message.NewUpdateCompleteResponsesWith(n)
}

// Delete handles a DELETE query.
func (service *Service) Delete(conn *postgresql.Conn, stmt *query.Delete) (message.Responses, error) {
	n, err := service.Service.Delete(conn, stmt)
	if err != nil {
		return nil, err
	}
	return message.NewDeleteCompleteResponsesWith(n)
}

// Copy handles a COPY query.
func (service *Service) Copy(conn *postgresql.Conn, stmt *query.Copy) (message.Responses, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("Copy")
	defer ctx.FinishSpan()

	store := service.Store()

	dbName := conn.Database()
	db, err := store.GetDatabase(ctx, dbName)
	if err != nil {
		return nil, err
	}

	txn, err := db.Transact(true)
	if err != nil {
		return nil, err
	}

	col, err := txn.GetCollection(ctx, stmt.TableName())
	if err != nil {
		return nil, service.CancelTransactionWithError(ctx, txn, err)
	}

	err = txn.Commit(ctx)
	if err != nil {
		return nil, err
	}

	schema, err := sql.NewQuerySchemaFrom(col)
	if err != nil {
		return nil, err
	}

	return postgresql.NewCopyInResponsesFrom(stmt, schema)
}

// Copy handles a COPY DATA message.
func (service *Service) CopyData(conn *postgresql.Conn, stmt *query.Copy, stream *postgresql.CopyStream) (message.Responses, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("CopyData")
	defer ctx.FinishSpan()

	store := service.Store()

	dbName := conn.Database()
	db, err := store.GetDatabase(ctx, dbName)
	if err != nil {
		return nil, err
	}

	txn, err := db.Transact(true)
	if err != nil {
		return nil, err
	}

	col, err := txn.GetCollection(ctx, stmt.TableName())
	if err != nil {
		return nil, service.CancelTransactionWithError(ctx, txn, err)
	}

	err = txn.Commit(ctx)
	if err != nil {
		return nil, err
	}

	schema, err := sql.NewQuerySchemaFrom(col)
	if err != nil {
		return nil, err
	}

	return postgresql.NewCopyCompleteResponsesFrom(stmt, stream, conn, schema, service.Executor)
}
