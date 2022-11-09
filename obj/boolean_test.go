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
	"testing"
)

func TestBool(t *testing.T) {
	NewBool()
}

func TestBoolBytes(t *testing.T) {
	testValues := []bool{
		true,
		false,
	}
	for _, testVal := range testValues {
		testBytes := NewBoolWithValue(testVal).Bytes()
		v, _, err := NewBoolWithBytes(testBytes)
		if err != nil {
			t.Error(err)
			continue
		}
		if v.value != testVal {
			t.Errorf("%t != %t", v.value, testVal)
		}
	}
}
