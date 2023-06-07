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
	"github.com/cybergarage/puzzledb-go/puzzledb/store/kv"
)

// CacheConfig represents a key-value cache store configuration.
type CacheConfig struct {
	keyPrefixMap map[byte]bool
}

// NewCacheConfig returns a new key-value cache store configuration.
func NewCacheConfig() *CacheConfig {
	return &CacheConfig{
		keyPrefixMap: map[byte]bool{},
	}
}

// RegisterCacheKeyPrefix registers a key header for the cache store.
func (conf *CacheConfig) RegisterCacheKeyHeader(header kv.KeyHeader) {
	conf.keyPrefixMap[header[0]] = true
}

// UnregisterCacheKeyPrefix unregisters a key header for the cache store.
func (conf *CacheConfig) UnregisterCacheKeyHeader(header kv.KeyHeader) {
	delete(conf.keyPrefixMap, header[0])
}

// IsRegisteredCacheKey returns true if the specified key is registered to the cache store.
func (conf *CacheConfig) IsRegisteredCacheKey(key kv.Key) bool {
	if len(key) == 0 {
		return false
	}
	header, ok := key[0].([]byte)
	if !ok {
		return false
	}
	if len(header) == 0 {
		return false
	}
	_, ok = conf.keyPrefixMap[header[0]]
	return ok
}

// EnableDatabaseCache enables a cache for all databases.
func (conf *CacheConfig) EnableDatabaseCache() {
	conf.RegisterCacheKeyHeader(kv.DatabaseKeyHeader)
}

// EnableCollectionCache enables a cache for all database collections.
func (conf *CacheConfig) EnableCollectionCache() {
	conf.RegisterCacheKeyHeader(kv.CollectionKeyHeader)
}
