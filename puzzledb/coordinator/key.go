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

package coordinator

import (
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
)

// Key represents an unique key for a key-value object.
type Key = document.Key

// NewKey returns a new blank key.
func NewKey() Key {
	return document.NewKey()
}

// NewKeyWith returns a new key from the specified key elements.
func NewKeyWith(elems ...any) Key {
	return document.NewKeyWith(elems...)
}
