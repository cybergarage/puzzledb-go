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

const (
	elementNameIdx = 1
	elementTypeIdx = 2
)

type elementMap = map[uint8]any

type element struct {
	data elementMap
}

// NewElement returns a blank element.
func NewElement() Element {
	e := &element{
		data: elementMap{},
	}
	return e
}

func newElementWith(obj any) (Element, error) {
	em, ok := obj.(elementMap)
	if !ok {
		return nil, newElementInvalidError(obj)
	}
	e := &element{
		data: em,
	}
	return e, nil
}

// SetName sets the specified name to the element.
func (e *element) SetName(name string) Element {
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
func (e *element) SetType(t ElementType) Element {
	e.data[elementTypeIdx] = uint8(t)
	return e
}

// Type returns the index type.
func (e *element) Type() ElementType {
	v, ok := e.data[elementTypeIdx]
	if !ok {
		return 0
	}
	et, err := NewElementTypeWith(v)
	if err != nil {
		return 0
	}
	return et
}

// Data returns the raw representation data in memory.
func (e *element) Data() any {
	return e.data
}
