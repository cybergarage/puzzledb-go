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
	"github.com/cybergarage/puzzledb-go/puzzledb/server/plugins/store/kv"
	store "github.com/cybergarage/puzzledb-go/puzzledb/store/kv"
)

// FoundationDB represents a FoundationDB instance.
type FoundationDB struct {
}

// New returns a new memdb store instance.
func NewStore() kv.Service {
	return &FoundationDB{}
}

// CreateDatabase creates a new database.
func (fdb *FoundationDB) CreateDatabase(name string) error {
	return nil
}

// GetDatabase retruns the specified database.
func (fdb *FoundationDB) GetDatabase(id string) (store.Database, error) {
	return nil, nil
}

// Start starts this memdb.
func (fdb *FoundationDB) Start() error {
	return nil
}

// Stop stops this memdb.
func (fdb FoundationDB) Stop() error {
	return nil
}
