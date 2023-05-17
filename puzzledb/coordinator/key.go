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

// Key represents an unique key for a key-value object.
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

// Equals returns true if the key is equal to the specified key.
func (key Key) Equals(other Key) bool {
	if len(key) != len(other) {
		return false
	}
	for n, elem := range key {
		if elem != other[n] {
			return false
		}
	}
	return true
}
