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

package kv

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/cybergarage/puzzledb-go/puzzledb/document"
	plugins "github.com/cybergarage/puzzledb-go/puzzledb/plugins/store/kv"
	store "github.com/cybergarage/puzzledb-go/puzzledb/store/kv"
)

const (
	testKeyCount  = 10
	testValBufMax = 8
)

//nolint:gosec,cyclop,gocognit,gocyclo,maintidx
func StoreTest(t *testing.T, kvStore plugins.Service) {
	t.Helper()

	testKeyPrefix := fmt.Sprintf("testkv%d", time.Now().UnixNano())

	// Starts the key-value store service.

	if err := kvStore.Start(); err != nil {
		t.Error(err)
		return
	}

	defer func() {
		if err := kvStore.Stop(); err != nil {
			t.Error(err)
			return
		}
	}()

	// Generates test keys and values.

	keys := make([]document.Key, testKeyCount)
	vals := make([][]byte, testKeyCount)
	for n := 0; n < testKeyCount; n++ {
		keys[n] = document.NewKeyWith(testKeyPrefix, fmt.Sprintf("key%d", n))
		vals[n] = make([]byte, testValBufMax)
		binary.LittleEndian.PutUint64(vals[n], rand.Uint64())
	}

	cancel := func(t *testing.T, txn store.Transaction) {
		t.Helper()
		if err := txn.Cancel(); err != nil {
			t.Error(err)
		}
	}

	// Inserts test keys and values.

	for n, key := range keys {
		tx, err := kvStore.Transact(true)
		if err != nil {
			t.Error(err)
			return
		}
		obj := &store.Object{
			Key:   key,
			Value: vals[n],
		}
		if err := tx.Set(obj); err != nil {
			cancel(t, tx)
			t.Error(err)
			return
		}
		if err := tx.Commit(); err != nil {
			t.Error(err)
			return
		}
	}

	// Selects inserted test keys.

	for n, key := range keys {
		tx, err := kvStore.Transact(false)
		if err != nil {
			t.Error(err)
			return
		}
		obj, err := tx.Get(key)
		if err != nil {
			cancel(t, tx)
			t.Error(err)
			return
		}
		if !bytes.Equal(obj.Value, vals[n]) {
			cancel(t, tx)
			t.Errorf("%s != %s", obj.Value, vals[n])
		}
		if err := tx.Commit(); err != nil {
			t.Error(err)
			return
		}
	}

	// Selects inserted test keys by range

	for n, key := range keys {
		tx, err := kvStore.Transact(false)
		if err != nil {
			t.Error(err)
			return
		}
		rs, err := tx.GetRange(key)
		if err != nil {
			cancel(t, tx)
			t.Error(err)
			return
		}
		if !rs.Next() {
			cancel(t, tx)
			t.Errorf("key (%v) is not found", key)
			return
		}
		obj := rs.Object()
		if !bytes.Equal(obj.Value, vals[n]) {
			cancel(t, tx)
			t.Errorf("%s != %s", obj.Value, vals[n])
			return
		}
		if rs.Next() {
			cancel(t, tx)
			t.Errorf("other key (%v) is found", key)
			return
		}
		if err := tx.Commit(); err != nil {
			t.Error(err)
			return
		}
	}

	// Updates inserted test values.

	for n := 0; n < testKeyCount; n++ {
		binary.LittleEndian.PutUint64(vals[n], rand.Uint64())
	}

	for n, key := range keys {
		tx, err := kvStore.Transact(true)
		if err != nil {
			t.Error(err)
			return
		}
		obj := &store.Object{
			Key:   key,
			Value: vals[n],
		}
		if err := tx.Set(obj); err != nil {
			cancel(t, tx)
			t.Error(err)
			return
		}
		if err := tx.Commit(); err != nil {
			t.Error(err)
			return
		}
	}

	// Selects updated test keys.

	for n, key := range keys {
		tx, err := kvStore.Transact(false)
		if err != nil {
			t.Error(err)
			return
		}
		obj, err := tx.Get(key)
		if err != nil {
			cancel(t, tx)
			t.Error(err)
			return
		}
		if !bytes.Equal(obj.Value, vals[n]) {
			cancel(t, tx)
			t.Errorf("%s != %s", obj.Value, vals[n])
			return
		}
		if err := tx.Commit(); err != nil {
			t.Error(err)
			return
		}
	}

	// Selects updated test keys by range.

	for n, key := range keys {
		tx, err := kvStore.Transact(false)
		if err != nil {
			t.Error(err)
			return
		}
		rs, err := tx.GetRange(key)
		if err != nil {
			cancel(t, tx)
			t.Error(err)
			return
		}
		if !rs.Next() {
			cancel(t, tx)
			t.Errorf("key (%v) is not found", key)
			return
		}
		obj := rs.Object()
		if !bytes.Equal(obj.Value, vals[n]) {
			cancel(t, tx)
			t.Errorf("%s != %s", obj.Value, vals[n])
			return
		}
		if rs.Next() {
			cancel(t, tx)
			t.Errorf("other key (%v) is found", key)
			return
		}
		if err := tx.Commit(); err != nil {
			t.Error(err)
			return
		}
	}

	// Removes inserted test keys.

	for _, key := range keys {
		tx, err := kvStore.Transact(true)
		if err != nil {
			t.Error(err)
			return
		}
		err = tx.Remove(key)
		if err != nil {
			cancel(t, tx)
			t.Error(err)
			return
		}
		if err := tx.Commit(); err != nil {
			t.Error(err)
			return
		}
	}

	// Selects removed test keys.

	for _, key := range keys {
		tx, err := kvStore.Transact(false)
		if err != nil {
			t.Error(err)
			return
		}
		_, err = tx.Get(key)
		if err == nil {
			t.Errorf("key (%v) is found", key)
			cancel(t, tx)
			t.Error(err)
			return
		}
		if !errors.Is(err, store.ErrNotExist) {
			t.Errorf("key (%v): %s", key, err.Error())
			cancel(t, tx)
			t.Error(err)
			return
		}
		if err := tx.Commit(); err != nil {
			t.Error(err)
			return
		}
	}

	// Selects removed test keys by range.

	for _, key := range keys {
		tx, err := kvStore.Transact(false)
		if err != nil {
			t.Error(err)
			return
		}
		rs, err := tx.GetRange(key)
		if err != nil {
			cancel(t, tx)
			t.Error(err)
			return
		}
		if rs.Next() {
			cancel(t, tx)
			t.Errorf("key (%v) is not removed", key)
			return
		}
		if err := tx.Commit(); err != nil {
			t.Error(err)
			return
		}
	}
}
