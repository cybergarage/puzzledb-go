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
	"github.com/cybergarage/puzzledb-go/puzzledb/cluster"
	"github.com/cybergarage/puzzledb-go/puzzledb/coordinator"
	"github.com/google/uuid"
)

// MessageObject represents a message object.
type MessageObject struct {
	ID    uuid.UUID
	Host  string
	Clock uint64
	Type  byte
	Bytes []byte
}

// NewMessageWith returns a new message with the specified message object.
func NewMessageWith(key coordinator.Key, obj *MessageObject) coordinator.Message {
	msg := coordinator.NewMessageWith(
		coordinator.MessageType(obj.Type),
		coordinator.NewObjectWith(key, obj.Bytes))
	msg.From().SetHost(obj.Host)
	msg.From().SetClock(obj.Clock)
	return msg
}

// NewMessageScanKey returns a new scan message key to get the latest message clock.
func NewMessageScanKey() coordinator.Key {
	return coordinator.NewKeyWith(coordinator.MessageObjectKeyHeader[:])
}

// NewMessageKeyWith returns a new message key with the specified message.
func NewMessageKeyWith(msg coordinator.Message, clock cluster.Clock) coordinator.Key {
	return coordinator.NewKeyWith(coordinator.MessageObjectKeyHeader[:], clock)
}

// NewMessageObject returns a new empty message.
func NewMessageObject() *MessageObject {
	return &MessageObject{
		ID:    uuid.Nil,
		Host:  "",
		Clock: 0,
		Type:  0,
		Bytes: nil,
	}
}

// NewMessageObjectWith returns a new message value with the specified message.
func NewMessageObjectWith(msg coordinator.Message, node cluster.Node, clock cluster.Clock) (*MessageObject, error) {
	return &MessageObject{
		ID:    node.ID(),
		Host:  node.Host(),
		Clock: uint64(clock),
		Type:  byte(msg.Type()),
		Bytes: msg.Object().Bytes(),
	}, nil
}
