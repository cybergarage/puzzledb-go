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
	"testing"

	"github.com/cybergarage/go-pict/pict"
	"github.com/cybergarage/puzzledb-go/puzzledb/coordinator"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/coordinator/core"
)

//go:embed go_types.pict
var goTypes []byte

func generateCoordinatorObjects() ([]coordinator.Object, error) {
	pict := pict.NewParserWithBytes(goTypes)
	err := pict.Parse()
	if err != nil {
		return []coordinator.Object{}, err
	}

	keys := make([]coordinator.Key, len(pict.Cases()))
	for i, pictCase := range pict.Cases() {
		key := coordinator.NewKey()
		key = append(key, fmt.Sprintf("key%d", i))
		for j, pictParam := range pict.Params() {
			kv, err := pictCase[j].CastType(string(pictParam))
			if err != nil {
				return []coordinator.Object{}, err
			}
			key = append(key, fmt.Sprintf("%v", kv))
		}
		keys[i] = key
	}

	vals := make([]coordinator.Value, len(pict.Cases()))
	for i, pictCase := range pict.Cases() {
		val := map[string]any{}
		for j, pictParam := range pict.Params() {
			name := string(pictParam)
			pictElem := pictCase[j]
			v, err := pictElem.CastType(name)
			if err != nil {
				return []coordinator.Object{}, err
			}
			val[name] = v
		}
		vals[i] = coordinator.NewValueWith(val)
	}

	objs := make([]coordinator.Object, len(pict.Cases()))
	for n, key := range keys {
		objs[n] = coordinator.NewObjectWith(key, vals[n])
	}

	return objs, nil
}

func updateCoordinatorObjects(objs []coordinator.Object) ([]coordinator.Object, error) {
	pict := pict.NewParserWithBytes(goTypes)
	err := pict.Parse()
	if err != nil {
		return []coordinator.Object{}, err
	}

	for i, pictCase := range pict.Cases() {
		val := []any{}
		for j, pictParam := range pict.Params() {
			name := string(pictParam)
			pictElem := pictCase[j]
			v, err := pictElem.CastType(name)
			if err != nil {
				return []coordinator.Object{}, err
			}
			val = append(val, v)
		}
		objs[i] = coordinator.NewObjectWith(objs[i].Key(), val)
	}

	return objs, nil
}

