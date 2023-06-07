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
	"sync"
	"testing"
	"time"

	"github.com/cybergarage/puzzledb-go/puzzledb/cluster"
	"github.com/cybergarage/puzzledb-go/puzzledb/coordinator"
	plugin "github.com/cybergarage/puzzledb-go/puzzledb/plugins/coordinator"
)

type testObserver struct {
	sync.Mutex
	receivedMsgs []coordinator.Message
}

type testMessage struct {
	Value int
}

func newTestObserver() *testObserver {
	return &testObserver{
		Mutex:        sync.Mutex{},
		receivedMsgs: []coordinator.Message{},
	}
}

func (observer *testObserver) TotalValue() int {
	totalValue := 0
	for _, msg := range observer.receivedMsgs {
		var testObj testMessage
		if err := msg.UnmarshalTo(&testObj); err != nil {
			continue
		}
		totalValue += testObj.Value
	}
	return totalValue
}

func (observer *testObserver) OnMessageReceived(msg coordinator.Message) {
	observer.Lock()
	defer observer.Unlock()
	observer.receivedMsgs = append(observer.receivedMsgs, msg)
}

func (observer *testObserver) IsEventReceived(msg coordinator.Message) bool {
	for _, receivedMsg := range observer.receivedMsgs {
		if msg.Equals(receivedMsg) {
			return true
		}
	}
	return false
}

func CoordinatorMessageTest(t *testing.T, coords []plugin.Service) {
	t.Helper()

	observer := newTestObserver()
	for _, coord := range coords {
		if err := truncateCoordinatorStore(coord); err != nil {
			t.Error(err)
			return
		}
		err := coord.AddObserver(observer)
		if err != nil {
			t.Error(err)
			return
		}
	}

	// Generates test messages

	msgs := []coordinator.Message{}
	expectedTotalMessageValue := 0
	for n := 0; n < 10; n++ {
		obj := &testMessage{
			Value: n,
		}
		msg, err := coordinator.NewMessageWith(
			coordinator.UserMessage,
			coordinator.CreatedEvent,
			obj)
		if err != nil {
			t.Error(err)
			return
		}
		msgs = append(msgs, msg)
		expectedTotalMessageValue += obj.Value
	}

	// Posts test messages

	for n, msg := range msgs {
		var err error
		if (n % 2) == 0 {
			err = coords[0].PostMessage(msg)
		} else {
			err = coords[1].PostMessage(msg)
		}
		if err != nil {
			t.Error(err)
			return
		}
	}

	// Wait messages

	for n := 0; n < 10; n++ {
		if len(observer.receivedMsgs) == len(msgs) {
			break
		}
		// Waits for the received messages
		time.Sleep(plugin.DefaultStoreScanInterval)
	}

	// Checks the received messages

	if len(observer.receivedMsgs) != len(msgs) {
		t.Errorf("the number of received messages (%d) is not matched to the number of posted messages (%d)", len(observer.receivedMsgs), len(msgs))
		return
	}

	for _, msg := range msgs {
		if !observer.IsEventReceived(msg) {
			t.Errorf("message (%v) is not received", msg.String())
			return
		}
	}

	// Checks the received message value

	if observer.TotalValue() != expectedTotalMessageValue {
		t.Errorf("the total value of received messages (%d) is not matched to the expected value (%d)", observer.TotalValue(), expectedTotalMessageValue)
		return
	}

	// Checks the received messages order

	lastClock := cluster.Clock(0)
	for _, msg := range msgs {
		if msg.Clock() < lastClock {
			t.Errorf("the received messages are not sorted by clock")
			return
		}
		lastClock = msg.Clock()
	}
}
