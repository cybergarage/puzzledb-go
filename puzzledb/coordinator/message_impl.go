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

	"github.com/cybergarage/puzzledb-go/puzzledb/cluster"
)

// messageImpl represents a coordinator messageImpl.
type messageImpl struct {
	cluster.Node
	typ EventType
	obj Object
}

// NewMessageWith returns a new message with the specified type and object.
func NewMessageWith(t EventType, obj Object) Message {
	return &messageImpl{
		Node: cluster.NewNode(),
		typ:  t,
		obj:  obj,
	}
}

// From returns the destination node of the message.
func (msg *messageImpl) From() cluster.Node {
	return msg.Node
}

// EventType returns the message event type.
func (msg *messageImpl) EventType() EventType {
	return msg.typ
}

// Object returns the object of the message.
func (msg *messageImpl) Object() Object {
	return msg.obj
}

// Equals returns true if the message is equal to the specified message.
func (msg *messageImpl) Equals(other Message) bool {
	if msg.EventType() != other.EventType() {
		return false
	}
	if msg.Object().Equals(other.Object()) {
		return false
	}
	return true
}

// String returns the string representation of the message.
func (msg *messageImpl) String() string {
	return fmt.Sprintf("%s %s", msg.typ.String(), msg.obj.String())
}
