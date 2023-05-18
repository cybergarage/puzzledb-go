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
	"errors"
	"testing"

	"github.com/cybergarage/puzzledb-go/puzzledb/coordinator"
	plugins "github.com/cybergarage/puzzledb-go/puzzledb/plugins/coordinator"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/coordinator/core"
)

type testObserver struct {
	receivedEvents []coordinator.Message
}

func newqTestObserver() *testObserver {
	return &testObserver{
		receivedEvents: []coordinator.Message{},
	}
}

func (observer *testObserver) MessageReceived(msg coordinator.Message) {
	observer.receivedEvents = append(observer.receivedEvents, msg)
}

func (observer *testObserver) IsEventReceived(msg coordinator.Message) bool {
	for _, event := range observer.receivedEvents {
		if msg.Equals(event) {
			return true
		}
	}
	return false
}

// nolint:goerr113, gocognit, gci, gocyclo, gosec, maintidx
func CoordinatorObserverTest(t *testing.T, core core.CoordinatorService) {
	t.Helper()

	coord := plugins.NewServiceWith(core)
	coord.SetKeyCoder(newTestKeyCoder())

	cancel := func(t *testing.T, txn coordinator.Transaction) {
		t.Helper()
		if err := txn.Cancel(); err != nil {
			t.Error(err)
		}
	}

	// Starts the coordinator service

	if err := coord.Start(); err != nil {
		t.Error(err)
		return
	}

	// Terminates the coordinator service

	defer func() {
		if err := coord.Stop(); err != nil {
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

	observers := make([]*testObserver, 10)
	for n := range observers {
		observers[n] = newqTestObserver()
		err := coord.AddObserver(observers[n])
		if err != nil {
			t.Error(err)
			return
		}
	}

	// Inserts new objects

	for _, obj := range objs {
		tx, err := coord.Transact()
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
		for _, w := range observers {
			e := coordinator.NewMessageWith(coordinator.ObjectCreated, obj)
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
		tx, err := coord.Transact()
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
		for _, w := range observers {
			e := coordinator.NewMessageWith(coordinator.ObjectUpdated, obj)
			if !w.IsEventReceived(e) {
				t.Errorf("watcher did not receive event: %s", e.String())
				return
			}
		}
	}

	// Removes updated objects

	for _, obj := range objs {
		tx, err := coord.Transact()
		if err != nil {
			t.Error(err)
			return
		}
		err = tx.Remove(obj.Key())
		if err != nil {
			cancel(t, tx)
			t.Error(err)
			return
		}
		if err := tx.Commit(); err != nil {
			t.Error(err)
			return
		}
		tx, err = coord.Transact()
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
		for _, w := range observers {
			e := coordinator.NewMessageWith(coordinator.ObjectDeleted, obj)
			if !w.IsEventReceived(e) {
				t.Errorf("watcher did not receive event: %s", e.String())
				return
			}
		}
	}
}
