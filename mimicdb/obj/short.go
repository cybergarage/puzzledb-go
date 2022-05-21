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

// Short represents a short value.
type Short struct {
	Value int16
}

// NewShort returns a short instance.
func NewShort() *Short {
	return NewShortWithValue(0)
}

// NewShortWithValue returns a short instance with the specified value.
func NewShortWithValue(val int16) *Short {
	return &Short{Value: val}
}

// NewShortWithBytes returns a short instance with the specified bytes.
func NewShortWithBytes(src []byte) (*Short, []byte, error) {
	val, src, err := ReadInt16Bytes(src)
	if err != nil {
		return nil, src, err
	}
	return &Short{Value: val}, src, nil
}

// GetData returns the value.
func (v *Short) GetData() interface{} {
	return v.Value
}

// SetValue sets a specified value.
func (v *Short) SetValue(value int16) {
	v.Value = value
}

// GetValue returns the stored value.
func (v *Short) GetValue() int16 {
	return v.Value
}

// Equals returns true when the specified value is s the same as this value, otherwise false.
func (v *Short) Equals(other Object) bool {
	if _, ok := other.(*Short); !ok {
		return false
	}
	otherValue, ok := other.GetData().(int16)
	if !ok {
		return false
	}
	if v.Value != otherValue {
		return false
	}
	return true
}

// Bytes returns the binary representation.
func (v *Short) Bytes() []byte {
	return AppendInt16Bytes(nil, v.Value)
}

// ReadInt16Bytes reads the specified bytes as a short integer.
func ReadInt16Bytes(src []byte) (int16, []byte, error) {
	srcLen := len(src)
	if srcLen < 2 {
		return 0, nil, fmt.Errorf(errorInvalidIntegerBytes, src)
	}
	return (int16(src[0])<<8 | int16(src[1])), src[2:], nil
}

// AppendInt16Bytes appends a value to the specified buffer.
func AppendInt16Bytes(buf []byte, val int16) []byte {
	return append(buf,
		byte(val>>8),
		byte(val),
	)
}
