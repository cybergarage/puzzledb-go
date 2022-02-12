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

func TestNewDataWithValues(t *testing.T) {
	testValues := []interface{}{
		nil,
		true,
		"a",
		int(123),
		int16(123),
		int32(123),
		int64(123),
		float32(1.0),
		float64(1.0),
		time.Now(),
		[]byte("a"),
	}
	for _, testVal := range testValues {
		_, err := NewDataWithValue(testVal)
		if err != nil {
			t.Error(err)
		}
	}
}
