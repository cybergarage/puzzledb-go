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

package plugins

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"testing"

	plugins "github.com/cybergarage/puzzledb-go/puzzledb/server/plugins/store/kv"
	store "github.com/cybergarage/puzzledb-go/puzzledb/store/kv"
)

const (
	testDbName    = "testdb"
	testKeyCount  = 100
	testValBufMax = 8
)

//nolint:gosec,cyclop
func StoreTest(t *testing.T, s plugins.Service) {
	t.Helper()

	if err := s.Start(); err != nil {
		t.Error(err)
		return
	}
	if err := s.CreateDatabase(testDbName); err != nil {
		t.Error(err)
		return
	}
	db, err := s.GetDatabase(testDbName)
	if err != nil {
		t.Error(err)
		return
	}

	keys := make([][]byte, testKeyCount)
	vals := make([][]byte, testKeyCount)
	for n := 0; n < testKeyCount; n++ {
		keys[n] = make([]byte, testValBufMax)
		binary.LittleEndian.PutUint64(keys[n], rand.Uint64())
		vals[n] = make([]byte, testValBufMax)
		binary.LittleEndian.PutUint64(vals[n], rand.Uint64())
	}

	// Insert test

	for n, key := range keys {
		tx, err := db.Transact(true)
		if err != nil {
			t.Error(err)
			break
		}
		val := vals[n]
		obj := &store.Object{
			Key:   []any{key},
			Value: val,
		}
		if err := tx.Set(obj); err != nil {
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
		tx, err := db.Transact(false)
		if err != nil {
			t.Error(err)
			break
		}
		rs, err := tx.Get([]any{key})
		if err != nil {
			t.Error(err)
			break
		}
		if !rs.Next() {
			t.Errorf("%v != 1", rs)
			break
		}
		obj := rs.Object()
		if !bytes.Equal(obj.Value, vals[n]) {
			t.Errorf("%s != %s", obj.Value, vals[n])
		}
		if err := tx.Cancel(); err != nil {
			t.Error(err)
			break
		}
	}

	if err := s.Stop(); err != nil {
		t.Error(err)
		return
	}
}
