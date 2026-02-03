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
	"fmt"

	"github.com/cybergarage/go-cbor/cbor"
	"github.com/cybergarage/puzzledb-go/puzzledb/cluster"
	"github.com/google/uuid"
)

// MessageObject represents a message object.
type MessageObject struct {
	FromID      uuid.UUID
	FromCluster string
	FromHost    string
	MsgClock    uint64
	MsgType     byte
	EvtType     byte
	EncBytes    []byte
}

// NewMessageFrom returns a new message with the specified message object.
func NewMessageFrom(obj *MessageObject) Message {
	return obj
}

// NewMessageWith returns a new message with the specified type and object.
func NewMessageWith(t MessageType, e EventType, obj any) (Message, error) {
	objBytes, err := cbor.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return &MessageObject{
		FromID:      uuid.Nil,
		FromCluster: "",
		FromHost:    "",
		MsgClock:    0,
		EvtType:     byte(e),
		MsgType:     byte(t),
		EncBytes:    objBytes,
	}, nil
}

// NewMessageObject returns a new empty message.
func NewMessageObject() *MessageObject {
	return &MessageObject{
		FromID:      uuid.Nil,
		FromCluster: "",
		FromHost:    "",
		MsgClock:    0,
		EvtType:     0,
		MsgType:     0,
		EncBytes:    nil,
	}
}

// NewMessageObjectWith returns a new message value with the specified message.
func NewMessageObjectWith(msg Message, node cluster.Node, clock cluster.Clock) (*MessageObject, error) {
	obj, err := msg.Object()
	if err != nil {
		return nil, err
	}
	objBytes, err := cbor.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return &MessageObject{
		FromID:      node.ID(),
		FromCluster: node.Cluster(),
		FromHost:    node.Host(),
		MsgClock:    uint64(clock),
		MsgType:     byte(msg.Type()),
		EvtType:     byte(msg.Event()),
		EncBytes:    objBytes,
	}, nil
}

// ID returns the message ID.
func (obj *MessageObject) ID() uuid.UUID {
	return obj.FromID
}

// Clock returns the message clock.
func (obj *MessageObject) Clock() cluster.Clock {
	return obj.MsgClock
}

// From returns the destination node of the message.
func (obj *MessageObject) From() cluster.Node {
	node := cluster.NewNode()
	node.SetID(obj.FromID)
	node.SetCluster(obj.FromCluster)
	node.SetHost(obj.FromHost)
	return node
}

// Type returns the message type.
func (obj *MessageObject) Type() MessageType {
	return MessageType(obj.MsgType)
}

// Event returns the message event type.
func (obj *MessageObject) Event() EventType {
	return EventType(obj.EvtType)
}

// Object returns the object of the message.
func (obj *MessageObject) Object() (any, error) {
	return cbor.Unmarshal(obj.EncBytes)
}

// UnmarshalTo unmarshals the object value to the specified object.
func (obj *MessageObject) UnmarshalTo(to any) error {
	return cbor.UnmarshalTo(obj.EncBytes, to)
}

// Equals returns true if the message is equal to the specified message.
func (obj *MessageObject) Equals(other Message) bool {
	if obj.Type() != other.Type() {
		return false
	}
	if obj.Event() != other.Event() {
		return false
	}
	return true
}

// String returns the string representation of the message.
func (obj *MessageObject) String() string {
	str := fmt.Sprintf("%s %s %d %s %s",
		obj.FromID,
		obj.FromHost,
		obj.MsgClock,
		obj.Type().String(),
		obj.Event().String())
	data, err := obj.Object()
	if err == nil {
		str += fmt.Sprintf(" %v", data)
	}
	return str
}
