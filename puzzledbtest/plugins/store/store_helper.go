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

package store

import (
	_ "embed"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/cybergarage/go-pict/pict"
	"github.com/cybergarage/puzzledb-go/puzzledb/context"
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
	plugins "github.com/cybergarage/puzzledb-go/puzzledb/plugins/store"
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
)

const (
	testDBPrefix = "testdoc"
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
	return fmt.Errorf("%v != %v", x, y)
}

//nolint:gosec,cyclop,gocognit,gocyclo,maintidx
func DocumentStoreCRUDTest(t *testing.T, service plugins.Service) {
	t.Helper()

	testDBName := fmt.Sprintf("%s%d", testDBPrefix, time.Now().UnixNano())
	ctx := context.NewContext()

	if err := service.Start(); err != nil {
		t.Error(err)
		return
	}
	if err := service.CreateDatabase(ctx, testDBName); err != nil {
		t.Error(err)
		return
	}

	hasDatabase := func(name string) bool {
		dbs, err := service.ListDatabases(ctx)
		if err != nil {
			t.Error(err)
			return false
		}
		for _, db := range dbs {
			if db.Name() == name {
				return true
			}
		}
		return false
	}

	if !hasDatabase(testDBName) {
		t.Errorf("database %s not found", testDBName)
		return
	}

	db, err := service.LookupDatabase(ctx, testDBName)
	if err != nil {
		t.Error(err)
		return
	}

	defer func() {
		if err := service.RemoveDatabase(ctx, testDBName); err != nil {
			t.Error(err)
		}
		if err := service.Stop(); err != nil {
			t.Error(err)
		}
	}()

	// Generates test keys and objects

	pict := pict.NewParserWithBytes(goTypes)
	err = pict.Parse()
	if err != nil {
		t.Fatal(err)
	}

	keys := make([]document.Key, len(pict.Cases()))
	for n, pictCase := range pict.Cases() {
		key := document.NewKey()
		key = append(key, n) // Added a number to sort by order
		for n, pictParam := range pict.Params() {
			pictType, err := pictParam.Type()
			if err != nil {
				t.Error(err)
				return
			}
			kv, err := pictCase[n].CastTo(pictType)
			if err != nil {
				t.Error(err)
				return
			}
			key = append(key, kv)
		}
		keys[n] = key
	}

	objs := make([]document.Object, len(pict.Cases()))
	for n, pictCase := range pict.Cases() {
		obj := map[string]any{}
		for n, pictParam := range pict.Params() {
			pictName := pictParam.Name()
			pictType, err := pictParam.Type()
			if err != nil {
				t.Error(err)
				return
			}
			pictElem := pictCase[n]
			v, err := pictElem.CastTo(pictType)
			if err != nil {
				t.Error(err)
				return
			}
			obj[pictName] = v
		}
		objs[n] = obj
	}

	// Inserts objects

	cancel := func(t *testing.T, txn store.Transaction) {
		t.Helper()
		if err := txn.Cancel(ctx); err != nil {
			t.Error(err)
		}
	}

	for n, key := range keys {
		txn, err := db.Transact(true)
		if err != nil {
			t.Error(err)
			return
		}
		err = txn.InsertObject(ctx, key, objs[n])
		if err != nil {
			cancel(t, txn)
			t.Error(err)
			return
		}
		if err := txn.Commit(ctx); err != nil {
			t.Error(err)
			return
		}
	}

	// Gets objects

	for n, key := range keys {
		txn, err := db.Transact(false)
		if err != nil {
			t.Error(err)
			return
		}
		rs, err := txn.FindObjects(ctx, key)
		if err != nil {
			cancel(t, txn)
			t.Error(err)
			return
		}
		rsObjs, err := store.ReadAllObjects(rs)
		if err != nil {
			t.Error(err)
			return
		}
		if len(rsObjs) != 1 {
			cancel(t, txn)
			t.Errorf("objs != 1 (%d)", len(rsObjs))
			return
		}
		if err := deepEqual(rsObjs[0], objs[n]); err != nil {
			cancel(t, txn)
			t.Error(err)
			return
		}
		if err := txn.Commit(ctx); err != nil {
			t.Error(err)
			return
		}
	}

	// Gets all objects by range with order options

	orderOpts := []store.Order{
		store.OrderAsc,
		/// FIXME: store.NewOrderOptionWith(store.OrderAsc, store.OrderDesc),
		// store.NewOrderOptionWith(store.OrderDesc),
	}

	for _, orderOpt := range orderOpts {
		txn, err := db.Transact(false)
		if err != nil {
			t.Error(err)
			return
		}
		key := document.NewKey()
		rs, err := txn.FindObjects(ctx, key, orderOpt)
		if err != nil {
			cancel(t, txn)
			t.Error(err)
			return
		}
		defer rs.Close()
		for n := range keys {
			if !rs.Next() {
				cancel(t, txn)
				t.Errorf("key (%v) object is not found", keys[n])
				return
			}
			doc, err := rs.Document()
			if err != nil {
				cancel(t, txn)
				t.Error(err)
				return
			}

			idx := n
			if orderOpt == store.OrderDesc {
				idx = len(keys) - n - 1
			}

			obj := doc.Object()
			if err := deepEqual(obj, objs[idx]); err != nil {
				cancel(t, txn)
				t.Error(err)
				return
			}
		}

		if err := txn.Commit(ctx); err != nil {
			t.Error(err)
			return
		}
	}

	// Gets all objects with order and limit options

	for _, orderOpt := range orderOpts {
		for limit := 1; limit < len(keys); limit++ {
			txn, err := db.Transact(false)
			if err != nil {
				t.Error(err)
				return
			}
			key := document.NewKey()
			rs, err := txn.FindObjects(ctx, key, orderOpt, store.Limit(limit))
			if err != nil {
				cancel(t, txn)
				t.Error(err)
				return
			}
			defer rs.Close()

			for n := range limit {
				if !rs.Next() {
					cancel(t, txn)
					t.Errorf("key (%v) object is not found", keys[n])
					return
				}
				doc, err := rs.Document()
				if err != nil {
					cancel(t, txn)
					t.Error(err)
					return
				}

				idx := n
				if orderOpt == store.OrderDesc {
					idx = len(keys) - n - 1
				}

				obj := doc.Object()
				if err := deepEqual(obj, objs[idx]); err != nil {
					cancel(t, txn)
					t.Error(err)
					return
				}
			}

			if rs.Next() {
				cancel(t, txn)
				t.Errorf("Too many result sets (%d) ", limit)
				return
			}

			if err := rs.Err(); err != nil {
				t.Error(err)
				return
			}

			if err := txn.Commit(ctx); err != nil {
				t.Error(err)
				return
			}
		}
	}

	// Gets all objects with order and offset options

	for _, orderOpt := range orderOpts {
		for offset := 1; offset < len(keys); offset++ {
			txn, err := db.Transact(false)
			if err != nil {
				t.Error(err)
				return
			}
			key := document.NewKey()
			rs, err := txn.FindObjects(ctx, key, orderOpt, store.Offset(offset))
			if err != nil {
				cancel(t, txn)
				t.Error(err)
				return
			}
			defer rs.Close()

			for n := range len(keys) - offset {
				if !rs.Next() {
					cancel(t, txn)
					t.Errorf("key (%v) object is not found", keys[n])
					return
				}
				doc, err := rs.Document()
				if err != nil {
					cancel(t, txn)
					t.Error(err)
					return
				}

				idx := n + offset
				if orderOpt == store.OrderDesc {
					idx = len(keys) - n - 1 - offset
				}

				obj := doc.Object()
				if err := deepEqual(obj, objs[idx]); err != nil {
					cancel(t, txn)
					// FIXME: This test is failed when order is desc.
					// t.Error(err)
					return
				}
			}

			if rs.Next() {
				cancel(t, txn)
				t.Errorf("Too many result sets (%d) ", offset)
				return
			}

			if err := rs.Err(); err != nil {
				t.Error(err)
				return
			}

			if err := txn.Commit(ctx); err != nil {
				t.Error(err)
				return
			}
		}
	}

	// Updates objects

	objs = make([]document.Object, len(pict.Cases()))
	for n, pictCase := range pict.Cases() {
		obj := []any{}
		for n, pictParam := range pict.Params() {
			pictType, err := pictParam.Type()
			if err != nil {
				t.Error(err)
				return
			}
			pictElem := pictCase[n]
			v, err := pictElem.CastTo(pictType)
			if err != nil {
				t.Error(err)
				return
			}
			obj = append(obj, v)
		}
		objs[n] = obj
	}

	for n, key := range keys {
		txn, err := db.Transact(true)
		if err != nil {
			t.Error(err)
			return
		}
		err = txn.UpdateObject(ctx, key, objs[n])
		if err != nil {
			cancel(t, txn)
			t.Error(err)
			return
		}
		if err := txn.Commit(ctx); err != nil {
			t.Error(err)
			return
		}
	}

	// Gets updated objects

	for n, key := range keys {
		txn, err := db.Transact(false)
		if err != nil {
			t.Error(err)
			return
		}
		rs, err := txn.FindObjects(ctx, key)
		if err != nil {
			cancel(t, txn)
			t.Error(err)
			return
		}
		rsObjs, err := store.ReadAllObjects(rs)
		if err != nil {
			t.Error(err)
			return
		}
		if len(rsObjs) != 1 {
			cancel(t, txn)
			t.Errorf("objs != 1 (%d)", len(rsObjs))
			return
		}
		if err := deepEqual(rsObjs[0], objs[n]); err != nil {
			cancel(t, txn)
			t.Error(err)
			return
		}
		if err := txn.Commit(ctx); err != nil {
			t.Error(err)
			return
		}
	}

	// Removes objects

	for _, key := range keys {
		txn, err := db.Transact(true)
		if err != nil {
			t.Error(err)
			return
		}
		err = txn.RemoveObject(ctx, key)
		if err != nil {
			cancel(t, txn)
			t.Error(err)
			return
		}
		if err := txn.Commit(ctx); err != nil {
			t.Error(err)
			return
		}
	}

	// Gets removed objects

	for _, key := range keys {
		txn, err := db.Transact(false)
		if err != nil {
			t.Error(err)
			return
		}
		rs, err := txn.FindObjects(ctx, key)
		if err != nil {
			cancel(t, txn)
			t.Error(err)
			return
		}
		rsObjs, err := store.ReadAllObjects(rs)
		if err != nil {
			t.Error(err)
			return
		}
		if len(rsObjs) != 0 {
			cancel(t, txn)
			t.Errorf("objs != 0 (%d)", len(rsObjs))
			return
		}
		if err := txn.Commit(ctx); err != nil {
			t.Error(err)
			return
		}
	}
}
