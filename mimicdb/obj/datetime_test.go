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
	"time"
)

func TestNewDatatime(t *testing.T) {
	NewDatetime()
}

func TestDatatimeBytes(t *testing.T) {
	testValues := []time.Time{
		time.Unix(time.Now().Unix(), 0),
	}
	for _, testVal := range testValues {
		testBytes := NewDatetimeWithValue(testVal).Bytes()
		v, _, err := NewDatetimeWithBytes(testBytes)
		if err != nil {
			t.Error(err)
			continue
		}
		if v.Value != testVal {
			t.Errorf("%v != %v", v.Value, testVal)
		}
	}
}
