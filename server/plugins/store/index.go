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
// 3: elements - []string

const (
	indexNameIdx     = 1
	indexTypeIdx     = 2
	indexElementsIdx = 3
)

type index struct {
	data     map[uint8]any
	elements []store.Element
}

// NewIndex returns a blank index.
func NewIndex() *index {
	idx := &index{
		data:     map[uint8]any{},
		elements: []store.Element{},
	}
	idx.data[indexElementsIdx] = []string{}
	return idx
}

// SetName sets the specified name to the index.
func (idx *index) SetName(name string) *index {
	idx.data[elementNameIdx] = name
	return idx
}

// Name returns the unique name.
func (idx *index) Name() string {
	v, ok := idx.data[indexNameIdx]
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
func (idx *index) SetType(t store.IndexType) *index {
	idx.data[indexTypeIdx] = uint8(t)
	return idx
}

// Type returns the index type.
func (idx *index) Type() store.IndexType {
	v, ok := idx.data[indexTypeIdx]
	if !ok {
		return 0
	}
	switch t := v.(type) {
	case store.IndexType:
		return store.IndexType(t)
	default:
		return 0
	}
}

// AddElement returns the schema elements.
func (idx *index) AddElement(elem store.Element) {
	idx.elements = append(idx.elements, elem)
	v, ok := idx.data[indexElementsIdx]
	if !ok {
		return
	}
	a, ok := v.([]string)
	if !ok {
		return
	}
	idx.data[indexElementsIdx] = append(a, elem.Name())
}

// Elements returns the schema elements.
func (idx *index) Elements() []store.Element {
	return idx.elements
}

// Data returns the raw representation data in memory.
func (idx *index) Data() any {
	return idx.data
}
