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

// Long represents a Long value.
type Long struct {
	Value int64
}

// NewLong returns a long instance.
func NewLong() *Long {
	return NewLongWithValue(0)
}

// NewLongWithValue returns a long instance with the specified value.
func NewLongWithValue(val int64) *Long {
	return &Long{Value: val}
}

// NewLongWithBytes returns a long instance with the specified bytes.
func NewLongWithBytes(src []byte) (*Long, []byte, error) {
	val, src, err := ReadInt64Bytes(src)
	if err != nil {
		return nil, src, err
	}
	return &Long{Value: val}, src, nil
}

// GetData returns the value.
func (v *Long) GetData() interface{} {
	return v.Value
}

// SetValue sets a specified value.
func (v *Long) SetValue(value int64) {
	v.Value = value
}

// GetValue returns the stored value.
func (v *Long) GetValue() int64 {
	return v.Value
}

// Equals returns true when the specified value is s the same as this value, otherwise false.
func (v *Long) Equals(other Data) bool {
	if v.GetType() != other.GetType() {
		return false
	}
	otherValue, ok := other.GetData().(int64)
	if !ok {
		return false
	}
	if v.Value != otherValue {
		return false
	}
	return true
}

// Bytes returns the binary representation.
func (v *Long) Bytes() []byte {
	return AppendInt64Bytes(nil, v.Value)
}

// ReadInt64Bytes reads the specified bytes as a long integer.
func ReadInt64Bytes(src []byte) (int64, []byte, error) {
	srcLen := len(src)
	if srcLen < 8 {
		return 0, nil, fmt.Errorf(errorInvalidIntegerBytes, src)
	}
	return (int64(src[0])<<56 | int64(src[1])<<48 | int64(src[2])<<40 | int64(src[3])<<32 | int64(src[4])<<24 | int64(src[5])<<16 | int64(src[6])<<8 | int64(src[7])), src[8:], nil
}

// AppendInt64Bytes appends a value to the specified buffer.
func AppendInt64Bytes(buf []byte, val int64) []byte {
	return append(buf,
		byte(val>>56),
		byte(val>>48),
		byte(val>>40),
		byte(val>>32),
		byte(val>>24),
		byte(val>>16),
		byte(val>>8),
		byte(val),
	)
}
