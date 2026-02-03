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

package key

import (
	_ "embed"
	"fmt"
	"math/rand"
	"reflect"
	"testing"

	"github.com/cybergarage/go-pict/pict"
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
)

//go:embed go_types.pict
var testKeyTypes []byte

func deepEqual(x, y any) error {
	if reflect.DeepEqual(x, y) {
		return nil
	}
	if fmt.Sprintf("%v", x) == fmt.Sprintf("%v", y) {
		return nil
	}
	return fmt.Errorf("%v != %v", x, y)
}

// KeyCoderTest runs key coder conformance tests against the specified coder.
func KeyCoderTest(t *testing.T, coder document.KeyCoder) { //nolint:gocognit,gci,gocyclo,gosec,maintidx
	t.Helper()

	pict := pict.NewParserWithBytes(testKeyTypes)
	err := pict.Parse()
	if err != nil {
		t.Fatal(err)
	}

	shuffleKey := func(key document.Key) {
		n := len(key)
		for i := n - 1; i > 0; i-- {
			j := rand.Intn(i + 1)
			key[i], key[j] = key[j], key[i]
		}
	}

	for _, pictCase := range pict.Cases() {
		key := document.NewKey()
		for n, pictParam := range pict.Params() {
			pictElem := pictCase[n]
			pictType, err := pictParam.Type()
			if err != nil {
				t.Error(err)
				return
			}
			kv, err := pictElem.CastTo(pictType)
			if err != nil {
				t.Error(err)
				return
			}
			key = append(key, kv)
		}

		// Non-shuffled key

		kb, err := coder.EncodeKey(key)
		if err != nil {
			t.Error(err)
			return
		}

		decKey, err := coder.DecodeKey(kb)
		if err != nil {
			t.Error(err)
			return
		}

		if !key.Equals(decKey) {
			t.Errorf("%s != %s", key, decKey)
			return
		}

		// Random shuffled key

		shuffleKey(key)

		kb, err = coder.EncodeKey(key)
		if err != nil {
			t.Error(err)
			return
		}

		decKey, err = coder.DecodeKey(kb)
		if err != nil {
			t.Error(err)
			return
		}

		if !key.Equals(decKey) {
			t.Errorf("%s != %s", key, decKey)
			return
		}

		// Random reduced key

		kn := rand.Intn(len(key)-1) + 1
		key = key[:kn]

		kb, err = coder.EncodeKey(key)
		if err != nil {
			t.Error(err)
			return
		}

		decKey, err = coder.DecodeKey(kb)
		if err != nil {
			t.Error(err)
			return
		}

		if !key.Equals(decKey) {
			t.Errorf("%s != %s", key, decKey)
			return
		}
	}
}
