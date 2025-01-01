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

package fdb

import (
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins"
	kvPlugin "github.com/cybergarage/puzzledb-go/puzzledb/plugins/store/kv"
	kvStore "github.com/cybergarage/puzzledb-go/puzzledb/store/kv"
)

const (
	RequiredAPIVersion int = 710
	ClusterFile            = "cluster_file"
)

// Store represents a FoundationDB store service instance.
type Store struct {
	*kvPlugin.BaseStore
	fdb.Database
}

// NewStore returns a new FoundationDB store instance.
func NewStore() kvPlugin.Service {
	return NewStoreWith(nil)
}

// NewStoreWith returns a new FoundationDB store instance with the specified key coder.
func NewStoreWith(coder document.KeyCoder) kvPlugin.Service {
	return &Store{ //nolint:all
		BaseStore: kvPlugin.NewBaseStoreWith(coder),
	}
}

// ServiceType returns the plug-in service type.
func (store *Store) ServiceType() plugins.ServiceType {
	return plugins.StoreKvService
}

// ServiceName returns the plug-in service name.
func (store *Store) ServiceName() string {
	return "fdb"
}

// Transact begin a new transaction.
func (store *Store) Transact(write bool) (kvStore.Transaction, error) {
	txn, err := store.Database.CreateTransaction()
	if err != nil {
		return nil, err
	}
	err = txn.Options().SetAccessSystemKeys()
	if err != nil {
		return nil, err
	}
	return newTransaction(txn, store.KeyCoder), nil
}

// GetClusterFile returns the cluster file configuration.
func (store *Store) GetClusterFile() (string, error) {
	e, err := store.LookupServiceConfigString(store, ClusterFile)
	if err != nil {
		return "", err
	}
	return e, nil
}

// Start starts this memdb.
func (store *Store) Start() error {
	err := fdb.APIVersion(RequiredAPIVersion)
	if err != nil {
		return err
	}

	clusterFile, err := store.GetClusterFile()

	var db fdb.Database
	if err == nil || 0 < len(clusterFile) {
		db, err = fdb.OpenDatabase(clusterFile)
	} else {
		db, err = fdb.OpenDefault()
	}

	if err != nil {
		return err
	}
	store.Database = db
	return nil
}

// Stop stops this memdb.
func (store *Store) Stop() error {
	return nil
}