// nolint:goerr113, gocognit, gci, gocyclo, gosec, maintidx
func CoordinatorStoreTest(t *testing.T, s core.CoordinatorService) {
	t.Helper()

	cancel := func(t *testing.T, txn coordinator.Transaction) {
		t.Helper()
		if err := txn.Cancel(); err != nil {
			t.Error(err)
		}
	}

	// Starts the coordinator service

	if err := s.Start(); err != nil {
		t.Error(err)
		return
	}

	// Terminates the coordinator service

	defer func() {
		if err := s.Stop(); err != nil {
			t.Error(err)
		}
	}()

	// Generates test keys and objects

	objs, err := generateCoordinatorObjects()
	if err != nil {
		t.Error(err)
		return
	}

	// Inserts new objects

	for _, obj := range objs {
		tx, err := s.Transact()
		if err != nil {
			t.Error(err)
			return
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

	// Selects inserted objects

	for _, obj := range objs {
		tx, err := s.Transact()
		if err != nil {
			t.Error(err)
			return
		}
		retObj, err := tx.Get(obj.Key())
		if err != nil {
			cancel(t, tx)
			t.Error(err)
			return
		}
		if !retObj.Equals(obj) {
			cancel(t, tx)
			t.Error(err)
			return
		}
		if err := tx.Commit(); err != nil {
			t.Error(err)
			return
		}
	}

	// Updates inserted objects

	objs, err = updateCoordinatorObjects(objs)
	if err != nil {
		t.Error(err)
		return
	}

	for _, obj := range objs {
		tx, err := s.Transact()
		if err != nil {
			t.Error(err)
			return
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

	// Selects update objects

	for _, obj := range objs {
		tx, err := s.Transact()
		if err != nil {
			t.Error(err)
			return
		}
		retObj, err := tx.Get(obj.Key())
		if err != nil {
			cancel(t, tx)
			t.Error(err)
			return
		}

		if !retObj.Equals(obj) {
			cancel(t, tx)
			t.Error(err)
			return
		}
		if err := tx.Commit(); err != nil {
			t.Error(err)
			return
		}
	}

	// Deletes updated objects

	for _, obj := range objs {
		tx, err := s.Transact()
		if err != nil {
			t.Error(err)
			return
		}
		err = tx.Delete(obj.Key())
		if err != nil {
			cancel(t, tx)
			t.Error(err)
			return
		}
		_, err = tx.Get(obj.Key())
		if !errors.Is(err, coordinator.ErrNotExist) {
			cancel(t, tx)
			t.Error(err)
			return
		}
		if err := tx.Commit(); err != nil {
			t.Error(err)
			return
		}
	}
}

type testWatcher struct {
	receivedEvents []coordinator.Event
}

func newqTestWatcher() *testWatcher {
	return &testWatcher{
		receivedEvents: []coordinator.Event{},
	}
}

func (w *testWatcher) ProcessEvent(e coordinator.Event) {
	w.receivedEvents = append(w.receivedEvents, e)
}

func (w *testWatcher) IsEventReceived(e coordinator.Event) bool {
	for _, event := range w.receivedEvents {
		if e.Equals(event) {
			return true
		}
	}
	return false
}

// nolint:goerr113, gocognit, gci, gocyclo, gosec, maintidx
func CoordinatorWatcherTest(t *testing.T, s core.CoordinatorService) {
	t.Helper()

	cancel := func(t *testing.T, txn coordinator.Transaction) {
		t.Helper()
		if err := txn.Cancel(); err != nil {
			t.Error(err)
		}
	}

	// Starts the coordinator service

	if err := s.Start(); err != nil {
		t.Error(err)
		return
	}

	// Terminates the coordinator service

	defer func() {
		if err := s.Stop(); err != nil {
			t.Error(err)
		}
	}()

	// Generates test keys and objects

	objs, err := generateCoordinatorObjects()
	if err != nil {
		t.Error(err)
		return
	}

	// Registers watcheres

	watchers := make([]*testWatcher, 10)
	for n := range watchers {
		watchers[n] = newqTestWatcher()
	}

	for _, obj := range objs {
		for _, w := range watchers {
			err := s.Watch(obj.Key(), w)
			if err != nil {
				t.Error(err)
				return
			}
		}
	}

	// Inserts new objects

	for _, obj := range objs {
		tx, err := s.Transact()
		if err != nil {
			t.Error(err)
			return
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

	// Checks if watchers received insert events

	for _, obj := range objs {
		for _, w := range watchers {
			e := coordinator.NewEventWith(coordinator.ObjectCreated, obj)
			if !w.IsEventReceived(e) {
				t.Errorf("watcher did not receive event: %s", e.String())
				return
			}
		}
	}

	// Updates inserted objects

	objs, err = updateCoordinatorObjects(objs)
	if err != nil {
		t.Error(err)
		return
	}

	for _, obj := range objs {
		tx, err := s.Transact()
		if err != nil {
			t.Error(err)
			return
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

	// Checks if watchers received update events

	for _, obj := range objs {
		for _, w := range watchers {
			e := coordinator.NewEventWith(coordinator.ObjectUpdated, obj)
			if !w.IsEventReceived(e) {
				t.Errorf("watcher did not receive event: %s", e.String())
				return
			}
		}
	}

	// Deletes updated objects

	for _, obj := range objs {
		tx, err := s.Transact()
		if err != nil {
			t.Error(err)
			return
		}
		err = tx.Delete(obj.Key())
		if err != nil {
			cancel(t, tx)
			t.Error(err)
			return
		}
		if err := tx.Commit(); err != nil {
			t.Error(err)
			return
		}
		tx, err = s.Transact()
		if err != nil {
			t.Error(err)
			return
		}
		_, err = tx.Get(obj.Key())
		if !errors.Is(err, coordinator.ErrNotExist) {
			cancel(t, tx)
			t.Error(err)
			return
		}
		if err := tx.Commit(); err != nil {
			t.Error(err)
			return
		}
	}

	// Checks if watchers received delete events

	for _, obj := range objs {
		for _, w := range watchers {
			e := coordinator.NewEventWith(coordinator.ObjectDeleted, obj)
			if !w.IsEventReceived(e) {
				t.Errorf("watcher did not receive event: %s", e.String())
				return
			}
		}
	}
}

func CoordinatorTest(t *testing.T, s core.CoordinatorService) {
	t.Helper()
	CoordinatorStoreTest(t, s)
	CoordinatorWatcherTest(t, s)
}
