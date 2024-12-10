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
	"errors"

	"github.com/cybergarage/puzzledb-go/puzzledb/document"
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
	"github.com/cybergarage/puzzledb-go/puzzledb/store/kv"
	"github.com/hashicorp/go-memdb"
)

const (
	tableName   = "document"
	idName      = "id"
	idFieldName = "Key"
	prefix      = "_prefix"
)

// Database represents a database.
type Database struct {
	*memdb.MemDB
	document.KeyCoder
}

// Document represents a document.
type Document struct {
	Key   []byte
	Value []byte
}

// BinaryFieldIndexer is a custom field indexer for binary keys.
type BinaryFieldIndexer struct {
	Field string
}

// FromArgs extracts the binary key from the arguments.
func (indexer *BinaryFieldIndexer) FromArgs(args ...interface{}) ([]byte, error) {
	if len(args) < 1 {
		return nil, errors.New("invalid arguments")
	}
	_, bytes, err := indexer.FromObject(args[0])
	return bytes, err
}

// FromObject extracts the binary key from the object.
func (indexer *BinaryFieldIndexer) FromObject(obj any) (bool, []byte, error) {
	binKey, ok := obj.([]byte)
	if ok {
		return true, binKey, nil
	}
	doc, ok := obj.(*Document)
	if ok {
		return true, doc.Key, nil
	}
	return false, nil, nil
}

// PrefixFromArgs returns the prefix of the key.
func (indexer *BinaryFieldIndexer) PrefixFromArgs(args ...interface{}) ([]byte, error) {
	if len(args) < 1 {
		return nil, errors.New("invalid arguments")
	}
	obj := args[0]
	binKey, ok := obj.([]byte)
	if ok {
		return binKey, nil
	}
	doc, ok := obj.(*Document)
	if ok {
		return doc.Key, nil
	}
	return nil, errors.New("invalid object")
}

// NewDatabaseWith returns a new database.
func NewDatabaseWith(coder document.KeyCoder) (*Database, error) {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			tableName: {
				Name: tableName,
				Indexes: map[string]*memdb.IndexSchema{
					idName: {
						Name:         idName,
						AllowMissing: false,
						Unique:       true,
						Indexer: &BinaryFieldIndexer{
							Field: idFieldName,
						},
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
		MemDB:    memDB,
		KeyCoder: coder,
	}, nil
}

// Transact begin a new transaction.
func (db *Database) Transact(write bool) (kv.Transaction, error) {
	if db.MemDB == nil {
		return nil, store.NewErrDatabaseNotExist("memdb")
	}
	return newTransaction(db.MemDB.Txn(write), db.KeyCoder), nil
}
