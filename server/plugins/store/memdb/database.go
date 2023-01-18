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

package memdb

import (
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
	"github.com/hashicorp/go-memdb"
)

const (
	tableName    = "document"
	idFieldName  = "id"
	keyFieldName = "Key"
)

// Database represents a database.
type Database struct {
	ID string
	*memdb.MemDB
}

// NewDatabaseWithID returns a new database with the specified ID.
func NewDatabaseWithID(id string) (*Database, error) {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			tableName: &memdb.TableSchema{
				Name: tableName,
				Indexes: map[string]*memdb.IndexSchema{
					idFieldName: &memdb.IndexSchema{
						Name:    idFieldName,
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Key"},
					},
				},
			},
		},
	}
	memDB, err := memdb.NewMemDB(schema)
	if err != nil {
		return nil, err
	}
	return &Database{
		ID:    id,
		MemDB: memDB,
	}, nil
}

// Name returns the unique name.
func (db *Database) Name() string {
	return db.ID
}

// Transact begin a new transaction.
func (db *Database) Transact(write bool) (store.Transaction, error) {
	if db.MemDB == nil {
		return nil, store.DatabaseNotFound
	}
	return newTransaction(db.MemDB.Txn(write)), nil
}
