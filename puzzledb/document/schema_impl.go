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
// 0: uint8 - version
// 1: string - name
// 2: elements - []map[uint8]any
//    1: name - string
//    2: type - uint8
// 3: indexes - []map[uint8]any
//    1: name - string
//    2: type - uint8
//    3: elements - []string (element name)

const (
	// SchemaVersion specifies a latest schema version.
	SchemaVersion = 1
)

const (
	schemaVersionIdx  = 0
	schemaNameIdx     = 1
	schemaElementsIdx = 2
	schemaIndexesIdx  = 3
)

type schemaMap = map[uint8]any

type schema struct {
	data     schemaMap
	elements []Element
	indexes  []Index
}

// NewSchema returns a blank schema.
func NewSchema() Schema {
	s := &schema{
		data:     schemaMap{},
		elements: []Element{},
		indexes:  []Index{},
	}
	s.SetVersion(SchemaVersion)
	s.data[schemaElementsIdx] = []elementMap{}
	s.data[schemaIndexesIdx] = []indexMap{}
	return s
}

// NewSchemaWith creates a schema from the specified object.
func NewSchemaWith(obj any) (Schema, error) {
	smap, ok := obj.(schemaMap)
	if !ok {
		amap, ok := obj.(map[any]any)
		if !ok {
			return nil, newSchemaInvalidError(obj)
		}
		smap = schemaMap{}
		for ak, av := range amap {
			switch k := ak.(type) {
			case int8:
				smap[uint8(k)] = av
			case uint8:
				smap[uint8(k)] = av
			default:
				return nil, newSchemaInvalidError(obj)
			}
		}
	}

	s := &schema{
		data:     smap,
		elements: []Element{},
		indexes:  []Index{},
	}

	// Caches elements

	ems, ok := s.elementMaps()
	if !ok {
		return nil, newSchemaInvalidError(obj)
	}

	for _, em := range ems {
		e, err := newElementWith(em)
		if err == nil {
			s.elements = append(s.elements, e)
		}
	}

	// Caches indexes

	ims, ok := s.indexMpas()
	if !ok {
		return nil, newSchemaInvalidError(obj)
	}

	for _, im := range ims {
		i, err := newIndexWith(s, im)
		if err == nil {
			s.indexes = append(s.indexes, i)
		}
	}

	return s, nil
}

// SetVersion sets the specified version to the schema.
func (s *schema) SetVersion(ver int) {
	s.data[schemaVersionIdx] = ver
}

// Version returns the schema version.
func (s *schema) Version() int {
	v, ok := s.data[schemaVersionIdx]
	if !ok {
		return 0
	}
	switch ver := v.(type) {
	case int:
		return int(ver)
	default:
		return 0
	}
}

// SetName sets the specified name to the schema.
func (s *schema) SetName(name string) {
	s.data[schemaNameIdx] = name
}

// Name returns the schema name.
func (s *schema) Name() string {
	v, ok := s.data[schemaNameIdx]
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

func (s *schema) elementMaps() ([]elementMap, bool) {
	v, ok := s.data[schemaElementsIdx]
	if !ok {
		return nil, false
	}
	ems, ok := v.([]elementMap)
	if !ok {
		return nil, false
	}
	return ems, true
}

// AddElement adds the specified element to the schema.
func (s *schema) AddElement(elem Element) {
	ems, ok := s.elementMaps()
	if !ok {
		return
	}
	em, ok := elem.Data().(elementMap)
	if !ok {
		return
	}
	s.data[schemaElementsIdx] = append(ems, em)
	// Add element to cache
	s.elements = append(s.elements, elem)
}

// Elements returns the schema elements.
func (s *schema) Elements() []Element {
	return s.elements
}

// FindElement returns the schema elements by the specified name.
func (s *schema) FindElement(name string) (Element, error) {
	es := s.Elements()
	for _, e := range es {
		if e.Name() == name {
			return e, nil
		}
	}
	return nil, newNotSupportedError(name)
}

func (s *schema) indexMpas() ([]indexMap, bool) {
	v, ok := s.data[schemaIndexesIdx]
	if !ok {
		return nil, false
	}
	ims, ok := v.([]indexMap)
	if !ok {
		return nil, false
	}
	return ims, true
}

// AddIndex adds the specified index to the schema.
func (s *schema) AddIndex(idx Index) {
	ims, ok := s.indexMpas()
	if !ok {
		return
	}
	im, ok := idx.Data().(indexMap)
	if !ok {
		return
	}
	s.data[schemaIndexesIdx] = append(ims, im)
	// Add index to cache
	s.indexes = append(s.indexes, idx)
}

// Indexes returns the schema indexes.
func (s *schema) Indexes() []Index {
	return s.indexes
}

// FindIndex returns the schema index by the spacified name.
func (s *schema) FindIndex(name string) (Index, error) {
	idxes := s.indexes
	for _, idx := range idxes {
		if idx.Name() == name {
			return idx, nil
		}
	}
	return nil, newIndexNotExistErrorr(name)
}

// PrimaryIndex returns the schema primary index.
func (s *schema) PrimaryIndex() (Index, error) {
	idxes := s.indexes
	for _, idx := range idxes {
		if idx.Type() == PrimaryIndex {
			return idx, nil
		}
	}
	return nil, newPrimaryIndexNotExistErrorr()
}

// Data returns the raw representation data in memory.
func (s *schema) Data() any {
	return s.data
}
