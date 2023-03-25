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
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/store/kv"
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
)

type Store struct {
	kvService kv.Service
	document.Serializer
}

func NewStoreWithKvStore(kvs kv.Service) *Store {
	return &Store{
		kvService:  kvs,
		Serializer: nil,
	}
}

func (store *Store) SetSerializer(serializer document.Serializer) {
	store.Serializer = serializer
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
		kv:         kvDB,
		Serializer: store.Serializer,
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
