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

import "fmt"

// Bool represents a bool value.
type Bool struct {
	value bool
}

// NewBool returns a bool instance.
func NewBool() *Bool {
	return NewBoolWithValue(false)
}

// NewBoolWithValue returns a bool instance with the specified value.
func NewBoolWithValue(v bool) *Bool {
	return &Bool{value: v}
}

// NewBoolWithBytes returns a bool instance with the specified bytes.
func NewBoolWithBytes(src []byte) (*Bool, []byte, error) {
	v, src, err := ReadBoolBytes(src)
	if err != nil {
		return nil, src, err
	}
	return &Bool{value: v}, src, nil
}

// Type returns the object type.
func (v *Bool) Type() Type {
	return BOOL
}

// Value returns the object value.
func (v *Bool) Value() interface{} {
	return v.value
}

// SetValue sets a specified value.
func (v *Bool) SetValue(value bool) {
	v.value = value
}

// GetValue returns the stored value.
func (v *Bool) GetValue() bool {
	return v.value
}

// Equals returns true when the specified value is s the same as this value, otherwise false.
func (v *Bool) Equals(other Object) bool {
	if _, ok := other.(*Bool); !ok {
		return false
	}
	otherValue, ok := other.Value().(bool)
	if !ok {
		return false
	}
	if v.value != otherValue {
		return false
	}
	return true
}

// Bytes returns the binary representation.
func (v *Bool) Bytes() []byte {
	return AppendBoolBytes(nil, v.value)
}

// ReadBoolBytes reads the specified bytes as a short integer.
func ReadBoolBytes(src []byte) (bool, []byte, error) {
	srcLen := len(src)
	if srcLen < 1 {
		return false, nil, fmt.Errorf(errorInvalidBooleanBytes, src)
	}
	if uint8(src[0]) != 0 {
		return true, src[1:], nil
	}
	return false, src[1:], nil
}

// AppendBoolBytes appends a value to the specified buffer.
func AppendBoolBytes(buf []byte, val bool) []byte {
	if val {
		return append(buf, byte(1))
	}
	return append(buf, byte(0))
}
