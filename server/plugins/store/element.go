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
// 1: name - string
// 2: type - uint8

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
	v, ok := e.data[elementNameIdx]
	if !ok {
		return ""
	}
	switch name := v.(type) {
	case string:
		return name
	default:
		return ""
	}
}

// SetType sets the specified type to the element.
func (e *element) SetType(t store.ElementType) *element {
	e.data[elementTypeIdx] = uint8(t)
	return e
}

// Type returns the index type.
func (e *element) Type() store.ElementType {
	v, ok := e.data[elementTypeIdx]
	if !ok {
		return 0
	}
	switch t := v.(type) {
	case store.ElementType:
		return store.ElementType(t)
	default:
		return 0
	}
}

// Data returns the raw representation data in memory.
func (e *element) Data() any {
	return e.data
}
