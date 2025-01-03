// Copyright (C) 2020 The PuzzleDB Authors.
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

package document

import (
	"reflect"
	"strconv"
	"testing"
)

var elementTypes = []ElementType{
	Int8Type,
	Int16Type,
	Int32Type,
	Int64Type,
	StringType,
	BinaryType,
	Float32Type,
	Float64Type,
	DatetimeType,
	BoolType,
}

func TestElement(t *testing.T) {
	for n, et := range elementTypes {
		e1 := NewElement().SetName(strconv.Itoa(n)).SetType(et)
		e2, err := newElementWith(e1.Data())
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(e1.Data(), e2.Data()) {
			t.Errorf("%v !=%v", e2, e1)
		}
	}
}
