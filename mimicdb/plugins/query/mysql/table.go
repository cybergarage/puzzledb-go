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
	"sync"

	"github.com/cybergarage/go-mysql/mysql/log"
	"github.com/cybergarage/go-mysql/mysql/query"
)

type Schema = query.Schema
type Row = query.Row
type Rows = query.Rows

// Table represents a destination or source database of query.
type Table struct {
	sync.Mutex
	value string
	*query.Schema
	*query.Rows
}

// NewTableWithNameAndSchema returns a new database with the specified string.
func NewTableWithNameAndSchema(name string, schema *Schema) *Table {
	tbl := &Table{
		value:  name,
		Schema: schema,
		Rows:   query.NewRows(),
	}
	return tbl
}

// NewTable returns a new database.
func NewTable() *Table {
	return NewTableWithNameAndSchema("", nil)
}

// SetSchema sets a specified schema.
func (tbl *Table) SetSchema(schema *Schema) {
	tbl.Schema = schema
}

// Name returns the database name.
func (tbl *Table) Name() string {
	return tbl.value
}

// Insert adds a row.
func (tbl *Table) Insert(row *Row) error {
	return tbl.AddRow(row)
}

// Select returns only matched and projected rows by the specified conditions and the columns.
func (tbl *Table) Select(cond *query.Condition) (*Rows, error) {
	matchedRows := tbl.FindMatchedRows(cond)
	return matchedRows, nil
}

// Update updates rows which are satisfied by the specified columns and conditions.
func (tbl *Table) Update(columns *query.Columns, cond *query.Condition) (int, error) {
	matchedRows := tbl.FindMatchedRows(cond)
	nUpdatedRows := 0
	for _, matchedRow := range matchedRows.Rows() {
		err := matchedRow.Update(columns)
		if err != nil {
			return 0, err
		}
		nUpdatedRows++
	}
	return nUpdatedRows, nil
}

// Delete deletes rows which are satisfied by the specified conditions.
func (tbl *Table) Delete(cond *query.Condition) (int, error) {
	matchedRows := tbl.FindMatchedRows(cond)
	nDeletedRows := 0
	for _, matchedRow := range matchedRows.Rows() {
		nDeletedRows += int(tbl.DeleteRow(matchedRow))
	}
	return nDeletedRows, nil
}

// DeleteAll deletes all rows in the table.
func (tbl *Table) DeleteAll() int {
	rows := tbl.Rows.Rows()
	nRowsCnt := len(rows)
	tbl.Rows = query.NewRows()
	return nRowsCnt
}

// String returns the string representation.
func (tbl *Table) String() string {
	return tbl.value
}

// Dump outputs all row values for debug.
func (tbl *Table) Dump() {
	log.Debug("%s", tbl.Name())
	for n, row := range tbl.Rows.Rows() {
		log.Debug("[%d] %s", n, row.String())
	}
}
