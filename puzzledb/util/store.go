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

package document

import (
	"fmt"

	"github.com/cybergarage/puzzledb-go/puzzledb/context"
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
	plugin "github.com/cybergarage/puzzledb-go/puzzledb/plugins/store"
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
)

// Store represents a new document store utility instance.
type Store struct {
	store        plugin.DocumentStore
	dumpChildren bool
}

// StoreOption represents a store option.
type StoreOption func(*Store)

func WithStoreDumpChildren(v bool) StoreOption {
	return func(s *Store) {
		s.dumpChildren = v
	}
}

// NewStoreWith returns a new document store utility instance with the specified store.
func NewStoreWith(docStore plugin.DocumentStore, opts ...StoreOption) *Store {
	store := &Store{
		store:        docStore,
		dumpChildren: true,
	}
	for _, opt := range opts {
		opt(store)
	}
	return store
}

// Dump returns a string array representation of the document store.
func (store *Store) Dump() ([]string, error) {
	ctx := context.NewContext()
	allLines := []string{}
	dbs, err := store.store.ListDatabases(ctx)
	if err != nil {
		return allLines, err
	}
	for n, db := range dbs {
		allLines = append(allLines, fmt.Sprintf("[%d]: %s", n, db.Name()))
		if !store.dumpChildren {
			continue
		}
		lines, err := store.DumpDatabase(ctx, db)
		if err != nil {
			return allLines, err
		}
		allLines = append(allLines, lines...)
	}
	return allLines, nil
}

// DumpDatabase returns a string array representation of the specified database.
func (store *Store) DumpDatabase(ctx context.Context, db store.Database) ([]string, error) {
	allLines := []string{}
	txn, err := db.Transact(false)
	if err != nil {
		return allLines, err
	}

	defer func() {
		txn.Commit(ctx)
	}()

	cols, err := txn.ListCollections(ctx)
	if err != nil {
		return allLines, err
	}
	for n, col := range cols {
		allLines = append(allLines, fmt.Sprintf("[%d]: %s", n, col.Name()))
		if !store.dumpChildren {
			continue
		}
		lines, err := store.DumpCollection(ctx, db, col, txn)
		if err != nil {
			return allLines, err
		}
		allLines = append(allLines, lines...)
	}
	return allLines, nil
}

func (store *Store) DumpCollection(ctx context.Context, db store.Database, col store.Collection, txn store.Transaction) ([]string, error) {
	allLines := []string{}
	key := document.NewKeyWith(db.Name(), col.Name())
	rs, err := txn.FindObjects(ctx, key)
	if err != nil {
		return allLines, err
	}
	defer rs.Close()
	for rs.Next() {
		doc, err := rs.Document()
		if err != nil {
			return allLines, err
		}
		allLines = append(allLines, fmt.Sprintf("[%s]: %s", doc.Key(), doc.Object()))
	}
	return allLines, nil
}
