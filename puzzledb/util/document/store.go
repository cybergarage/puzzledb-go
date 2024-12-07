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
	"github.com/cybergarage/puzzledb-go/puzzledb/context"
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
	plugin "github.com/cybergarage/puzzledb-go/puzzledb/plugins/store"
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
)

// Store represents a new document store utility instance.
type Store struct {
	store plugin.DocumentStore
}

type DumpOptions struct {
	// DumpAll specifies whether to dump all the documents.
	DumpAll bool
}

// NewStoreWith returns a new document store utility instance with the specified store.
func NewStoreWith(store plugin.DocumentStore) *Store {
	return &Store{
		store: store,
	}
}

// Dump returns a string array representation of the document store.
func (doc *Store) Dump(opts DumpOptions) ([]string, error) {
	allLines := []string{}
	dbs, err := doc.store.ListDatabases(nil)
	if err != nil {
		return allLines, err
	}
	ctx := context.NewContext()
	for _, db := range dbs {
		lines, err := doc.dumpDatabase(ctx, db)
		if err != nil {
			return allLines, err
		}
		allLines = append(allLines, lines...)
	}
	return allLines, nil
}

// dumpDatabase returns a string array representation of the specified database.
func (doc *Store) dumpDatabase(ctx context.Context, db store.Database) ([]string, error) {
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
	for _, col := range cols {
		lines, err := doc.dumpCollection(ctx, db, col, txn)
		if err != nil {
			return allLines, err
		}
		allLines = append(allLines, lines...)
	}
	return allLines, nil
}

func (doc *Store) dumpCollection(ctx context.Context, db store.Database, col store.Collection, txn store.Transaction) ([]string, error) {
	allLines := []string{}
	key := document.NewKeyWith(db.Name(), col.Name())
	rs, err := txn.FindObjects(ctx, key)
	if err != nil {
		return allLines, err
	}
	for rs.Next() {
		// rs.Object()
		// key := rs.Key()
		// obj := rs.Object()
		// line := fmt.Sprintf("[%s]: %s", key, obj)
		// allLines = append(allLines, line)
	}
	return allLines, nil
}
