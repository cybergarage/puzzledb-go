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
	"github.com/cybergarage/go-cbor/cbor"
)

type object struct {
	key   Key
	value Value
}

// NewObjectWith creates a new object with the specified key and value.
func NewObjectWith(key Key, value Value) Object {
	return &object{
		key:   key,
		value: value,
	}
}

// Key returns the key of the object.
func (obj *object) Key() Key {
	return obj.key
}

// Value returns the value of the object.
func (obj *object) Value() Value {
	return obj.value
}

// Encode encodes the object.
func (obj *object) Encode() ([]byte, error) {
	return cbor.Marshal(obj.value)
}
