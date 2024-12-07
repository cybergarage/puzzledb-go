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
	kvPlugins "github.com/cybergarage/puzzledb-go/puzzledb/plugins/store/kv"
	"github.com/cybergarage/puzzledb-go/puzzledb/store/kv"
)

const (
	testKeyCount  = 10
	testValBufMax = 8
)

//nolint:gosec,cyclop,gocognit,gocyclo,maintidx
func StoreTest(t *testing.T, kvStore kvPlugins.Service) {
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

	// Generates test objects and values.

	keys := make([]document.Key, testKeyCount)
	vals := make([][]byte, testKeyCount)
	for n := 0; n < testKeyCount; n++ {
		keys[n] = document.NewKeyWith(testKeyPrefix, fmt.Sprintf("key%d", n))
		vals[n] = make([]byte, testValBufMax)
		binary.LittleEndian.PutUint64(vals[n], uint64(n))
	}

	cancel := func(t *testing.T, txn kv.Transaction) {
		t.Helper()
		if err := txn.Cancel(); err != nil {
			t.Error(err)
		}
	}

	// Inserts test objects and values.

	for n, key := range keys {
		txn, err := kvStore.Transact(true)
		if err != nil {
			t.Error(err)
			return
		}
		obj := kv.NewObject(key, vals[n])
		if err := txn.Set(obj); err != nil {
			cancel(t, txn)
			t.Error(err)
			return
		}
		if err := txn.Commit(); err != nil {
			t.Error(err)
			return
		}
	}

	// Selects inserted test objects.

	for n, key := range keys {
		txn, err := kvStore.Transact(false)
		if err != nil {
			t.Error(err)
			return
		}
		obj, err := txn.Get(key)
		if err != nil {
			cancel(t, txn)
			t.Error(err)
			return
		}
		if !bytes.Equal(obj.Value(), vals[n]) {
			cancel(t, txn)
			t.Errorf("%s != %s", obj.Value(), vals[n])
		}
		if err := txn.Commit(); err != nil {
			t.Error(err)
			return
		}
	}

	// Selects inserted test objects by range

	for n, key := range keys {
		txn, err := kvStore.Transact(false)
		if err != nil {
			t.Error(err)
			return
		}
		rs, err := txn.GetRange(key)
		if err != nil {
			cancel(t, txn)
			t.Error(err)
			return
		}
		if !rs.Next() {
			cancel(t, txn)
			t.Errorf("key (%v) is not found", key)
			return
		}
		obj := rs.Object()
		if !bytes.Equal(obj.Value(), vals[n]) {
			cancel(t, txn)
			t.Errorf("%s != %s", obj.Value(), vals[n])
			return
		}
		if rs.Next() {
			cancel(t, txn)
			t.Errorf("other key (%v) is found", key)
			return
		}
		if err := txn.Commit(); err != nil {
			t.Error(err)
			return
		}
	}

	// Selects all inserted test objects by range with order options

	orderOpts := []*kv.OrderOption{
		kv.NewOrderOptionWith(kv.OrderAsc),
		kv.NewOrderOptionWith(kv.OrderDesc),
	}

	for _, orderOpt := range orderOpts {
		txn, err := kvStore.Transact(false)
		if err != nil {
			t.Error(err)
			return
		}

		prefixKey := document.NewKeyWith(testKeyPrefix)
		rs, err := txn.GetRange(prefixKey, orderOpt)
		if err != nil {
			cancel(t, txn)
			t.Error(err)
			return
		}

		for n := 0; n < testKeyCount; n++ {
			if !rs.Next() {
				cancel(t, txn)
				t.Errorf("key (%v) object is not found", keys[n])
				return
			}
			obj := rs.Object()

			idx := n
			if orderOpt.Order == kv.OrderDesc {
				idx = testKeyCount - n - 1
			}

			if !obj.Key().Equals(keys[idx]) {
				cancel(t, txn)
				t.Errorf("%s != %s", obj.Key, keys[idx])
				return
			}
			if !bytes.Equal(obj.Value(), vals[idx]) {
				cancel(t, txn)
				t.Errorf("%s != %s", obj.Value(), vals[idx])
				return
			}
		}

		if err := txn.Commit(); err != nil {
			t.Error(err)
			return
		}
	}

	// Selects all inserted test objects by range with desc order and limit options

	for _, orderOpt := range orderOpts {
		for limit := 1; limit < testKeyCount; limit++ {
			txn, err := kvStore.Transact(false)
			if err != nil {
				t.Error(err)
				return
			}

			prefixKey := document.NewKeyWith(testKeyPrefix)
			rs, err := txn.GetRange(prefixKey, orderOpt, kv.NewLimitOption(limit))
			if err != nil {
				cancel(t, txn)
				t.Error(err)
				return
			}

			for n := 0; n < limit; n++ {
				if !rs.Next() {
					cancel(t, txn)
					t.Errorf("key (%v) object is not found", keys[n])
					return
				}
				obj := rs.Object()

				idx := n
				if orderOpt.Order == kv.OrderDesc {
					idx = testKeyCount - n - 1
				}

				if !obj.Key().Equals(keys[idx]) {
					cancel(t, txn)
					t.Errorf("%s != %s", obj.Key, keys[idx])
					return
				}
				if !bytes.Equal(obj.Value(), vals[idx]) {
					cancel(t, txn)
					t.Errorf("%s != %s", obj.Value(), vals[idx])
					return
				}
			}

			if rs.Next() {
				cancel(t, txn)
				t.Errorf("Too many result sets (%d) ", limit)
				return
			}

			if err := txn.Commit(); err != nil {
				t.Error(err)
				return
			}
		}
	}

	// Selects all inserted test objects by range with desc order and offset options

	for _, orderOpt := range orderOpts {
		for offset := 0; offset < testKeyCount; offset++ {
			txn, err := kvStore.Transact(false)
			if err != nil {
				t.Error(err)
				return
			}

			prefixKey := document.NewKeyWith(testKeyPrefix)
			rs, err := txn.GetRange(prefixKey, orderOpt, kv.NewOffsetOption(uint(offset)))
			if err != nil {
				cancel(t, txn)
				t.Error(err)
				return
			}

			for n := 0; n < (testKeyCount - offset); n++ {
				if !rs.Next() {
					cancel(t, txn)
					t.Errorf("key (%v) object is not found", keys[n])
					return
				}
				obj := rs.Object()

				idx := n + offset
				if orderOpt.Order == kv.OrderDesc {
					idx = testKeyCount - n - 1 - offset
				}

				if !obj.Key().Equals(keys[idx]) {
					cancel(t, txn)
					t.Errorf("%s != %s", obj.Key, keys[idx])
					return
				}
				if !bytes.Equal(obj.Value(), vals[idx]) {
					cancel(t, txn)
					t.Errorf("%s != %s", obj.Value(), vals[idx])
					return
				}
			}

			if rs.Next() {
				cancel(t, txn)
				t.Errorf("Too many result sets (%d) ", offset)
				return
			}

			if err := txn.Commit(); err != nil {
				t.Error(err)
				return
			}
		}
	}

	// Updates inserted test object values.

	for n := 0; n < testKeyCount; n++ {
		binary.LittleEndian.PutUint64(vals[n], rand.Uint64())
	}

	for n, key := range keys {
		txn, err := kvStore.Transact(true)
		if err != nil {
			t.Error(err)
			return
		}
		obj := kv.NewObject(key, vals[n])
		if err := txn.Set(obj); err != nil {
			cancel(t, txn)
			t.Error(err)
			return
		}
		if err := txn.Commit(); err != nil {
			t.Error(err)
			return
		}
	}

	// Selects updated test objects.

	for n, key := range keys {
		txn, err := kvStore.Transact(false)
		if err != nil {
			t.Error(err)
			return
		}
		obj, err := txn.Get(key)
		if err != nil {
			cancel(t, txn)
			t.Error(err)
			return
		}
		if !bytes.Equal(obj.Value(), vals[n]) {
			cancel(t, txn)
			t.Errorf("%s != %s", obj.Value(), vals[n])
			return
		}
		if err := txn.Commit(); err != nil {
			t.Error(err)
			return
		}
	}

	// Selects updated test objects by range.

	for n, key := range keys {
		txn, err := kvStore.Transact(false)
		if err != nil {
			t.Error(err)
			return
		}
		rs, err := txn.GetRange(key)
		if err != nil {
			cancel(t, txn)
			t.Error(err)
			return
		}
		if !rs.Next() {
			cancel(t, txn)
			t.Errorf("key (%v) is not found", key)
			return
		}
		obj := rs.Object()
		if !bytes.Equal(obj.Value(), vals[n]) {
			cancel(t, txn)
			t.Errorf("%s != %s", obj.Value(), vals[n])
			return
		}
		if rs.Next() {
			cancel(t, txn)
			t.Errorf("other key (%v) is found", key)
			return
		}
		if err := txn.Commit(); err != nil {
			t.Error(err)
			return
		}
	}

	// Removes inserted test objects.

	for _, key := range keys {
		txn, err := kvStore.Transact(true)
		if err != nil {
			t.Error(err)
			return
		}
		err = txn.Remove(key)
		if err != nil {
			cancel(t, txn)
			t.Error(err)
			return
		}
		if err := txn.Commit(); err != nil {
			t.Error(err)
			return
		}
	}

	// Selects removed test objects.

	for _, key := range keys {
		txn, err := kvStore.Transact(false)
		if err != nil {
			t.Error(err)
			return
		}
		_, err = txn.Get(key)
		if err == nil {
			t.Errorf("key (%v) is found", key)
			cancel(t, txn)
			t.Error(err)
			return
		}
		if !errors.Is(err, kv.ErrNotExist) {
			t.Errorf("key (%v): %s", key, err.Error())
			cancel(t, txn)
			t.Error(err)
			return
		}
		if err := txn.Commit(); err != nil {
			t.Error(err)
			return
		}
	}

	// Selects removed test objects by range.

	for _, key := range keys {
		txn, err := kvStore.Transact(false)
		if err != nil {
			t.Error(err)
			return
		}
		rs, err := txn.GetRange(key)
		if err != nil {
			cancel(t, txn)
			t.Error(err)
			return
		}
		if rs.Next() {
			cancel(t, txn)
			t.Errorf("key (%v) is not removed", key)
			return
		}
		if err := txn.Commit(); err != nil {
			t.Error(err)
			return
		}
	}
}
