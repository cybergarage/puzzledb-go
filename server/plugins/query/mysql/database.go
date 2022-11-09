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
	"github.com/cybergarage/go-mysql/mysql/query"
)

// Database represents a destination or source database of query.
type Database struct {
	*query.Database
	tables map[string]*Table
}

// NewDatabaseWithName returns a new database with the specified string.
func NewDatabaseWithName(name string) *Database {
	ks := &Database{
		Database: query.NewDatabaseWithName(name),
		tables:   map[string]*Table{},
	}
	return ks
}

// NewDatabase returns a new database.
func NewDatabase() *Database {
	return NewDatabaseWithName("")
}

// AddTable adds a specified table into the database.
func (ks *Database) AddTable(table *Table) {
	tableName := table.Name()
	ks.tables[tableName] = table
}

// AddTables adds a specified tables into the database.
func (ks *Database) AddTables(tables []*Table) {
	for _, table := range tables {
		ks.AddTable(table)
	}
}

// DropTable remove the specified table.
func (ks *Database) DropTable(table *Table) bool {
	name := table.TableName()
	delete(ks.tables, name)
	_, ok := ks.tables[name]
	return !ok
}

// GetTable returns a table with the specified name.
func (ks *Database) GetTable(name string) (*Table, bool) {
	table, ok := ks.tables[name]
	return table, ok
}

// GetTables returns all tables in the database.
func (ks *Database) GetTables() []*Table {
	tables := make([]*Table, 0)
	for _, table := range ks.tables {
		tables = append(tables, table)
	}
	return tables
}
