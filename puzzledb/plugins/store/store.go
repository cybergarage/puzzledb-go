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
func (s *Store) SetKvStore(kvs kv.Store) {
	s.kvStore = kvs
	s.kvStore.SetKeyCoder(s.KeyCoder)
}

// SetDocumentCoder sets the document coder.
func (s *Store) SetDocumentCoder(coder document.Coder) {
	s.Coder = coder
}

// SetKeyCoder sets the key coder.
func (s *Store) SetKeyCoder(coder document.KeyCoder) {
	s.KeyCoder = coder
	if s.kvStore != nil {
		s.kvStore.SetKeyCoder(coder)
	}
}

// ServiceType returns the plug-in service type.
func (s *Store) ServiceType() plugins.ServiceType {
	return plugins.StoreDocumentService
}

// ServiceName returns the plug-in service name.
func (s *Store) ServiceName() string {
	return "kv"
}

// CreateDatabase creates a new database.
func (s *Store) CreateDatabase(name string) error {
	return s.kvStore.CreateDatabase(name)
}

// GetDatabase retruns the specified database.
func (s *Store) GetDatabase(name string) (store.Database, error) {
	kvDB, err := s.kvStore.GetDatabase(name)
	if err != nil {
		return nil, err
	}
	db := &database{
		kv:       kvDB,
		Coder:    s.Coder,
		KeyCoder: s.KeyCoder,
	}
	return db, nil
}

// RemoveDatabase removes the specified database.
func (s *Store) RemoveDatabase(name string) error {
	return s.kvStore.RemoveDatabase((name))
}

// ListDatabases returns the all databases.
func (s *Store) ListDatabases() ([]store.Database, error) {
	dbs := make([]store.Database, 0)
	// dbs := make([]store.Database, len(kvDB))
	// for n, kvDB := range kvDB {
	// 	dbs[n] = &database{
	// 		kv:       kvDB,
	// 		Coder:    s.Coder,
	// 		KeyCoder: s.KeyCoder,
	// 	}
	// }
	return dbs, nil
}

// Start starts this store.
func (s *Store) Start() error {
	return nil
}

// Stop stops this store.
func (s *Store) Stop() error {
	return nil
}
