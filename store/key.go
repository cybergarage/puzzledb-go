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

package store

import (
	"bytes"
	"fmt"
)

// Key represents an object key.
type Key = []any

func KeyToBytes(key Key) ([]byte, error) {
	var keyBuf bytes.Buffer
	for _, elem := range key {
		switch v := elem.(type) {
		case string:
			keyBuf.WriteString(v)
		default:
			return nil, fmt.Errorf("%w: (%T)", KeyTypeError, elem)
		}
	}
	return keyBuf.Bytes(), nil
}
