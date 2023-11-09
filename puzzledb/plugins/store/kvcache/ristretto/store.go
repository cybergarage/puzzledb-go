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
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
	pluginkv "github.com/cybergarage/puzzledb-go/puzzledb/plugins/store/kv"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/store/kvcache"
	"github.com/cybergarage/puzzledb-go/puzzledb/store/kv"
	"github.com/dgraph-io/ristretto"
)

// Store represents a cache store service instance.
type Store struct {
	*kvcache.BaseStore
	Cache *ristretto.Cache
}

// NewStore returns a new FoundationDB store instance.
func NewStore() pluginkv.Service {
	return NewStoreWith(nil)
}

// NewStoreWith returns a new FoundationDB store instance with the specified key coder.
func NewStoreWith(kvStore kv.Store) pluginkv.Service {
	store := &Store{
		BaseStore: kvcache.NewBaseStore(),
		Cache:     nil,
	}
	store.BaseStore.SetStore(kvStore)
	return store
}

// ServiceName returns the plug-in service name.
func (store *Store) ServiceName() string {
	return "ristretto"
}

func (store *Store) GetNumCounters() (int64, error) {
	v, err := store.GetServiceConfigInt(store, NumCounters)
	if err != nil {
		return DefaultNumCounters, nil //nolint: nilerr
	}
	return int64(v), err
}

func (store *Store) GetMaxCost() (int64, error) {
	v, err := store.GetServiceConfigInt(store, MaxCost)
	if err != nil {
		return DefaultMaxCost, nil //nolint: nilerr
	}
	return int64(v), err
}

func (store *Store) GeBufferItems() (int64, error) {
	v, err := store.GetServiceConfigInt(store, BufferItems)
	if err != nil {
		return DefaultBufferItems, nil //nolint: nilerr
	}
	return int64(v), err
}

func (store *Store) GeMetrics() (bool, error) {
	v, err := store.GetServiceConfigBool(store, Metrics)
	if err != nil {
		return DefalutMetrics, nil //nolint: nilerr
	}
	return v, err
}

// Transact begin a new transaction.
func (store *Store) Transact(write bool) (kv.Transaction, error) {
	kvStore := store.BaseStore.Store
	txn, err := kvStore.Transact(write)
	if err != nil {
		return nil, err
	}
	return NewTransaction(txn, store), nil
}

// SetCache sets a cache for the specified key.
func (store *Store) SetCache(obj *kv.Object) error {
	key := obj.Key
	b, err := store.EncodeKey(key)
	if err != nil {
		return err
	}
	if !store.Cache.Set(b, obj.Value, 0) {
		return kvcache.NewErrSetCache(obj.Key)
	}
	return nil
}

// EraseCache deletes a cache for the specified key.
func (store *Store) EraseCache(key kv.Key) error {
	b, err := store.EncodeKey(key)
	if err != nil {
		return err
	}
	store.Cache.Del(b)
	return nil
}

// EraseDatabaseCache deletes a cache for the specified database.
func (store *Store) EraseDatabaseCache(database string) error {
	return store.EraseCache(kv.NewKeyWith(kv.DatabaseKeyHeader, document.NewKeyWith(database)))
}

// EraseCollectionCache deletes a cache for the specified collection.
func (store *Store) EraseCollectionCache(database string, collection string) error {
	return store.EraseCache(kv.NewKeyWith(kv.CollectionKeyHeader, document.NewKeyWith(database, collection)))
}

// Start starts the ristretto store.
func (store *Store) Start() error {
	numCounters, err := store.GetNumCounters()
	if err != nil {
		return err
	}
	maxCost, err := store.GetMaxCost()
	if err != nil {
		return err
	}
	bufferItems, err := store.GeBufferItems()
	if err != nil {
		return err
	}
	metrics, err := store.GeMetrics()
	if err != nil {
		return err
	}
	conf := &ristretto.Config{
		NumCounters:        numCounters,
		MaxCost:            maxCost,
		BufferItems:        bufferItems,
		Metrics:            metrics,
		OnEvict:            nil,
		OnReject:           nil,
		OnExit:             nil,
		KeyToHash:          nil,
		Cost:               nil,
		IgnoreInternalCost: false,
	}
	cache, err := ristretto.NewCache(conf)
	if err != nil {
		return err
	}
	store.Cache = cache
	return nil
}

// Stop stops the ristretto store.
func (store Store) Stop() error {
	return nil
}
