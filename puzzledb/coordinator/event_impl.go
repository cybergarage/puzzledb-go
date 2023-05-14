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

import "fmt"

// event represents a coordinator event.
type event struct {
	typ MessageType
	obj Object
}

// NewEventWith returns a new event with the specified type and object.
func NewEventWith(t MessageType, obj Object) Event {
	return &event{
		typ: t,
		obj: obj,
	}
}

// Type returns the event type.
func (e *event) Type() MessageType {
	return e.typ
}

// Object returns the object of the event.
func (e *event) Object() Object {
	return e.obj
}

// Equals returns true if the event is equal to the specified event.
func (e *event) Equals(other Event) bool {
	if e.Type() != other.Type() {
		return false
	}
	if e.Object().Equals(other.Object()) {
		return true
	}
	return false
}

// String returns the string representation of the event.
func (e *event) String() string {
	return fmt.Sprintf("%s %s", e.typ.String(), e.obj.String())
}
