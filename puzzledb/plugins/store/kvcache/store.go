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

package kvcache

import (
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins"
	"github.com/cybergarage/puzzledb-go/puzzledb/store/kv"
)

// Store represents a key-value cache store interface.
type Store interface {
	kv.Store
	plugins.Service
	// SetStore sets the key-value store.
	SetStore(s kv.Store)
}

// BaseStore represents a cache store service instance.
type BaseStore struct {
	kv.Store
}

// NewStore returns a new FoundationDB store instance.
func NewBaseStore() *BaseStore {
	return &BaseStore{
		Store: nil,
	}
}

// SetStore sets the key-value store service.
func (store *BaseStore) SetStore(s kv.Store) {
	store.Store = s
}

// ServiceType returns the plug-in service type.
func (store *BaseStore) ServiceType() plugins.ServiceType {
	return plugins.StoreKvCacheService
}
