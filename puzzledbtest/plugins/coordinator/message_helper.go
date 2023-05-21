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
	"testing"
	"time"

	"github.com/cybergarage/puzzledb-go/puzzledb/coordinator"
	plugin "github.com/cybergarage/puzzledb-go/puzzledb/plugins/coordinator"
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
func CoordinatorMessageTest(t *testing.T, coord coordinator.Coordinator) {
	t.Helper()

	observer := newqTestObserver()
	err := coord.AddObserver(observer)
	if err != nil {
		t.Error(err)
		return
	}

	// Generates test messages
	msgs := []coordinator.Message{}
	for n := 0; n < 10; n++ {
		obj := coordinator.NewObjectWith(
			coordinator.NewKeyWith(n),
			n)
		msg := coordinator.NewMessageWith(
			coordinator.ObjectCreated,
			obj)
		msgs = append(msgs, msg)
	}

	// Posts test messages
	for _, msg := range msgs {
		err := coord.PostMessage(msg)
		if err != nil {
			t.Error(err)
			return
		}
	}

	// Waits for the received messages

	time.Sleep(plugin.DefaultStoreScanInterval * 2)

	// Checks the received messages
	for _, msg := range msgs {
		if !observer.IsEventReceived(msg) {
			t.Errorf("message (%v) is not received", msg.Object().Key())
			return
		}
	}
}
