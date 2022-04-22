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

package memdb

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"testing"

	"github.com/cybergarage/mimicdb/mimicdb/plugins/store"
)

const (
	testKeyCount = 100
)

func TestStores(t *testing.T) {
	stores := []store.Store{
		NewStore(),
	}

	for _, store := range stores {
		testStore(t, store)
	}
}

func testStore(t *testing.T, store store.Store) {
	if err := store.Start(); err != nil {
		t.Error(err)
	}
	if err := store.Open("testdb"); err != nil {
		t.Error(err)
	}

	keys := make([][]byte, testKeyCount)
	vals := make([][]byte, testKeyCount)
	for n := 0; n < testKeyCount; n++ {
		keys[n] = make([]byte, 8)
		binary.LittleEndian.PutUint64(keys[n], rand.Uint64())
		vals[n] = make([]byte, 8)
		binary.LittleEndian.PutUint64(vals[n], rand.Uint64())
	}

	// Insert test

	for n, key := range keys {
		tx, err := store.Transact(true)
		if err != nil {
			t.Error(err)
			break
		}
		val := vals[n]
		obj := &Object{
			Key:   string(key),
			Value: val,
		}
		if err := tx.Insert(obj); err != nil {
			t.Error(err)
			break
		}
		if err := tx.Commit(); err != nil {
			t.Error(err)
			break
		}
	}

	// Select test

	for n, key := range keys {
		tx, err := store.Transact(false)
		if err != nil {
			t.Error(err)
			break
		}
		obj, err := tx.Select(string(key))
		if err != nil {
			t.Error(err)
			break
		}
		if bytes.Compare(obj.Value, vals[n]) != 0 {
			t.Errorf("%s != %s", obj.Value, vals[n])
		}
		if err := tx.Cancel(); err != nil {
			t.Error(err)
			break
		}
	}

	if err := store.Close(); err != nil {
		t.Error(err)
	}
	if err := store.Stop(); err != nil {
		t.Error(err)
	}
}
