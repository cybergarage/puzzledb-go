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
	return fmt.Errorf("%v != %v", x, y) // nolint:goerr113
}

//nolint:gosec,cyclop,gocognit,gocyclo,maintidx
func DocumentStoreCRUDTest(t *testing.T, service plugins.Service) {
	t.Helper()

	testDBName := fmt.Sprintf("%s%d", testDBPrefix, time.Now().UnixNano())

	if err := service.Start(); err != nil {
		t.Error(err)
		return
	}
	if err := service.CreateDatabase(testDBName); err != nil {
		t.Error(err)
		return
	}

	hasDatabase := func(name string) bool {
		dbs, err := service.ListDatabases()
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

	db, err := service.GetDatabase(testDBName)
	if err != nil {
		t.Error(err)
		return
	}

	defer func() {
		if err := service.RemoveDatabase(testDBName); err != nil {
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
		for n, pictParam := range pict.Params() {
			kv, err := pictCase[n].CastType(string(pictParam))
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
			name := string(pictParam)
			pictElem := pictCase[n]
			v, err := pictElem.CastType(name)
			if err != nil {
				t.Error(err)
				return
			}
			obj[name] = v
		}
		objs[n] = obj
	}

	// Inserts objects

	cancel := func(t *testing.T, tx store.Transaction) {
		t.Helper()
		if err := tx.Cancel(); err != nil {
			t.Error(err)
		}
	}

	for n, key := range keys {
		tx, err := db.Transact(true)
		if err != nil {
			t.Error(err)
			return
		}
		err = tx.InsertDocument(key, objs[n])
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

	// Gets objects

	for n, key := range keys {
		tx, err := db.Transact(false)
		if err != nil {
			t.Error(err)
			return
		}
		rs, err := tx.FindDocuments(key)
		if err != nil {
			cancel(t, tx)
			t.Error(err)
			return
		}
		rsObjs := rs.Objects()
		if len(rsObjs) != 1 {
			cancel(t, tx)
			t.Errorf("objs != 1 (%d)", len(rsObjs))
			return
		}
		if err := deepEqual(rsObjs[0], objs[n]); err != nil {
			cancel(t, tx)
			t.Error(err)
			return
		}
		if err := tx.Commit(); err != nil {
			t.Error(err)
			return
		}
	}

	// Updates objects

	objs = make([]document.Object, len(pict.Cases()))
	for n, pictCase := range pict.Cases() {
		obj := []any{}
		for n, pictParam := range pict.Params() {
			name := string(pictParam)
			pictElem := pictCase[n]
			v, err := pictElem.CastType(name)
			if err != nil {
				t.Error(err)
				return
			}
			obj = append(obj, v)
		}
		objs[n] = obj
	}

	for n, key := range keys {
		tx, err := db.Transact(true)
		if err != nil {
			t.Error(err)
			return
		}
		err = tx.UpdateDocument(key, objs[n])
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

	// Gets updated objects

	for n, key := range keys {
		tx, err := db.Transact(false)
		if err != nil {
			t.Error(err)
			return
		}
		rs, err := tx.FindDocuments(key)
		if err != nil {
			cancel(t, tx)
			t.Error(err)
			return
		}
		rsObjs := rs.Objects()
		if len(rsObjs) != 1 {
			cancel(t, tx)
			t.Errorf("objs != 1 (%d)", len(rsObjs))
			return
		}
		if err := deepEqual(rsObjs[0], objs[n]); err != nil {
			cancel(t, tx)
			t.Error(err)
			return
		}
		if err := tx.Commit(); err != nil {
			t.Error(err)
			return
		}
	}

	// Removes objects

	for _, key := range keys {
		tx, err := db.Transact(true)
		if err != nil {
			t.Error(err)
			return
		}
		err = tx.RemoveDocument(key)
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

	// Gets removed objects

	for _, key := range keys {
		tx, err := db.Transact(false)
		if err != nil {
			t.Error(err)
			return
		}
		rs, err := tx.FindDocuments(key)
		if err != nil {
			cancel(t, tx)
			t.Error(err)
			return
		}
		rsObjs := rs.Objects()
		if len(rsObjs) != 0 {
			cancel(t, tx)
			t.Errorf("objs != 0 (%d)", len(rsObjs))
			return
		}
		if err := tx.Commit(); err != nil {
			t.Error(err)
			return
		}
	}
}
