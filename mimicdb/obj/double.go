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

// Double represents a double value.
type Double struct {
	Value float64
}

// NewDouble returns a double instance.
func NewDouble() *Double {
	return NewDoubleWithValue(0)
}

// NewDoubleWithValue returns a double instance with the specified value.
func NewDoubleWithValue(val float64) *Double {
	return &Double{Value: val}
}

// NewDoubleWithBytes returns a double instance with the specified bytes.
func NewDoubleWithBytes(src []byte) (*Double, []byte, error) {
	val, src, err := ReadFloat64Bytes(src)
	if err != nil {
		return nil, src, err
	}
	return &Double{Value: val}, src, nil
}

// GetType returns the primitive type.
func (v *Double) GetType() Type {
	return DOUBLE
}

// GetData returns the value.
func (v *Double) GetData() interface{} {
	return v.Value
}

// SetValue sets a specified value.
func (v *Double) SetValue(value float64) {
	v.Value = value
}

// GetValue returns the stored value.
func (v *Double) GetValue() float64 {
	return v.Value
}

// Equals returns true when the specified value is s the same as this value, otherwise false.
func (v *Double) Equals(other Data) bool {
	if v.GetType() != other.GetType() {
		return false
	}
	otherValue, ok := other.GetData().(float64)
	if !ok {
		return false
	}
	if v.Value != otherValue {
		return false
	}
	return true
}

// Bytes returns the binary representation.
func (v *Double) Bytes() []byte {
	return AppendFloat64Bytes(nil, v.Value)
}

// ReadFloat64Bytes reads the specified bytes as a long Float.
func ReadFloat64Bytes(src []byte) (float64, []byte, error) {
	bits, src, err := ReadUint64Bytes(src)
	if err != nil {
		return 0, src, err
	}
	return math.Float64frombits(bits), src, nil
}

// AppendFloat64Bytes appends a value to the specified buffer.
func AppendFloat64Bytes(buf []byte, val float64) []byte {
	return AppendUint64Bytes(buf, math.Float64bits(val))
}
