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
	"sync/atomic"

	"github.com/cybergarage/puzzledb-go/puzzledb/plugins"
	"github.com/cybergarage/puzzledb-go/puzzledb/store/kv"
)

// BaseStore represents a cache store service instance.
type BaseStore struct {
	plugins.Config
	kv.Store
	*CacheConfig
	reqCnt int64
	hitCnt int64
}

// NewStore returns a new FoundationDB store instance.
func NewBaseStore() *BaseStore {
	return &BaseStore{
		Config:      plugins.NewConfig(),
		Store:       nil,
		CacheConfig: NewCacheConfig(),
		reqCnt:      0,
		hitCnt:      0,
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

// IncrementRequestCount increments the number of cache requests.
func (store *BaseStore) IncrementRequestCount() {
	atomic.AddInt64(&store.reqCnt, 1)
}

// IncrementHitCount increments the number of cache hits.
func (store *BaseStore) IncrementHitCount() {
	atomic.AddInt64(&store.hitCnt, 1)
}

// CacheRequestCount returns the number of cache requests.
func (store *BaseStore) CacheRequestCount() int64 {
	return store.reqCnt
}

// CacheMissCount returns the number of cache misses.
func (store *BaseStore) CacheHitCount() int64 {
	return store.hitCnt
}

// CacheMissCount returns the number of cache misses.
func (store *BaseStore) CacheHitRate() float64 {
	return float64(store.hitCnt) / float64(store.reqCnt)
}

// CacheMissCount returns the number of cache misses.
func (store *BaseStore) CacheMissCount() int64 {
	return store.reqCnt - store.hitCnt
}

// CacheMissRate returns the cache miss rate.
func (store *BaseStore) CacheMissRate() float64 {
	return 1.0 - store.CacheHitRate()
}
