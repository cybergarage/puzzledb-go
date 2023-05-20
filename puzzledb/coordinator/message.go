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

// MessageType represents a coordinator event type.
type MessageType uint8

const (
	// ObjectCreated represents a object created event.
	ObjectCreated MessageType = 1
	// ObjectUpdated represents a object updated event.
	ObjectUpdated MessageType = 2
	// ObjectDeleted represents a object deleted event.
	ObjectDeleted MessageType = 3
)

func (t MessageType) String() string {
	switch t {
	case ObjectCreated:
		return "created"
	case ObjectUpdated:
		return "updated"
	case ObjectDeleted:
		return "deleted"
	default:
		return "unknown"
	}
}

// Message represents a  coordinator event.
type Message interface {
	Process
	// Clock returns the message clock.
	Clock() Clock
	// From returns the destination process of the message.
	From() Process
	// Type returns the message type.
	Type() MessageType
	// Object returns the object of the message.
	Object() Object
	// Equals returns true if the message is equal to the specified event.
	Equals(Message) bool
	// String returns the string representation of the message.
	String() string
}
