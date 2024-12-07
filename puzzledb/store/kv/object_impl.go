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

package kv

type object struct {
	key   Key
	value []byte
}

// NewObject returns a new object.
func NewObject(key Key, value []byte) Object {
	obj := &object{
		key:   key,
		value: value,
	}
	return obj
}

// Key returns a key of the object.
func (obj *object) Key() Key {
	return obj.key
}

// Value returns a value of the object.
func (obj *object) Value() []byte {
	return obj.value
}
