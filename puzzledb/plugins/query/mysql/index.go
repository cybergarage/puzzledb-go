// Copyright (C) 2022 PuzzleDB Contributors.
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

package mysql

import (
	"github.com/cybergarage/go-mysql/mysql/query"
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
)

// NewPrimaryIndexWith creates an index from the specified element.
func NewPrimaryIndexWith(elem document.Element) (document.Index, error) {
	idx := document.NewIndex()
	idx.SetName(elem.Name())
	idx.SetType(document.PrimaryIndex)
	idx.AddElement(elem)
	return idx, nil
}

// NewIndexWith creates an index from the specified coulumn definition.
func NewIndexWith(s document.Schema, def *query.IndexDefinition) (document.Index, error) {
	if def.Info.Spatial || def.Info.Fulltext {
		return nil, newIndexNotSupportedError(def.Info.Type)
	}

	idx := document.NewIndex()

	idx.SetName(def.Info.Name.Lowered())

	if def.Info.Primary {
		idx.SetType(document.PrimaryIndex)
	} else {
		idx.SetType(document.SecondaryIndex)
	}

	for _, col := range def.Columns {
		elem, err := s.FindElement(col.Column.String())
		if err != nil {
			return nil, err
		}
		idx.AddElement(elem)
	}

	return idx, nil
}
