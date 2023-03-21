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

package coordinator

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"testing"

	"github.com/cybergarage/puzzledb-go/puzzledb/coordinator"
	plugins "github.com/cybergarage/puzzledb-go/puzzledb/server/plugins/coordinator/core"
)

const (
	testKeyCount  = 100
	testValBufMax = 8
)

//nolint:gosec,cyclop,revive
func CoordinatorTest(t *testing.T, s plugins.CoordinatorService) {
	t.Helper()

	cancel := func(t *testing.T, tx coordinator.Transaction) {
		t.Helper()
		if err := tx.Cancel(); err != nil {
			t.Error(err)
		}
	}

	if err := s.Start(); err != nil {
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
		tx, err := s.Transact()
		if err != nil {
			t.Error(err)
			break
		}
		val := vals[n]
		obj := coordinator.NewObjectWith(
			coordinator.NewKeyWith(key),
			val)

		if err := tx.Set(obj); err != nil {
			cancel(t, tx)
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
		tx, err := s.Transact()
		if err != nil {
			t.Error(err)
			break
		}
		obj, err := tx.Get([]any{key})
		if err != nil {
			cancel(t, tx)
			t.Error(err)
			break
		}
		val, ok := obj.Value().([]byte)
		if !ok {
			cancel(t, tx)
			t.Errorf("invalid value type: %T", obj.Value())
			break
		}
		if !bytes.Equal(val, vals[n]) {
			cancel(t, tx)
			t.Errorf("%s != %s", val, vals[n])
		}
		if err := tx.Commit(); err != nil {
			t.Error(err)
			break
		}
	}

	if err := s.Stop(); err != nil {
		t.Error(err)
		return
	}
}
