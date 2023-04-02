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
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins"
	store "github.com/cybergarage/puzzledb-go/puzzledb/plugins/store/kv"
)

// Store represents a Memdb store service instance.
type Store struct {
	*Databases
}

// New returns a new memdb store instance.
func NewStore() store.Service {
	return &Store{
		Databases: NewDatabases(),
	}
}

// Type returns the plug-in service type.
func (store *Store) Type() plugins.ServiceType {
	return plugins.KvStoreService
}

// Name returns the plug-in service name.
func (store *Store) Name() string {
	return "memdb"
}

// Start starts this memdb.
func (store *Store) Start() error {
	return nil
}

// Stop stops this memdb.
func (store Store) Stop() error {
	return nil
}
