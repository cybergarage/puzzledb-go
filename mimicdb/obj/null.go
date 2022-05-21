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

// Null represents a bool value.
type Null struct {
}

// NewNull returns a bool instance.
func NewNull() *Null {
	return &Null{}
}

// NewNullWithBytes returns a bool instance with the specified bytes.
func NewNullWithBytes(src []byte) (*Null, []byte, error) {
	src, err := ReadNullBytes(src)
	if err != nil {
		return nil, src, err
	}
	return NewNull(), src, nil
}

// Type returns the object type.
func (v *Null) Type() Type {
	return NULL
}

// GetData returns the value.
func (v *Null) GetData() interface{} {
	return nil
}

// Equals returns true when the specified value is s the same as this value, otherwise false.
func (v *Null) Equals(other Object) bool {
	if _, ok := other.(*Null); !ok {
		return false
	}
	return true
}

// Bytes returns the binary representation.
func (v *Null) Bytes() []byte {
	return AppendNullBytes(nil)
}

// ReadNullBytes reads the specified bytes as a null value.
func ReadNullBytes(src []byte) ([]byte, error) {
	srcLen := len(src)
	if srcLen < 1 {
		return nil, fmt.Errorf(errorInvalidNullBytes, src)
	}
	return src[1:], nil
}

// AppendNullBytes appends a null value to the specified buffer.
func AppendNullBytes(buf []byte) []byte {
	return append(buf, byte(0))
}
