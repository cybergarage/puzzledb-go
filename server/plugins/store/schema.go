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

type schema struct {
	data map[uint8]any
}

// NewSchema returns a blank schema.
func NewSchema() store.Schema {
	return &schema{
		data: map[uint8]any{},
	}
}

// Name returns the schema name.
func (s *schema) SetName(name string) {
}

// Name returns the schema name.
func (s *schema) Name() string {
	return ""
}

// Elements returns the schema elements.
func (s *schema) Elements() []store.Element {
	return []store.Element{}

}

// Elements returns the schema elements.
func (s *schema) Indexes() []store.Index {
	return []store.Index{}
}
