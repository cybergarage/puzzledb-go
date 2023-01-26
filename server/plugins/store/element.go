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
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
)

// Schema format (version 1)
//
// 0: uint8 - version
// 1: string - name
// 2: colums - map[int8]any
//
// colums
//

const (
	elementNameIdx = 1
	elementTypeIdx = 2
)

type element struct {
	data map[uint8]any
}

// NewElement returns a blank schema.
func NewElement() *element {
	e := &element{
		data: map[uint8]any{},
	}
	return e
}

// SetName sets the specified name to the element.
func (e *element) SetName(name string) *element {
	e.data[elementNameIdx] = name
	return e
}

// Name returns the unique name.
func (e *element) Name() string {
	return ""
}

// Type returns the index type.
func (e *element) Type() store.ElementType {
	return 0
}
