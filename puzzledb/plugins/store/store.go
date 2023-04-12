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
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/store/kv"
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
)

type Store struct {
	kvService kv.Service
	document.Coder
}

func NewStoreWithKvStore(kvs kv.Service) *Store {
	return &Store{
		kvService: kvs,
		Coder:     nil,
	}
}

// SetDocumentCoder sets the document coder.
func (store *Store) SetDocumentCoder(coder document.Coder) {
	store.Coder = coder
}

// ServiceType returns the plug-in service type.
func (store *Store) ServiceType() plugins.ServiceType {
	return plugins.StoreDocumentService
}

// ServiceName returns the plug-in service name.
func (store *Store) ServiceName() string {
	return "document"
}

// CreateDatabase creates a new database.
func (store *Store) CreateDatabase(name string) error {
	return store.kvService.CreateDatabase(name)
}

// GetDatabase retruns the specified database.
func (store *Store) GetDatabase(name string) (store.Database, error) {
	kvDB, err := store.kvService.GetDatabase(name)
	if err != nil {
		return nil, err
	}
	db := &database{
		kv:    kvDB,
		Coder: store.Coder,
	}
	return db, nil
}

// RemoveDatabase removes the specified database.
func (store *Store) RemoveDatabase(name string) error {
	return store.kvService.RemoveDatabase((name))
}

// Start starts this store.
func (store *Store) Start() error {
	return store.kvService.Start()
}

// Stop stops this store.
func (store *Store) Stop() error {
	return store.kvService.Stop()
}
