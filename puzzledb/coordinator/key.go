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
	"strings"
)

const (
	sep = " "
)

// Key represents an unique key for a key-value object.
type Key []string

// NewKey returns a new blank key.
func NewKey() Key {
	return Key{}
}

// NewKeyWith returns a new key from the specified key elements.
func NewKeyWith(elems ...string) Key {
	elemArray := make([]string, len(elems))
	copy(elemArray, elems)
	return elemArray
}

// NewKeyFrom returns a new key from the specified value.
func NewKeyFrom(v any) (Key, error) {
	switch v := v.(type) {
	case string:
		return NewKeyWith(strings.Split(v, sep)...), nil
	case []string:
		return NewKeyWith(v...), nil
	case []byte:
		return NewKeyWith(strings.Split(string(v), sep)...), nil
	case Key:
		return v, nil
	}
	return nil, newKeyInvalidError(v)
}

// Elements returns all elements of the key.
func (key Key) Elements() []string {
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

// Encode encodes the key to a byte array.
func (key Key) Encode() (string, error) {
	return strings.Join(key, sep), nil
}

// String returns the string representation of the key.
func (key Key) String() string {
	return fmt.Sprintf("[%s]", strings.Join(key, " "))
}
