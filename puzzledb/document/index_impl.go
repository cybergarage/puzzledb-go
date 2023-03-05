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

// Schema format (version 1)
//
// map[uint8]any
// 1: name - string
// 2: type - uint8
// 3: elements - []string

const (
	indexNameIdx     = 1
	indexTypeIdx     = 2
	indexElementsIdx = 3
)

type indexMap = map[uint8]any
type indexElements = []string

type index struct {
	data     map[uint8]any
	elements []Element
}

// NewIndex returns a blank index.
func NewIndex() Index {
	idx := &index{
		data:     indexMap{},
		elements: []Element{},
	}
	idx.data[indexElementsIdx] = indexElements{}
	return idx
}

func newIndexWith(s *schema, obj any) (Index, error) {
	im, ok := obj.(indexMap)
	if !ok {
		return nil, newErrIndexInvalid(obj)
	}
	i := &index{
		data:     im,
		elements: nil,
	}

	// Caches index elements

	ies, ok := i.indexElements()
	if !ok {
		return nil, newErrSchemaInvalid(s)
	}
	i.elements = []Element{}
	for _, ie := range ies {
		em, err := s.FindElement(ie)
		if err != nil {
			return nil, newErrSchemaInvalid(s)
		}
		i.elements = append(i.elements, em)
	}

	return i, nil
}

// SetName sets the specified name to the index.
func (idx *index) SetName(name string) Index {
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
func (idx *index) SetType(t IndexType) Index {
	idx.data[indexTypeIdx] = t
	return idx
}

// Type returns the index type.
func (idx *index) Type() IndexType {
	v, ok := idx.data[indexTypeIdx]
	if !ok {
		return 0
	}
	switch t := v.(type) {
	case IndexType:
		return IndexType(t)
	default:
		return 0
	}
}

func (idx *index) indexElements() (indexElements, bool) {
	v, ok := idx.data[indexElementsIdx]
	if !ok {
		return nil, false
	}
	es, ok := v.(indexElements)
	if !ok {
		return nil, false
	}
	return es, true
}

// AddElement returns the schema elements.
func (idx *index) AddElement(elem Element) Index {
	es, ok := idx.indexElements()
	if !ok {
		return idx
	}
	idx.data[indexElementsIdx] = append(es, elem.Name())
	// Add element to cache
	idx.elements = append(idx.elements, elem)
	return idx
}

// Elements returns the schema elements.
func (idx *index) Elements() []Element {
	return idx.elements
}

// Data returns the raw representation data in memory.
func (idx *index) Data() any {
	return idx.data
}
