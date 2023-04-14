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

package document

import (
	"bytes"
	"fmt"
	"math"
	"reflect"
	"testing"

	"github.com/cybergarage/puzzledb-go/puzzledb/document"
)

// nolint:goerr113
func DeepEqual(x, y any) error {
	if x == y {
		return nil
	}
	if reflect.DeepEqual(x, y) {
		return nil
	}
	if fmt.Sprintf("%v", x) == fmt.Sprintf("%v", y) {
		return nil
	}
	return fmt.Errorf("%v != %v", x, y)
}

func DocumentCoderPrimitiveTest(t *testing.T, coder document.Coder) {
	t.Helper()

	tests := []struct {
		name string
		obj  any
	}{
		{"int", math.MaxInt},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var w bytes.Buffer
			err := coder.Encode(&w, test.obj)
			if err != nil {
				t.Error(err)
				return
			}

			r := bytes.NewReader(w.Bytes())
			decObj, err := coder.Decode(r)
			if err != nil {
				t.Error(err)
				return
			}

			err = DeepEqual(decObj, test.obj)
			if err != nil {
				t.Error(err)
				return
			}
		})
	}
}

func DocumentCoderTest(t *testing.T, coder document.Coder) {
	t.Helper()
	testFuncs := []struct {
		name string
		fn   func(*testing.T, document.Coder)
	}{
		{"primitive", DocumentCoderPrimitiveTest},
	}

	for _, testFunc := range testFuncs {
		t.Run(testFunc.name, func(t *testing.T) {
			testFunc.fn(t, coder)
		})
	}
}
