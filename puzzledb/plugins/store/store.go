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

package store

import (
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins"
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
	"github.com/cybergarage/puzzledb-go/puzzledb/store/kv"
)

type Store struct {
	kvStore kv.Store
	document.Coder
	document.KeyCoder
}

// NewStore returns a new store.
func NewStore() *Store {
	return &Store{
		kvStore:  nil,
		Coder:    nil,
		KeyCoder: nil,
	}
}

// NewStoreWith returns a new store with the specified key-value store service.
func NewStoreWith(kvs kv.Store) *Store {
	return &Store{
		kvStore:  kvs,
		Coder:    nil,
		KeyCoder: nil,
	}
}

// SetKvStore sets the key-value store service.
func (store *Store) SetKvStore(kvs kv.Store) {
	store.kvStore = kvs
	store.kvStore.SetKeyCoder(store.KeyCoder)
}

// SetDocumentCoder sets the document coder.
func (store *Store) SetDocumentCoder(coder document.Coder) {
	store.Coder = coder
}

// SetKeyCoder sets the key coder.
func (store *Store) SetKeyCoder(coder document.KeyCoder) {
	store.KeyCoder = coder
	if store.kvStore != nil {
		store.kvStore.SetKeyCoder(coder)
	}
}

// ServiceType returns the plug-in service type.
func (store *Store) ServiceType() plugins.ServiceType {
	return plugins.StoreDocumentService
}

// ServiceName returns the plug-in service name.
func (store *Store) ServiceName() string {
	return "kv"
}

// CreateDatabase creates a new database.
func (store *Store) CreateDatabase(name string) error {
	return store.kvStore.CreateDatabase(name)
}

// GetDatabase retruns the specified database.
func (store *Store) GetDatabase(name string) (store.Database, error) {
	kvDB, err := store.kvStore.GetDatabase(name)
	if err != nil {
		return nil, err
	}
	db := &database{
		kv:       kvDB,
		Coder:    store.Coder,
		KeyCoder: store.KeyCoder,
	}
	return db, nil
}

// RemoveDatabase removes the specified database.
func (store *Store) RemoveDatabase(name string) error {
	return store.kvStore.RemoveDatabase((name))
}

// ListDatabases returns the database list.
func (store *Store) ListDatabases() ([]store.Database, error) {
	return nil, nil
}

// Start starts this store.
func (store *Store) Start() error {
	return nil
}

// Stop stops this store.
func (store *Store) Stop() error {
	return nil
}
