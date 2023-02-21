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
	// "github.com/apple/foundationdb/bindings/go/src/fdb"
	store "github.com/cybergarage/puzzledb-go/puzzledb/server/plugins/store/kv"
	"github.com/cybergarage/puzzledb-go/puzzledb/store/kv"
)

// FoundationDB represents a FoundationDB instance.
type FoundationDB struct {
}

// New returns a new memdb store instance.
func NewStore() store.Service {
	// fdb.MustAPIVersion(720)
	return &FoundationDB{}
}

// CreateDatabase creates a new database.
func (store *FoundationDB) CreateDatabase(name string) error {
	return nil
}

// GetDatabase retruns the specified database.
func (store *FoundationDB) GetDatabase(id string) (kv.Database, error) {
	return newDatabaseWithID(id), nil
}

// Start starts this memdb.
func (store *FoundationDB) Start() error {
	return nil
}

// Stop stops this memdb.
func (store *FoundationDB) Stop() error {
	return nil
}
