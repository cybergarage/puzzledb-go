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
	"github.com/cybergarage/mimicdb/mimicdb/errors"
	"github.com/hashicorp/go-memdb"
)

// Memdb represents a Memdb instance.
type Memdb struct {
	*memdb.MemDB
}

type document struct {
	Key   []byte
	Value []byte
}

// New returns a new memdb store instance.
func NewStore() *Memdb {
	return &Memdb{
		MemDB: nil,
	}
}

// Open opens the specified store.
func (db *Memdb) Open(name string) error {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"document": &memdb.TableSchema{
				Name: "document",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Key"},
					},
				},
			},
		},
	}
	var err error
	db.MemDB, err = memdb.NewMemDB(schema)
	if err != nil {
		return err
	}
	return nil
}

// Transact opens a transaction.
func (db *Memdb) Transact() (mimicdb.Transaction, error) {
	if db.MemDB == nil {
		return nil, errors.DatabaseNotFound
	}
	return newTransaction(db.MemDB.Txn(true)), nil
}

// Close closes this store.
func (db *Memdb) Close() error {
	return nil
}

// Start starts this memdb.
func (db *Memdb) Start() error {
	return nil
}

// Stop stops this memdb.
func (db Memdb) Stop() error {
	return nil
}
