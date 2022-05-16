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
	"math"
)

// Float represents a float value.
type Float struct {
	Value float32
}

// NewFloat returns a float instance.
func NewFloat() *Float {
	return NewFloatWithValue(0)
}

// NewFloatWithValue returns a float instance with the specified value.
func NewFloatWithValue(val float32) *Float {
	return &Float{Value: val}
}

// NewFloatWithBytes returns a float instance with the specified bytes.
func NewFloatWithBytes(src []byte) (*Float, []byte, error) {
	val, src, err := ReadFloat32Bytes(src)
	if err != nil {
		return nil, src, err
	}
	return &Float{Value: val}, src, nil
}

// GetData returns the value.
func (v *Float) GetData() interface{} {
	return v.Value
}

// SetValue sets a specified value.
func (v *Float) SetValue(value float32) {
	v.Value = value
}

// GetValue returns the stored value.
func (v *Float) GetValue() float32 {
	return v.Value
}

// Equals returns true when the specified value is s the same as this value, otherwise false.
func (v *Float) Equals(other Data) bool {
	if v.GetType() != other.GetType() {
		return false
	}
	otherValue, ok := other.GetData().(float32)
	if !ok {
		return false
	}
	if v.Value != otherValue {
		return false
	}
	return true
}

// Bytes returns the binary representation.
func (v *Float) Bytes() []byte {
	return AppendFloat32Bytes(nil, v.Value)
}

// ReadFloat32Bytes reads the specified bytes as a Float.
func ReadFloat32Bytes(src []byte) (float32, []byte, error) {
	bits, src, err := ReadUint32Bytes(src)
	if err != nil {
		return 0, src, err
	}
	return math.Float32frombits(bits), src, nil
}

// AppendFloat32Bytes appends a value to the specified buffer.
func AppendFloat32Bytes(buf []byte, val float32) []byte {
	return AppendUint32Bytes(buf, math.Float32bits(val))
}
