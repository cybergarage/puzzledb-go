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
	"fmt"

	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/store"
	"github.com/cybergarage/puzzledb-go/puzzledb/store/kv"
)

type Store struct {
	store.Service
}

func NewStoreWith(service store.Service) *Store {
	return &Store{
		Service: service,
	}
}

func (s *Store) String() string {
	out := ""

	docStore, ok := s.Service.(*store.Store)
	if !ok {
		return out
	}

	kvStore := docStore.KvStore()
	tx, err := kvStore.Transact(false)
	if err != nil {
		return out
	}

	keys := []kv.Key{
		kv.NewKeyWith(kv.DatabaseKeyHeader, kv.Key{}),
		kv.NewKeyWith(kv.SchemaKeyHeader, kv.Key{}),
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
			obj := rs.Object()
			out += fmt.Sprintf("%v\n", obj)
		}
	}

	err = tx.Commit()
	if err != nil {
		return out
	}

	return out
}
