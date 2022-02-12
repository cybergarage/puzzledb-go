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

func TestNewArray(t *testing.T) {
	NewArray()
}

func TestArrayBytes(t *testing.T) {
	now := time.Now()

	vals := []Data{
		NewBoolWithValue(true),
		NewStringWithValue("abc"),
		NewShortWithValue(123),
		NewIntWithValue(123),
		NewLongWithValue(123),
		NewFloatWithValue(123),
		NewDoubleWithValue(123),
		NewTimestampWithValue(time.Unix(now.Unix(), now.UnixNano()%1e3)),
		NewDatetimeWithValue(time.Unix(now.Unix(), 0)),
		NewBinaryWithValue([]byte("abc")),
	}

	array := NewArray()
	for _, val := range vals {
		array.Append(val)
	}

	testBytes := array.Bytes()

	readArray, _, err := NewArrayWithBytes(testBytes)
	if err != nil {
		t.Error(err)
		return
	}

	if !array.Equals(readArray) {
		t.Errorf("%v != %v", array, readArray)
	}
}
