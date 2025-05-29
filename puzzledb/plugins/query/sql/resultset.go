// Copyright (C) 2024 The PuzzleDB Authors.
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
	"github.com/cybergarage/go-sqlparser/sql"
	"github.com/cybergarage/go-sqlparser/sql/query/response/resultset"
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
)

type resultSet struct {
	storeDB store.Database
	schema  resultset.Schema
	storeRs store.ResultSet
}

// NewResultSet returns a new result set.
func NewResultSetFrom(storeDB store.Database, storeCol store.Collection, storeRs store.ResultSet) (sql.ResultSet, error) {
	rs := &resultSet{
		storeDB: storeDB,
		schema:  nil,
		storeRs: storeRs,
	}
	schema, err := NewQuerySchemaFrom(storeCol)
	if err != nil {
		return nil, err
	}
	rs.schema = resultset.NewSchema(
		resultset.WithSchemaDatabaseName(rs.storeDB.Name()),
		resultset.WithSchemaTableName(schema.TableName()),
		resultset.WithSchemaColumns(resultset.NewColumnsFrom(schema.Columns())),
	)
	return rs, nil
}

// Row returns the current row.
func (rs *resultSet) Row() (resultset.Row, error) {
	rsDoc, err := rs.storeRs.Document()
	if err != nil {
		return nil, err
	}
	row, err := NewRowFromObject(rs.schema, rsDoc.Object())
	if err != nil {
		return nil, err
	}
	return row, nil
}

// Schema returns the schema.
func (rs *resultSet) Schema() resultset.Schema {
	return rs.schema
}

// RowsAffected returns the number of rows affected.
func (rs *resultSet) RowsAffected() uint {
	return 0
}

// Next returns the next row.
func (rs *resultSet) Next() bool {
	if rs.storeRs == nil {
		return false
	}
	return rs.storeRs.Next()
}

// Close closes the resultset.
func (rs *resultSet) Close() error {
	return nil
}
