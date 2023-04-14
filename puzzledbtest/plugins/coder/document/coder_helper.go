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
	_ "embed"
	"fmt"
	"reflect"
	"testing"

	"github.com/cybergarage/go-pict/pict"
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
)

//go:embed go_types.pict
var goTypes []byte

// nolint:goerr113
func DeepEqual(x, y any) error {
	if reflect.DeepEqual(x, y) {
		return nil
	}
	if fmt.Sprintf("%v", x) == fmt.Sprintf("%v", y) {
		return nil
	}
	return fmt.Errorf("%v != %v", x, y)
}

func PrimitiveDocumentTest(t *testing.T, coder document.Coder) {
	t.Helper()

	pict := pict.NewParserWithBytes(goTypes)
	err := pict.Parse()
	if err != nil {
		t.Fatal(err)
	}

	pictParams := pict.Params()
	for _, pictCase := range pict.Cases() {
		for n, pictParam := range pictParams {
			pictElem := pictCase[n]
			obj, err := pictElem.CastType(string(pictParam))
			if err != nil {
				t.Error(err)
				return
			}

			var w bytes.Buffer
			err = coder.EncodeDocument(&w, obj)
			if err != nil {
				t.Error(err)
				return
			}

			r := bytes.NewReader(w.Bytes())
			decObj, err := coder.DecodeDocument(r)
			if err != nil {
				t.Error(err)
				return
			}

			err = DeepEqual(decObj, obj)
			if err != nil {
				t.Error(err)
				return
			}
		}
	}
}

func DocumentCoderTest(t *testing.T, coder document.Coder) {
	t.Helper()
	testFuncs := []struct {
		name string
		fn   func(*testing.T, document.Coder)
	}{
		{"primitive", PrimitiveDocumentTest},
	}

	for _, testFunc := range testFuncs {
		t.Run(testFunc.name, func(t *testing.T) {
			testFunc.fn(t, coder)
		})
	}
}
