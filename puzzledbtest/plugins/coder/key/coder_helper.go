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
	"strconv"

	"fmt"
	"testing"

	"github.com/cybergarage/go-pict/pict"
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
)

//go:embed key_types.pict
var testKeyTypes []byte

// nolint:goerr113, gocognit, gci, gocyclo
func CoderTest(t *testing.T, coder document.KeyCoder) {
	t.Helper()

	pict := pict.NewParserWithBytes(testKeyTypes)
	err := pict.Parse()
	if err != nil {
		t.Fatal(err)
	}

	getKeyValue := func(t string, v string) (any, error) {
		switch t {
		case "string":
			return v, nil
		case "bytes":
			return []byte(v), nil
		case "bool":
			return strconv.ParseBool(v)
		case "nil":
			return nil, nil
		case "int":
			i, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return nil, err
			}
			return int(i), nil
		case "int8":
			i, err := strconv.ParseInt(v, 10, 8)
			if err != nil {
				return nil, err
			}
			return int8(i), nil
		case "int16":
			i, err := strconv.ParseInt(v, 10, 16)
			if err != nil {
				return nil, err
			}
			return int16(i), nil
		case "int32":
			i, err := strconv.ParseInt(v, 10, 32)
			if err != nil {
				return nil, err
			}
			return int32(i), nil
		case "int64":
			i, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return nil, err
			}
			return int64(i), nil
		case "uint":
			i, err := strconv.ParseUint(v, 10, 64)
			if err != nil {
				return nil, err
			}
			return uint(i), nil
		case "uint8":
			i, err := strconv.ParseUint(v, 10, 8)
			if err != nil {
				return nil, err
			}
			return uint8(i), nil
		case "uint16":
			i, err := strconv.ParseUint(v, 10, 16)
			if err != nil {
				return nil, err
			}
			return uint16(i), nil
		case "uint32":
			i, err := strconv.ParseUint(v, 10, 32)
			if err != nil {
				return nil, err
			}
			return uint32(i), nil
		case "uint64":
			i, err := strconv.ParseUint(v, 10, 64)
			if err != nil {
				return nil, err
			}
			return uint64(i), nil
		case "float32":
			i, err := strconv.ParseFloat(v, 32)
			if err != nil {
				return nil, err
			}
			return float32(i), nil
		case "float64":
			i, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return nil, err
			}
			return float64(i), nil
		default:
			return nil, fmt.Errorf("unknown type: %s", t)
		}
	}

	pictParams := pict.Params()
	for i, pictCase := range pict.Cases() {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			key := document.NewKey()
			for n, pictParam := range pictParams {
				kv, err := getKeyValue(pictParam, pictCase[n])
				if err != nil {
					t.Error(err)
					return
				}
				key = append(key, kv)
			}
			kb, err := coder.Encode(key)
			if err != nil {
				t.Error(err)
				return
			}

			decKey, err := coder.Decode(kb)
			if err != nil {
				t.Error(err)
				return
			}

			if !key.Equals(decKey) {
				t.Errorf("%s != %s", key, decKey)
				return
			}
		})
	}
}
