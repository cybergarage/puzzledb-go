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

// Store represents a coordination store inteface.
type Store interface {
	// SetKeyCoder sets the key coder.
	SetKeyCoder(coder KeyCoder)
	// DecodeKey returns the decoded key from the specified bytes if available, otherwise returns an error.
	DecodeKey([]byte) (Key, error)
	// EncodeKey returns the encoded bytes from the specified key if available, otherwise returns an error.
	EncodeKey(Key) ([]byte, error)
	// Transact begin a new transaction.
	Transact() (Transaction, error)
}

// Coordinator represents a coordination service.
type Coordinator interface {
	Store
	cluster.Node
	// SetNode sets the coordinator node.
	SetNode(node cluster.Node)
	// SetStateObject sets the state object for the specified key.
	SetStateObject(t StateType, obj Object) error
	// GetObject gets the object for the specified key and state type.
	GetStateObject(t StateType, key Key) (Object, error)
	// GetRangeObjects gets the result set for the specified key and state type.
	GetStateObjects(t StateType) (ResultSet, error)
	// PostMessage posts the specified message to the coordinator.
	PostMessage(msg Message) error
	// AddObserver adds the specified observer to the coordinator.
	AddObserver(observer Observer) error
}
