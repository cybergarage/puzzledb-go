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

package util

import (
	"bytes"
	"errors"
	"fmt"

	dockv "github.com/cybergarage/puzzledb-go/puzzledb/document/kv"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/store"
	"github.com/cybergarage/puzzledb-go/puzzledb/store/kv"
)

// Store represents a document store.
type Store struct {
	*store.Store
}

// NewStoreWith creates a new store with the given object.
func NewStoreWith(obj any) (*Store, error) {
	store, ok := obj.(*store.Store)
	if !ok {
		return nil, errors.New("failed to get kv store")
	}
	return &Store{
		Store: store,
	}, nil
}

// Dump returns a string representation of the store.
func (store *Store) Dump() ([]string, error) {
	lines := []string{}

	docStore := store.Store
	kvStore := docStore.KvStore()

	tx, err := kvStore.Transact(false)
	if err != nil {
		return lines, nil
	}

	keys := []kv.Key{
		kv.NewKeyWith(kv.DatabaseKeyHeader, kv.Key{}),
		kv.NewKeyWith(kv.CollectionKeyHeader, kv.Key{}),
		kv.NewKeyWith(kv.PrimaryIndexHeader, kv.Key{}),
		kv.NewKeyWith(kv.SecondaryIndexHeader, kv.Key{}),
		kv.NewKeyWith(kv.DocumentKeyHeader, kv.Key{}),
	}

	for _, key := range keys {
		rs, err := tx.GetRange(key)
		if err != nil {
			continue
		}
		for rs.Next() {
			obj, err := rs.Object()
			if err != nil {
				continue
			}
			keys := obj.Key().Elements()
			keyHeaderBytes, ok := keys[0].([]byte)
			if !ok {
				lines = append(lines, fmt.Sprintf("%v: %v", keys[1:], obj.Value()))
			}
			keyHeader := dockv.NewKeyHeaderFrom(keyHeaderBytes)

			switch keyHeader.Type() {
			case dockv.DatabaseObject, dockv.CollectionObject, dockv.DocumentObject:
				r := bytes.NewReader(obj.Value())
				val, err := docStore.DecodeDocument(r)
				if err != nil {
					lines = append(lines, fmt.Sprintf("%v %v: %v", keyHeader, keys[1:], obj.Value()))
					continue
				}
				lines = append(lines, fmt.Sprintf("%v %v: %v", keyHeader, keys[1:], val))
			case dockv.IndexObject:
				idxKeys, err := docStore.DecodeKey(obj.Value())
				if err != nil {
					lines = append(lines, fmt.Sprintf("%v %v: %v", keyHeader, keys[1:], obj.Value()))
					continue
				}
				idxKeyHederBytes, ok := idxKeys[0].([]byte)
				if !ok {
					lines = append(lines, fmt.Sprintf("%v %v: %v", keyHeader, keys[1:], idxKeys))
				}
				idxKeyHeder := dockv.NewKeyHeaderFrom(idxKeyHederBytes)
				lines = append(lines, fmt.Sprintf("%v %v: %v %v", keyHeader, keys[1:], idxKeyHeder, idxKeys[1:]))
			}
		}
	}
	return lines, tx.Commit()
}
