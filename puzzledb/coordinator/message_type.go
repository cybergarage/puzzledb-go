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

// MessageType represents a coordinator message type.
type MessageType byte

const (
	// ObjectMessage represents a object message type.
	ObjectMessage MessageType = 'O'
	// DatabaseMessage represents a database message type.
	DatabaseMessage MessageType = 'D'
	// CollectionMessage represents a schema message type.
	CollectionMessage MessageType = 'C'
	// UserMessage represents a user message type.
	UserMessage MessageType = 'U'
)

// String returns the string representation of the message type.
func (t MessageType) String() string {
	switch t {
	case ObjectMessage:
		return "object"
	case DatabaseMessage:
		return "database"
	case CollectionMessage:
		return "collection"
	case UserMessage:
		return "user"
	default:
		return Unknown
	}
}
