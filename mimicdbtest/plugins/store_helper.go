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

	"github.com/cybergarage/mimicdb/mimicdb/plugins/store"
)

const (
	testKeyCount  = 100
	testValBufMax = 8
)

//nolint:gosec,cyclop
func StoreTest(t *testing.T, s store.Store) {
	t.Helper()

	if err := s.Start(); err != nil {
		t.Error(err)
	}
	if err := s.Open("testdb"); err != nil {
		t.Error(err)
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
		tx, err := s.Transact(true)
		if err != nil {
			t.Error(err)
			break
		}
		val := vals[n]
		obj := &store.Object{
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
		tx, err := s.Transact(false)
		if err != nil {
			t.Error(err)
			break
		}
		obj, err := tx.Select(string(key))
		if err != nil {
			t.Error(err)
			break
		}
		if !bytes.Equal(obj.Value, vals[n]) {
			t.Errorf("%s != %s", obj.Value, vals[n])
		}
		if err := tx.Cancel(); err != nil {
			t.Error(err)
			break
		}
	}

	if err := s.Close(); err != nil {
		t.Error(err)
	}
	if err := s.Stop(); err != nil {
		t.Error(err)
	}
}
