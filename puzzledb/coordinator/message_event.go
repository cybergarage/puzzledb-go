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
