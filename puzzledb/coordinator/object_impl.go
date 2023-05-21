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
	"reflect"

	"github.com/cybergarage/go-cbor/cbor"
)

type object struct {
	key   Key
	bytes []byte
}

// NewObjectWith creates a new object with the specified key and value.
func NewObjectWith(key Key, bytes []byte) Object {
	return &object{
		key:   key,
		bytes: bytes,
	}
}

// Key returns the key of the object.
func (obj *object) Key() Key {
	return obj.key
}

// Bytes returns the encoded object value.
func (obj *object) Bytes() []byte {
	return obj.bytes
}

// Unmarshal unmarshals the object value to the specified object.
func (obj *object) Unmarshal(to any) error {
	return cbor.UnmarshalTo(obj.bytes, to)
}

// Equals returns true if the object is equal to the specified object.
func (obj *object) Equals(other Object) bool {
	if obj == nil || other == nil {
		return false
	}
	if !obj.key.Equals(other.Key()) {
		return false
	}
	if reflect.DeepEqual(obj.bytes, other.Bytes()) {
		return true
	}
	return false
}

// String returns the string representation of the event.
func (obj *object) String() string {
	return fmt.Sprintf("%v %v", obj.key, obj.bytes)
}
