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

	"github.com/cybergarage/puzzledb-go/puzzledb/store/errors"
)

// Key represents an unique key for a document object.
type Key []any

// NewKeyWith returns a new key from the specified key elements.
func NewKeyWith(elems ...any) Key {
	elemArray := make([]any, len(elems))
	for n, elem := range elems {
		elemArray[n] = elem
	}
	return elemArray
}

// Encode encodes the key to a byte array.
func (key Key) Encode() ([]byte, error) {
	var keyBuf bytes.Buffer
	for _, elem := range key {
		switch v := elem.(type) {
		case string:
			if _, err := keyBuf.WriteString(v); err != nil {
				return nil, err
			}
		case []byte:
			if _, err := keyBuf.Write(v); err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("%w: (%T)", errors.KeyTypeError, elem)
		}
	}
	return keyBuf.Bytes(), nil
}
