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
	// StringMaxLen is the maximum length of String data
	StringMaxLen = math.MaxUint16
)

// String represents a string value.
type String struct {
	Value string
}

// NewString returns a string instance.
func NewString() *String {
	return NewStringWithValue("")
}

// NewStringWithValue returns a string instance with the specified value.
func NewStringWithValue(val string) *String {
	return &String{Value: val}
}

// NewStringWithBytes returns a string instance with the specified bytes.
func NewStringWithBytes(src []byte) (*String, []byte, error) {
	val, src, err := ReadStringBytes(src)
	if err != nil {
		return nil, src, err
	}
	return &String{Value: val}, src, nil
}

// GetData returns the value.
func (v *String) GetData() interface{} {
	return v.Value
}

// SetValue sets a specified value.
func (v *String) SetValue(value string) {
	v.Value = value
}

// GetValue returns the stored value.
func (v *String) GetValue() string {
	return v.Value
}

// Equals returns true when the specified value is s the same as this value, otherwise false.
func (v *String) Equals(other Object) bool {
	if _, ok := other.(*String); !ok {
		return false
	}
	otherValue, ok := other.GetData().(string)
	if !ok {
		return false
	}
	if v.Value != otherValue {
		return false
	}
	return true
}

// Bytes returns the binary representation.
func (v *String) Bytes() []byte {
	return AppendStringBytes(nil, v.Value)
}

// ReadStringBytes reads the specified bytes as a string.
func ReadStringBytes(src []byte) (string, []byte, error) {
	strLen, src, err := ReadUint16Bytes(src)
	if err != nil {
		return "", nil, err
	}
	srcLen := len(src)
	if srcLen < int(strLen) {
		return "", nil, fmt.Errorf(errorInvalidStringBytes, srcLen, strLen, src)
	}
	return string(src[:strLen]), src[strLen:], nil
}

// AppendStringBytes appends a string to the specified buffer.
func AppendStringBytes(buf []byte, val string) []byte {
	valByteLen := len(val)
	if StringMaxLen < valByteLen {
		valByteLen = StringMaxLen
	}
	buf = AppendUint16Bytes(buf, uint16(valByteLen))
	valBytes := []byte(val)
	return append(buf, valBytes[:valByteLen]...)
}
