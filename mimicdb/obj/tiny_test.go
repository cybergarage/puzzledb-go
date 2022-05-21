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
	"testing"
)

func TestNewTiny(t *testing.T) {
	NewTiny()
}

func TestTinyBytes(t *testing.T) {
	testValues := []int8{
		math.MinInt8,
		math.MinInt8 + 1,
		0,
		math.MaxInt8 / 2,
		math.MaxInt8,
	}
	for _, testVal := range testValues {
		testObj := NewTinyWithValue(testVal)
		testBytes := testObj.Bytes()
		parsedObj, _, err := NewTinyWithBytes(testBytes)
		if err != nil {
			t.Error(err)
			continue
		}
		// Compares generated objects using Equals() and private values strictly
		if !parsedObj.Equals(testObj) {
			t.Errorf("%v != %v", parsedObj, testObj)
		}
		if parsedObj.value != testObj.value {
			t.Errorf("%v != %v", parsedObj, testObj)
		}
	}
}
