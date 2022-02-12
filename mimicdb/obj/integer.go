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

// Int represents an integer value.
type Int struct {
	Value int32
}

// NewInt returns an integer instance.
func NewInt() *Int {
	return NewIntWithValue(0)
}

// NewIntWithValue returns an integer instance with the specified value.
func NewIntWithValue(val int32) *Int {
	return &Int{Value: val}
}

// NewIntWithBytes returns an integer instance with the specified bytes.
func NewIntWithBytes(src []byte) (*Int, []byte, error) {
	val, src, err := ReadInt32Bytes(src)
	if err != nil {
		return nil, src, err
	}
	return &Int{Value: val}, src, nil
}

// GetType returns the primitive type.
func (v *Int) GetType() Type {
	return INT
}

// GetData returns the value.
func (v *Int) GetData() interface{} {
	return v.Value
}

// SetValue sets a specified value.
func (v *Int) SetValue(value int32) {
	v.Value = value
}

// GetValue returns the stored value.
func (v *Int) GetValue() int32 {
	return v.Value
}

// Equals returns true when the specified value is s the same as this value, otherwise false.
func (v *Int) Equals(other Data) bool {
	if v.GetType() != other.GetType() {
		return false
	}
	otherValue, ok := other.GetData().(int32)
	if !ok {
		return false
	}
	if v.Value != otherValue {
		return false
	}
	return true
}

// Bytes returns the binary representation.
func (v *Int) Bytes() []byte {
	return AppendInt32Bytes(nil, v.Value)
}

// ReadInt32Bytes reads the specified bytes as a integer.
func ReadInt32Bytes(src []byte) (int32, []byte, error) {
	srcLen := len(src)
	if srcLen < 4 {
		return 0, nil, fmt.Errorf(errorInvalidIntegerBytes, src)
	}
	return (int32(src[0])<<24 | int32(src[1])<<16 | int32(src[2])<<8 | int32(src[3])), src[4:], nil
}

// AppendInt32Bytes appends a value to the specified buffer.
func AppendInt32Bytes(buf []byte, val int32) []byte {
	return append(buf,
		byte(val>>24),
		byte(val>>16),
		byte(val>>8),
		byte(val),
	)
}
