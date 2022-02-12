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

func TestNewFloat(t *testing.T) {
	NewFloat()
}

func TestFloatBytes(t *testing.T) {
	testValues := []float32{
		math.SmallestNonzeroFloat32,
		0,
		math.MaxFloat32 / 2,
		math.MaxFloat32,
	}
	for _, testVal := range testValues {
		testBytes := NewFloatWithValue(testVal).Bytes()
		v, _, err := NewFloatWithBytes(testBytes)
		if err != nil {
			t.Error(err)
			continue
		}
		if v.Value != testVal {
			t.Errorf("%f != %f", v.Value, testVal)
		}
	}
}
