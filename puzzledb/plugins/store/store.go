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
	"bytes"

	"github.com/cybergarage/puzzledb-go/puzzledb/config"
	"github.com/cybergarage/puzzledb-go/puzzledb/context"
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins"
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
	"github.com/cybergarage/puzzledb-go/puzzledb/store/kv"
)

type Store struct {
	kvStore kv.Store
	document.Coder
	document.KeyCoder
}

// NewStore returns a new store.
func NewStore() *Store {
	return &Store{
		kvStore:  nil,
		Coder:    nil,
		KeyCoder: nil,
	}
}

// NewStoreWith returns a new store with the specified key-value store service.
func NewStoreWith(kvs kv.Store) *Store {
	return &Store{
		kvStore:  kvs,
		Coder:    nil,
		KeyCoder: nil,
	}
}

// SetConfig sets the specified configuration.
func (s *Store) SetConfig(conf config.Config) {
}

// SetKvStore sets the key-value store service.
func (s *Store) SetKvStore(kvs kv.Store) {
	s.kvStore = kvs
	s.kvStore.SetKeyCoder(s.KeyCoder)
}

// KvStore returns the key-value store.
func (s *Store) KvStore() kv.Store {
	return s.kvStore
}

// SetDocumentCoder sets the document coder.
func (s *Store) SetDocumentCoder(coder document.Coder) {
	s.Coder = coder
}

// SetKeyCoder sets the key coder.
func (s *Store) SetKeyCoder(coder document.KeyCoder) {
	s.KeyCoder = coder
	if s.kvStore != nil {
		s.kvStore.SetKeyCoder(coder)
	}
}

// ServiceType returns the plug-in service type.
func (s *Store) ServiceType() plugins.ServiceType {
	return plugins.StoreDocumentService
}

// ServiceName returns the plug-in service name.
func (s *Store) ServiceName() string {
	return "dockv"
}

// CreateDatabase creates a new database.
func (s *Store) CreateDatabase(ctx context.Context, name string) error {
	ctx.StartSpan("CreateDatabase")
	defer ctx.FinishSpan()

	txn, err := s.kvStore.Transact(true)
	if err != nil {
		return err
	}

	kvDBKey := kv.NewKeyWith(kv.DatabaseKeyHeader, document.Key{name})

	kvDBObj, err := txn.Get(kvDBKey)
	if err == nil && kvDBObj != nil {
		if err := txn.Cancel(); err != nil {
			return err
		}
		return store.NewDatabaseExistError(name)
	}

	var opts store.DatabaseOptions
	var objBytes bytes.Buffer
	err = s.EncodeDocument(&objBytes, opts)
	if err != nil {
		if err := txn.Cancel(); err != nil {
			return err
		}
		return err
	}

	kvObj := kv.Object{
		Key:   kvDBKey,
		Value: objBytes.Bytes(),
	}
	err = txn.Set(&kvObj)
	if err != nil {
		if err := txn.Cancel(); err != nil {
			return err
		}
		return err
	}
	err = txn.Commit()
	if err != nil {
		if err := txn.Cancel(); err != nil {
			return err
		}
		return err
	}

	return nil
}

// GetDatabase retruns the specified database.
func (s *Store) GetDatabase(ctx context.Context, name string) (store.Database, error) {
	ctx.StartSpan("GetDatabase")
	defer ctx.FinishSpan()

	txn, err := s.kvStore.Transact(false)
	if err != nil {
		return nil, err
	}
	kvDBKey := kv.NewKeyWith(kv.DatabaseKeyHeader, document.Key{name})

	kvDBObj, err := txn.Get(kvDBKey)
	if err != nil && kvDBObj == nil {
		if err := txn.Cancel(); err != nil {
			return nil, err
		}
		return nil, store.NewDatabaseNotExistError(name)
	}
	err = txn.Commit()
	if err != nil {
		if err := txn.Cancel(); err != nil {
			return nil, err
		}
		return nil, err
	}

	dbOptsObj, err := s.DecodeDocument(bytes.NewReader(kvDBObj.Value))
	if err != nil {
		return nil, err
	}

	dbOpts, err := newDatabaeOptionsFrom(dbOptsObj)
	if err != nil {
		return nil, err
	}

	db := &database{
		name:            name,
		DatabaseOptions: dbOpts,
		Store:           s.kvStore,
		Coder:           s.Coder,
		KeyCoder:        s.KeyCoder,
	}

	return db, nil
}

