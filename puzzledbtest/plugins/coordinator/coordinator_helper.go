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
	_ "embed"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/cybergarage/go-pict/pict"
	"github.com/cybergarage/puzzledb-go/puzzledb/coordinator"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/coordinator/core"
)

//go:embed go_types.pict
var goTypes []byte

func deepEqual(x, y any) error {
	if reflect.DeepEqual(x, y) {
		return nil
	}
	if fmt.Sprintf("%v", x) == fmt.Sprintf("%v", y) {
		return nil
	}
	return fmt.Errorf("%v != %v", x, y) // nolint:goerr113
}

// nolint:goerr113, gocognit, gci, gocyclo, gosec, maintidx
func CoordinatorStoreTest(t *testing.T, s core.CoordinatorService) {
	t.Helper()

	cancel := func(t *testing.T, tx coordinator.Transaction) {
		t.Helper()
		if err := tx.Cancel(); err != nil {
			t.Error(err)
		}
	}

	// Starts the coordinator service

	if err := s.Start(); err != nil {
		t.Error(err)
		return
	}

	// Generates test keys and objects

	pict := pict.NewParserWithBytes(goTypes)
	err := pict.Parse()
	if err != nil {
		t.Fatal(err)
	}

	keys := make([]coordinator.Key, len(pict.Cases()))
	for n, pictCase := range pict.Cases() {
		key := coordinator.NewKey()
		for n, pictParam := range pict.Params() {
			kv, err := pictCase[n].CastType(string(pictParam))
			if err != nil {
				t.Error(err)
				return
			}
			key = append(key, fmt.Sprintf("%v", kv))
		}
		keys[n] = key
	}

	vals := make([]coordinator.Value, len(pict.Cases()))
	for n, pictCase := range pict.Cases() {
		val := map[string]any{}
		for n, pictParam := range pict.Params() {
			name := string(pictParam)
			pictElem := pictCase[n]
			v, err := pictElem.CastType(name)
			if err != nil {
				t.Error(err)
				return
			}
			val[name] = v
		}
		vals[n] = coordinator.NewValueWith(val)
	}

	// Inserts new objects

	for n, key := range keys {
		tx, err := s.Transact()
		if err != nil {
			t.Error(err)
			break
		}
		obj := coordinator.NewObjectWith(key, vals[n])
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

	// Selects inserted objects

	for n, key := range keys {
		tx, err := s.Transact()
		if err != nil {
			t.Error(err)
			break
		}
		obj, err := tx.Get(key)
		if err != nil {
			cancel(t, tx)
			t.Error(err)
			break
		}
		if err := deepEqual(obj.Value(), vals[n]); err != nil {
			cancel(t, tx)
			t.Error(err)
			break
		}
		if err := tx.Commit(); err != nil {
			t.Error(err)
			break
		}
	}

	// Updates inserted objects

	for n, pictCase := range pict.Cases() {
		val := []any{}
		for n, pictParam := range pict.Params() {
			name := string(pictParam)
			pictElem := pictCase[n]
			v, err := pictElem.CastType(name)
			if err != nil {
				t.Error(err)
				return
			}
			val = append(val, v)
		}
		vals[n] = coordinator.NewValueWith(val)
	}

	for n, key := range keys {
		tx, err := s.Transact()
		if err != nil {
			t.Error(err)
			break
		}
		obj := coordinator.NewObjectWith(key, vals[n])
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

	// Selects update objects

	for n, key := range keys {
		tx, err := s.Transact()
		if err != nil {
			t.Error(err)
			break
		}
		obj, err := tx.Get(key)
		if err != nil {
			cancel(t, tx)
			t.Error(err)
			break
		}

		if err := deepEqual(obj.Value(), vals[n]); err != nil {
			cancel(t, tx)
			t.Error(err)
			break
		}
		if err := tx.Commit(); err != nil {
			t.Error(err)
			break
		}
	}

	// Delete test

	for _, key := range keys {
		tx, err := s.Transact()
		if err != nil {
			t.Error(err)
			break
		}
		err = tx.Delete(key)
		if err != nil {
			cancel(t, tx)
			t.Error(err)
			break
		}
		_, err = tx.Get(key)
		if !errors.Is(err, coordinator.ErrNotExist) {
			cancel(t, tx)
			t.Error(err)
			break
		}
		if err := tx.Commit(); err != nil {
			t.Error(err)
			break
		}
	}

	// Terminates the coordinator service

	if err := s.Stop(); err != nil {
		t.Error(err)
		return
	}
}
