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

import "strings"

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
	smap, ok := schemaMapFrom(obj)
	if !ok {
		return nil, newSchemaInvalidError(obj)
	}

	s := &schema{
		data:     smap,
		elements: []Element{},
		indexes:  []Index{},
	}

	return s, s.updateCashes()
}

func (s *schema) updateCashes() error {
	// Caches elements

	ems, ok := s.elementMaps()
	if !ok {
		return newElementMapNotExist()
	}

	s.elements = []Element{}
	for _, em := range ems {
		e, err := newElementWith(em)
		if err != nil {
			return err
		}
		s.elements = append(s.elements, e)
	}

	// Caches indexes

	ims, ok := s.indexMpas()
	if !ok {
		return newIndexMapNotExist()
	}

	s.indexes = []Index{}
	for _, im := range ims {
		i, err := newIndexWith(s, im)
		if err != nil {
			return err
		}
		s.indexes = append(s.indexes, i)
	}

	return nil
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
	ems, ok := schemaMapsFrom(v)
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

// DropElement drops the specified element from the schema.
func (s *schema) DropElement(name string) error {
	ems, ok := s.elementMaps()
	if !ok {
		return newElementMapNotExist()
	}
	for i, em := range ems {
		if strings.EqualFold(em[elementNameIdx].(string), name) {
			s.data[schemaElementsIdx] = append(ems[:i], ems[i+1:]...)
			return s.updateCashes()
		}
	}
	return newElementNotExistError(name)
}

// Elements returns the schema elements.
func (s *schema) Elements() Elements {
	return s.elements
}

// FindElement returns the schema elements by the specified name.
func (s *schema) FindElement(name string) (Element, error) {
	es := s.Elements()
	for _, e := range es {
		if strings.EqualFold(e.Name(), name) {
			return e, nil
		}
	}
	return nil, newElementNotExistError(name)
}

func (s *schema) indexMpas() ([]indexMap, bool) {
	v, ok := s.data[schemaIndexesIdx]
	if !ok {
		return nil, false
	}
	ims, ok := schemaMapsFrom(v)
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

// DropIndex drops the specified index from the schema.
func (s *schema) DropIndex(name string) error {
	ims, ok := s.indexMpas()
	if !ok {
		return newIndexMapNotExist()
	}
	for i, im := range ims {
		if strings.EqualFold(im[indexNameIdx].(string), name) {
			s.data[schemaIndexesIdx] = append(ims[:i], ims[i+1:]...)
			return s.updateCashes()
		}
	}
	return newIndexNotExistError(name)
}

// Indexes returns the schema indexes.
func (s *schema) Indexes() Indexes {
	return s.indexes
}

// FindIndex returns the schema index by the spacified name.
func (s *schema) FindIndex(name string) (Index, error) {
	idxes := s.indexes
	for _, idx := range idxes {
		if strings.EqualFold(idx.Name(), name) {
			return idx, nil
		}
	}
	return nil, newIndexNotExistError(name)
}

// PrimaryIndex returns the schema primary index.
func (s *schema) PrimaryIndex() (Index, error) {
	for _, idx := range s.indexes {
		if idx.Type() == PrimaryIndex {
			return idx, nil
		}
	}
	return nil, newPrimaryIndexNotExistErrorr()
}

// SecondaryIndexes returns the schema secondary indexes.
func (s *schema) SecondaryIndexes() (Indexes, error) {
	secIdxes := []Index{}
	for _, idx := range s.indexes {
		if idx.Type() != SecondaryIndex {
			continue
		}
		secIdxes = append(secIdxes, idx)
	}
	return secIdxes, nil
}

// Data returns the raw representation data in memory.
func (s *schema) Data() any {
	return s.data
}
