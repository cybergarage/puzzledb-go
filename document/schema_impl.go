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
// 3: elements - []map[uint8]any
//    1: name - string
//    2: type - uint8

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

type schema struct {
	data    map[uint8]any
	indexes []Index
}

// NewSchema returns a blank schema.
func NewSchema() Schema {
	s := &schema{
		data:    map[uint8]any{},
		indexes: []Index{},
	}
	s.SetVersion(SchemaVersion)
	s.data[schemaElementsIdx] = []any{}
	s.data[schemaIndexesIdx] = []any{}
	return s
}

// NewSchemaWith creates a schema from the specified object.
func NewSchemaWith(obj any) (Schema, error) {
	s := NewSchema()
	return s, nil
}

// SetVersion sets the specified version to the schema.
func (s *schema) SetVersion(ver int) {
	s.data[schemaVersionIdx] = uint8(ver)
}

// Version returns the schema version.
func (s *schema) Version() int {
	v, ok := s.data[schemaVersionIdx]
	if !ok {
		return 0
	}
	switch ver := v.(type) {
	case uint8:
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
}

// Elements returns the schema elements.
func (s *schema) Elements() []Element {
	es := []Element{}
	ems, ok := s.elementMaps()
	if ok {
		for _, em := range ems {
			es = append(es, newElementWith(em))
		}
	}
	return es
}

// FindElement returns the schema elements by the name.
func (s *schema) FindElement(name string) (Element, error) {
	es := s.Elements()
	for _, e := range es {
		if e.Name() == name {
			return e, nil
		}
	}
	return nil, newErrorNotSupported(name)
}

// AddIndex adds the specified index to the schema.
func (s *schema) AddIndex(idx Index) {
	s.indexes = append(s.indexes, idx)
	v, ok := s.data[schemaElementsIdx]
	if !ok {
		return
	}
	a, ok := v.([]any)
	if !ok {
		return
	}
	s.data[schemaIndexesIdx] = append(a, idx.Data())
}

// Elements returns the schema elements.
func (s *schema) Indexes() []Index {
	return []Index{}
}

// Data returns the raw representation data in memory.
func (s *schema) Data() any {
	return s.data
}