// RemoveDatabase removes the specified database.
func (s *Store) RemoveDatabase(ctx context.Context, name string) error {
	ctx.StartSpan("RemoveDatabase")
	defer ctx.FinishSpan()

	db, err := s.GetDatabase(ctx, name)
	if err != nil {
		return err
	}

	// Remove all database objects

	dbTxn, err := db.Transact(true)
	if err != nil {
		return err
	}

	err = dbTxn.TruncateDocuments(ctx)
	if err != nil {
		if err := dbTxn.Cancel(ctx); err != nil {
			return err
		}
		return err
	}

	err = dbTxn.TruncateIndexes(ctx)
	if err != nil {
		if err := dbTxn.Cancel(ctx); err != nil {
			return err
		}
		return err
	}

	err = dbTxn.TruncateCollections(ctx)
	if err != nil {
		if err := dbTxn.Cancel(ctx); err != nil {
			return err
		}
		return err
	}

	err = dbTxn.Commit(ctx)
	if err != nil {
		if err := dbTxn.Cancel(ctx); err != nil {
			return err
		}
		return err
	}

	// Remove database objects

	kvDBKey := kv.NewKeyWith(kv.DatabaseKeyHeader, document.Key{name})

	txn, err := s.kvStore.Transact(true)
	if err != nil {
		return err
	}

	err = txn.Remove(kvDBKey)
	if err != nil {
		if err := txn.Cancel(); err != nil {
			return err
		}
		return err
	}
	err = txn.Commit()
	if err != nil {
		if err := txn.Cancel(); err != nil {
			return err
		}
		return err
	}

	return nil
}

// ListDatabases returns the all databases.
func (s *Store) ListDatabases(ctx context.Context) ([]store.Database, error) {
	ctx.StartSpan("ListDatabases")
	defer ctx.FinishSpan()

	txn, err := s.kvStore.Transact(false)
	if err != nil {
		return nil, err
	}

	kvDBKey := kv.NewKeyWith(kv.DatabaseKeyHeader, document.Key{})
	kvRs, err := txn.GetRange(kvDBKey)
	if err != nil {
		if err := txn.Cancel(); err != nil {
			return nil, err
		}
		return nil, err
	}

	dbs := make([]store.Database, 0)
	for kvRs.Next() {
		kvObj := kvRs.Object()
		kvKeys := kvObj.Key.Elements()
		kvKeyLen := len(kvKeys)
		if kvKeyLen == 0 {
			continue
		}
		lastKvKey := kvKeys[kvKeyLen-1]
		name, ok := lastKvKey.(string)
		if !ok {
			continue
		}

		dbOptsObj, err := s.DecodeDocument(bytes.NewReader(kvObj.Value))
		if err != nil {
			return nil, err
		}

		dbOpts, err := newDatabaeOptionsFrom(dbOptsObj)
		if err != nil {
			return nil, err
		}

		db := &database{
			name:            name,
			DatabaseOptions: dbOpts,
			Store:           s.kvStore,
			Coder:           s.Coder,
			KeyCoder:        s.KeyCoder,
		}
		dbs = append(dbs, db)
	}

	err = txn.Commit()
	if err != nil {
		if err := txn.Cancel(); err != nil {
			return nil, err
		}
		return nil, err
	}

	return dbs, nil
}

// Start starts this store.
func (s *Store) Start() error {
	s.kvStore.SetKeyCoder(s.KeyCoder)
	return nil
}

// Stop stops this store.
func (s *Store) Stop() error {
	return nil
}
