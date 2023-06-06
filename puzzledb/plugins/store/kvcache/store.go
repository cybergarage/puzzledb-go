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

// CacheStore represents a key-value cache store interface.
type CacheStore interface {
	kv.Store
	// SetStore sets a base key-value store.
	SetStore(s kv.Store)
	// RegisterCacheKeyPrefix registers a key header for the cache store.
	RegisterCacheKeyHeader(header kv.KeyHeader)
	// IsRegisteredCacheKey returns true if the specified key is registered to the cache store.
	IsRegisteredCacheKey(key kv.Key) bool
	// EraseCache deletes a cache for the specified key.
	EraseCache(key kv.Key) error
}

// Service represents a key-value cache store service interface.
type Service interface {
	CacheStore
	plugins.Service
}
