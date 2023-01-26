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
	indexNameIdx = 1
	indexTypeIdx = 2
)

type index struct {
	data map[uint8]any
}

// NewIndex returns a blank index.
func NewIndex() store.Index {
	idx := &index{
		data: map[uint8]any{},
	}
	return idx
}

// Name returns the unique name.
func (idx *index) Name() string {
	return ""
}

// Type returns the index type.
func (idx *index) Type() store.IndexType {
	return 0
}

// Elements returns the schema elements.
func (idx *index) Elements() []store.Element {
	return []store.Element{}
}
