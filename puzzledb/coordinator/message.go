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

// Message represents a  coordinator event.
type Message interface {
	// Clock returns the message clock.
	Clock() cluster.Clock
	// From returns the destination node of the message.
	From() cluster.Node
	// Type returns the message type.
	Type() MessageType
	// Event returns the message event type.
	Event() EventType
	// Object returns the object of the message.
	Object() (any, error)
	// UnmarshalTo unmarshals the object value to the specified object.
	UnmarshalTo(to any) error
	// Equals returns true if the message is equal to the specified event.
	Equals(Message) bool
	// String returns the string representation of the message.
	String() string
}
