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

// Object represents a key-value object.
type Object struct {
	Key   Key
	Value []byte
}

// NewObject returns a new object.
func NewObject(key Key, value []byte) *Object {
	obj := &Object{
		Key:   key,
		Value: value,
	}
	return obj
}
