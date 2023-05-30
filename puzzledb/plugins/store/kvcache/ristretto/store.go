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

package ristretto

import (
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/store/kv"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/store/kvcache"
	"github.com/dgraph-io/ristretto"
)

// Store represents a cache store service instance.
type Store struct {
	*kvcache.BaseStore
	kv.Service
	Cache *ristretto.Cache
}

// NewStore returns a new FoundationDB store instance.
func NewStore() kv.Service {
	return NewStoreWith(nil)
}

// NewStoreWith returns a new FoundationDB store instance with the specified key coder.
func NewStoreWith(service kv.Service) kv.Service {
	return &Store{
		BaseStore: kvcache.NewBaseStore(),
		Service:   service,
		Cache:     nil,
	}
}

// ServiceType returns the plug-in service type.
func (store *Store) ServiceType() plugins.ServiceType {
	return plugins.StoreKvCacheService
}

// ServiceName returns the plug-in service name.
func (store *Store) ServiceName() string {
	return "ristretto"
}

// Start starts the ristretto store.
func (store *Store) Start() error {
	return nil
}

// Stop stops the ristretto store.
func (store Store) Stop() error {
	return nil
}
