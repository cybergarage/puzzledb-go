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

package fdb

import (
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	store "github.com/cybergarage/puzzledb-go/puzzledb/server/plugins/store/kv"
	"github.com/cybergarage/puzzledb-go/puzzledb/store/kv"
)

var RequiredAPIVersion = 630

// Store represents a FoundationDB store service instance.
type Store struct {
	fdb.Database
}

// New returns a new memdb store instance.
func NewStore() store.Service {
	return &Store{}
}

// CreateDatabase creates a new database.
func (store *Store) CreateDatabase(name string) error {
	return nil
}

// GetDatabase retruns the specified database.
func (store *Store) GetDatabase(id string) (kv.Database, error) {
	return newDatabaseWith(id, store.Database), nil
}

// Start starts this memdb.
func (store *Store) Start() error {
	err := fdb.APIVersion(RequiredAPIVersion)
	if err != nil {
		return err
	}
	db, err := fdb.OpenDefault()
	if err != nil {
		return err
	}
	store.Database = db
	return nil
}

// Stop stops this memdb.
func (store *Store) Stop() error {
	return nil
}
