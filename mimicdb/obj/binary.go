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
	"fmt"
	"math"
)

const (
	// BinaryMaxLen is the maximum length of Binary data
	BinaryMaxLen = math.MaxUint32
)

// Binary represents a binary value.
type Binary struct {
	Value []byte
}

// NewBinary returns a binary instance.
func NewBinary() *Binary {
	return NewBinaryWithValue(make([]byte, 0))
}

// NewBinaryWithValue returns a binary instance with the specified value.
func NewBinaryWithValue(val []byte) *Binary {
	return &Binary{Value: val}
}

// NewBinaryWithBytes returns a binary instance with the specified bytes.
func NewBinaryWithBytes(src []byte) (*Binary, []byte, error) {
	val, src, err := ReadBinaryBytes(src)
	if err != nil {
		return nil, src, err
	}
	return &Binary{Value: val}, src, nil
}

// GetType returns the primitive type.
func (v *Binary) GetType() Type {
	return BINARY
}

// GetData returns the value.
func (v *Binary) GetData() interface{} {
	return v.Value
}

// SetValue sets a specified value.
func (v *Binary) SetValue(value []byte) {
	v.Value = value
}

// GetValue returns the stored value.
func (v *Binary) GetValue() []byte {
	return v.Value
}

// Equals returns true when the specified value is s the same as this value, otherwise false.
func (v *Binary) Equals(other Data) bool {
	if v.GetType() != other.GetType() {
		return false
	}
	otherValue, ok := other.GetData().([]byte)
	if !ok {
		return false
	}
	if string(v.Value) != string(otherValue) {
		return false
	}
	return true
}

// Bytes returns the binary representation.
func (v *Binary) Bytes() []byte {
	return AppendBinaryBytes(nil, v.Value)
}

// ReadBinaryBytes reads the specified bytes as a string.
func ReadBinaryBytes(src []byte) ([]byte, []byte, error) {
	binaryLen, src, err := ReadUint32Bytes(src)
	if err != nil {
		return nil, nil, err
	}
	srcLen := len(src)
	if srcLen < int(binaryLen) {
		return nil, nil, fmt.Errorf(errorInvalidBinaryBytes, srcLen, binaryLen, src)
	}
	return src[:binaryLen], src[binaryLen:], nil
}

// AppendBinaryBytes appends a string to the specified buffer.
func AppendBinaryBytes(buf []byte, val []byte) []byte {
	valByteLen := len(val)
	if StringMaxLen < valByteLen {
		valByteLen = BinaryMaxLen
	}
	buf = AppendUint32Bytes(buf, uint32(valByteLen))
	return append(buf, val[:valByteLen]...)
}
