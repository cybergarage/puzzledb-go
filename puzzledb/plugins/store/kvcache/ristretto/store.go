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
	store := &Store{
		BaseStore: kvcache.NewBaseStore(),
		Service:   service,
		Cache:     nil,
	}
	store.SetStore(store)
	return store
}

// ServiceType returns the plug-in service type.
func (store *Store) ServiceType() plugins.ServiceType {
	return plugins.StoreKvCacheService
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
