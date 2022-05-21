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

import "math"

const (
	// ArrayMaxSize is the maximum size of Array
	ArrayMaxSize = math.MaxUint16
)

// Array represents the data array.
type Array []Object

// NewArray returns an array instance.
func NewArray() Array {
	return make([]Object, 0)
}

// NewArrayWithBytes returns a array instance with the specified bytes.
func NewArrayWithBytes(src []byte) (Array, []byte, error) {
	var err error

	var nData uint16
	nData, src, err = ReadUint16Bytes(src)
	if err != nil {
		return nil, nil, err
	}

	array := NewArray()

	var val Object
	for n := 0; n < int(nData); n++ {
		val, src, err = NewObjectWithBytes(src)
		if err != nil {
			return nil, nil, err
		}
		array = append(array, val)
	}

	return array, src, nil
}

// GetData returns the value.
func (array Array) GetData() interface{} {
	return array
}

// Bytes returns the binary description.
func (array Array) Bytes() []byte {
	return AppendArrayBytes(nil, array)
}

// Append appends a value into the array.
func (array Array) Append(val Object) {
	array = append(array, val)
}

// Equals returns true when the specified array is the same as this array, otherwise false.
func (array Array) Equals(other Object) bool {
	otherArray, ok := other.(Array)
	if !ok {
		return false
	}

	arraySize := len(array)
	if arraySize != len(otherArray) {
		return false
	}

	for n := 0; n < int(arraySize); n++ {
		if !array[n].Equals(otherArray[n]) {
			return false
		}
	}
	return true
}

// AppendArrayBytes appends a value to the specified byte buffer.
func AppendArrayBytes(buf []byte, array Array) []byte {
	nData := len(array)
	if ArrayMaxSize < nData {
		nData = ArrayMaxSize
	}
	buf = AppendUint16Bytes(buf, uint16(nData))
	for _, val := range array {
		buf = AppendObjectBytes(buf, val)
	}
	return buf
}
