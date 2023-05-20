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
	"github.com/cybergarage/puzzledb-go/puzzledb/coordinator"
	"github.com/google/uuid"
)

// ProcessObject represents a store process state object.
type MessageObject struct {
	ID     uuid.UUID
	Host   string
	Clock  uint64
	Type   uint8
	Object []byte
}

// NewScanMessageKey returns a new scan message key to get the latest message clock.
func NewScanMessageKey() coordinator.Key {
	return coordinator.NewKeyWith(coordinator.MessageObjectKeyHeader)
}

// NewMessageKeyWith returns a new message key with the specified message.
func NewMessageKeyWith(msg coordinator.Message, clock coordinator.Clock) coordinator.Key {
	return coordinator.NewKeyWith(coordinator.MessageObjectKeyHeader, clock, uint8(msg.Type()))
}

// NewMessageValueWith returns a new message value with the specified message.
func NewMessageValueWith(msg coordinator.Message, process coordinator.Process, clock coordinator.Clock) (coordinator.Value, error) {
	objBytes, err := msg.Object().Encode()
	if err != nil {
		return nil, err
	}
	obj := &MessageObject{
		ID:     process.ID(),
		Host:   process.Host(),
		Clock:  uint64(clock),
		Type:   uint8(msg.Type()),
		Object: objBytes,
	}
	return coordinator.NewValueWith(obj), nil
}
