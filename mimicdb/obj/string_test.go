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

func TestNewString(t *testing.T) {
	NewString()
}

func TestStringBytes(t *testing.T) {
	testValues := []string{
		"a",
		strings.Repeat("a", StringMaxLen/2),
		strings.Repeat("a", StringMaxLen),
		strings.Repeat("a", StringMaxLen+1),
		strings.Repeat("a", StringMaxLen*2),
	}
	for _, testVal := range testValues {
		testBytes := NewStringWithValue(testVal).Bytes()
		v, _, err := NewStringWithBytes(testBytes)
		if err != nil {
			t.Error(err)
			continue
		}
		vLen := len(v.Value)
		if StringMaxLen < vLen {
			t.Errorf("%d < %d", StringMaxLen, vLen)
		}
		if v.Value != testVal[:vLen] {
			t.Errorf("%s != %s", v.Value, testVal)
		}
	}
}
