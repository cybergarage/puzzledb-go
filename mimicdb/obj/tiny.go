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

package obj

import (
	"strconv"
)

// Tiny represents a tiny value.
type Tiny struct {
	value int8
}

// NewTiny returns a tiny instance.
func NewTiny() *Tiny {
	return NewTinyWithValue(0)
}

// NewTinyWithValue returns a tiny instance with the specified value.
func NewTinyWithValue(v int8) *Tiny {
	return &Tiny{value: v}
}

// NewTinyWithBytes returns a tiny instance with the specified bytes.
func NewTinyWithBytes(src []byte) (*Tiny, []byte, error) {
	v, src, err := ReadInt8Bytes(src)
	if err != nil {
		return nil, src, err
	}
	return &Tiny{value: v}, src, nil
}

// Type returns the object type.
func (v *Tiny) Type() Type {
	return TINY
}

// Value returns the object value.
func (v *Tiny) Value() any {
	return v.value
}

// SetValue sets a specified value.
func (v *Tiny) SetValue(value int8) {
	v.value = value
}

// GetValue returns the stored value.
func (v *Tiny) GetValue() int8 {
	return v.value
}

// Equals returns true when the specified value is s the same as this value, otherwise false.
func (v *Tiny) Equals(other Object) bool {
	if _, ok := other.(*Tiny); !ok {
		return false
	}
	otherValue, ok := other.Value().(int8)
	if !ok {
		return false
	}
	if v.value != otherValue {
		return false
	}
	return true
}

// Bytes returns the binary representation.
func (v *Tiny) Bytes() []byte {
	return AppendInt8Bytes(nil, v.value)
}

// String returns the string representation.
func (v *Tiny) String() string {
	return strconv.FormatInt(int64(v.value), 10)
}
