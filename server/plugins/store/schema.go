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

const (
	schemaNameIdx = 0
)

type schema struct {
	data map[uint8]any
}

// NewSchema returns a blank schema.
func NewSchema() store.Schema {
	return &schema{
		data: map[uint8]any{},
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

// AddElement adds the specified element to the schema.
func (s *schema) AddElement(elem store.Element) {
}

// Elements returns the schema elements.
func (s *schema) Elements() []store.Element {
	return []store.Element{}

}

// AddIndex adds the specified index to the schema.
func (s *schema) AddIndex(idx store.Index) {
}

// Elements returns the schema elements.
func (s *schema) Indexes() []store.Index {
	return []store.Index{}
}
