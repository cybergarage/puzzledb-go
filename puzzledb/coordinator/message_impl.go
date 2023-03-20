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

type message struct {
	typ MessageType
	obj any
}

// NewMessageWith creates a new message with the specified ID, type, and object.
func NewMessageWith(typ MessageType, obj any) Message {
	return &message{
		typ: typ,
		obj: obj,
	}
}

// Type returns the type of the message.
func (msg *message) Type() MessageType {
	return msg.typ
}

// Object returns the object of the message.
func (msg *message) Object() any {
	return msg.obj
}
