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
	"bytes"
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

// Encode encodes the key to a byte array.
func (key Key) Encode() ([]byte, error) {
	var keyBuf bytes.Buffer
	for _, elem := range key {
		switch v := elem.(type) {
		case string:
			if _, err := keyBuf.WriteString(v); err != nil {
				return nil, err
			}
		case []byte:
			if _, err := keyBuf.Write(v); err != nil {
				return nil, err
			}
		default:
			if _, err := keyBuf.WriteString(fmt.Sprintf("%v", v)); err != nil {
				return nil, err
			}
		}
	}
	return keyBuf.Bytes(), nil
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
