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
	"strings"
	"testing"
)

func TestNewBinary(t *testing.T) {
	NewBinary()
}

func TestBinaryBytes(t *testing.T) {
	testValues := [][]byte{
		[]byte("a"),
		//NOTE: The following tests are disabled due to the long execution time
		//[]byte(strings.Repeat("a", BinaryMaxLen/2)),
		//[]byte(strings.Repeat("a", BinaryMaxLen)),
		[]byte(strings.Repeat("a", BinaryMaxLen+1)),
	}
	for _, testVal := range testValues {
		testBytes := NewBinaryWithValue(testVal).Bytes()
		v, _, err := NewBinaryWithBytes(testBytes)
		if err != nil {
			t.Error(err)
			continue
		}
		vLen := len(v.value)
		if BinaryMaxLen < vLen {
			t.Errorf("%d < %d", BinaryMaxLen, vLen)
		}
		if string(v.value) != string(testVal[:vLen]) {
			t.Errorf("%s != %s", string(v.value), string(testVal[:vLen]))
		}
	}
}
