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

func TestNewTimestamp(t *testing.T) {
	NewTimestamp()
}

func TestTimestampBytes(t *testing.T) {
	now := time.Now()
	testValues := []time.Time{
		time.Unix(now.Unix(), now.UnixNano()%1e3),
	}
	for _, testVal := range testValues {
		testBytes := NewTimestampWithValue(testVal).Bytes()
		v, _, err := NewTimestampWithBytes(testBytes)
		if err != nil {
			t.Error(err)
			continue
		}
		if v.value != testVal {
			t.Errorf("%v != %v", v.value, testVal)
		}
	}
}
