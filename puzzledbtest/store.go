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

package puzzledbtest

import (
	"bytes"
	"fmt"
	"strings"

	dockv "github.com/cybergarage/puzzledb-go/puzzledb/document/kv"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/store"
	"github.com/cybergarage/puzzledb-go/puzzledb/store/kv"
)

// Store represents a document store.
type Store struct {
	store.Service
}

// NewStoreWith creates a new store with the given object.
func NewStoreWith(service store.Service) *Store {
	return &Store{
		Service: service,
	}
}

// Dump returns a string representation of the store.
func (s *Store) Dump() ([]string, error) {
	lines := []string{}

	docStore, ok := s.Service.(*store.Store)
	if !ok {
		return lines, fmt.Errorf("invalid store")
	}

	kvStore := docStore.KvStore()
	tx, err := kvStore.Transact(false)
	if err != nil {
		return lines, err
	}

	keys := []kv.Key{
		kv.NewKeyWith(kv.DatabaseKeyHeader, kv.Key{}),
		kv.NewKeyWith(kv.CollectionKeyHeader, kv.Key{}),
		kv.NewKeyWith(kv.IndexKeyHeader, kv.Key{}),
		kv.NewKeyWith(kv.DocumentKeyHeader, kv.Key{}),
	}

	for _, key := range keys {
		rs, err := tx.Scan(key)
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
				lines = append(lines, fmt.Sprintf("%v %v:", keyHeader, keys[1:]))
			}
		}
		if err := rs.Err(); err != nil {
			return lines, err
		}
		if err := rs.Close(); err != nil {
			return lines, err
		}
	}
	return lines, tx.Commit()
}

// String returns a string representation of the store.
func (s *Store) String() string {
	lines, err := s.Dump()
	if err != nil {
		return ""
	}
	return strings.Join(lines, "\n")
}
