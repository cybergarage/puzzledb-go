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

package tuple

import (
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
)

func newTupleWith(key document.Key) tuple.Tuple {
	tpl := make([]tuple.TupleElement, len(key))
	for n, keyElem := range key {
		switch v := keyElem.(type) {
		case int8:
			tpl[n] = int(v)
		case int16:
			tpl[n] = int(v)
		case int32:
			tpl[n] = int(v)
		case uint8:
			tpl[n] = int(v)
		case uint16:
			tpl[n] = int(v)
		case uint32:
			tpl[n] = int(v)
		default:
			tpl[n] = v
		}
	}
	return tpl
}

func newKeyWith(tpl tuple.Tuple) document.Key {
	key := make([]any, len(tpl))
	for n, tplElem := range tpl {
		key[n] = tplElem
	}
	return key
}
