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
)

// MessageType represents a coordinator message type.
type MessageType byte

const (
	// ObjectMessage represents a object message type.
	ObjectMessage MessageType = 'O'
	// SchemaMessage represents a schema message type.
	SchemaMessage MessageType = 'S'
)

// String returns the string representation of the message type.
func (t MessageType) String() string {
	switch t {
	case ObjectMessage:
		return "object"
	case SchemaMessage:
		return "schema"
	default:
		return Unknown
	}
}

// EventType represents a coordinator event type.
type EventType byte

const (
	// CreatedEvent represents a created event.
	CreatedEvent EventType = 'C'
	// UpdatedEvent represents a object updated event.
	UpdatedEvent EventType = 'U'
	// DeletedEvent represents a object deleted event.
	DeletedEvent EventType = 'O'
)

// String returns the string representation of the message event type.
func (t EventType) String() string {
	switch t {
	case CreatedEvent:
		return "created"
	case UpdatedEvent:
		return "updated"
	case DeletedEvent:
		return "deleted"
	default:
		return Unknown
	}
}

// Message represents a  coordinator event.
type Message interface {
	// Clock returns the message clock.
	Clock() cluster.Clock
	// From returns the destination node of the message.
	From() cluster.Node
	// Type returns the message type.
	Type() MessageType
	// EventType returns the message event type.
	EventType() EventType
	// Object returns the object of the message.
	Object() Object
	// Equals returns true if the message is equal to the specified event.
	Equals(Message) bool
	// String returns the string representation of the message.
	String() string
}
