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

package document

import (
	"fmt"
)

// Key represents an unique key for a document object.
type Key []any

// NewKey returns a new blank key.
func NewKey() Key {
	return Key{}
}

// NewKeyWith returns a new key from the specified key elements.
func NewKeyWith(elems ...any) Key {
	elemArray := make([]any, len(elems))
	copy(elemArray, elems)
	return elemArray
}

// Elements returns all elements of the key.
func (key Key) Elements() []any {
	return key
}

// Len returns the number of elements of the key.
func (key Key) Len() int {
	return len(key)
}

// Database returns the database name of the key.
func (key Key) Database() (string, error) {
	if key.Len() < 1 {
		return "", newDatabaseKeyNotFoundError(key)
	}
	v, ok := key[0].(string)
	if !ok {
		return "", newDatabaseKeyNotFoundError(key)
	}
	return v, nil
}

// Collection returns the collection name of the key.
func (key Key) Collection() (string, error) {
	if key.Len() < 2 {
		return "", newCollectionKeyNotFoundError(key)
	}
	v, ok := key[1].(string)
	if !ok {
		return "", newCollectionKeyNotFoundError(key)
	}
	return v, nil
}

// Equals returns true if the specified key is equal to the key.
func (key Key) Equals(other Key) bool {
	if len(key) != len(other) {
		return false
	}
	for n, elem := range key {
		es := fmt.Sprintf("%v", elem)
		os := fmt.Sprintf("%v", other[n])
		if es != os {
			return false
		}
	}
	return true
}

// String returns a string representation of the key.
func (key Key) String() string {
	var s string
	for _, elem := range key {
		s += fmt.Sprintf("%v", elem)
	}
	return s
}
