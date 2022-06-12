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

package plugins

import (
	"testing"
	"time"

	"github.com/cybergarage/mimicdb/mimicdb/obj"
	"github.com/cybergarage/mimicdb/mimicdb/plugins/serializer"
)

func SerializerTest(t *testing.T, serializer serializer.Serializer) {
	t.Helper()
	SerializerArrayTest(t, serializer)
}

//nolint:gomnd
func SerializerArrayTest(t *testing.T, serializer serializer.Serializer) {
	t.Helper()

	now := time.Now()
	now = now.Add(time.Duration((now.Nanosecond() % 1e6)) * -1)

	vals := []obj.Object{
		obj.NewBoolWithValue(true),
		obj.NewStringWithValue("abc"),
		obj.NewShortWithValue(123),
		obj.NewIntWithValue(123),
		obj.NewLongWithValue(123),
		obj.NewFloatWithValue(123),
		obj.NewDoubleWithValue(123),
		obj.NewTimestampWithValue(time.Unix(now.Unix(), int64(now.Nanosecond()))),
		obj.NewDatetimeWithValue(time.Unix(now.Unix(), 0)),
		obj.NewBinaryWithValue([]byte("abc")),
	}

	array := obj.NewArray()
	for _, val := range vals {
		array.Append(val)
	}

	testBytes, err := serializer.Encode(array)
	if err != nil {
		t.Error(err)
		return
	}

	parsedArray, _, err := obj.NewArrayWithBytes(testBytes)
	if err != nil {
		t.Error(err)
		return
	}

	if !array.Equals(parsedArray) {
		t.Errorf("%v != %v", array, parsedArray)
	}
}
