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

// NewCollectionWith creates a new schema from the specified schema object.
func NewCollectionWith(schema *query.Schema) (document.Schema, error) {
	s := document.NewSchema()
	s.SetName(schema.TableName())
	// Columns
	for _, col := range schema.GetTableSpec().Columns {
		e, err := NewElementWith(col)
		if err != nil {
			return nil, err
		}
		err = s.AddElement(e)
		if err != nil {
			return nil, err
		}
		// Primary Index
		if col.Type.Options.KeyOpt == query.ColKeyPrimary {
			i, err := NewPrimaryIndexWith(e)
			if err != nil {
				return nil, err
			}
			err = s.AddIndex(i)
			if err != nil {
				return nil, err
			}
		}
	}
	// Indexes
	for _, idx := range schema.GetTableSpec().Indexes {
		i, err := NewIndexWith(s, idx)
		if err != nil {
			return nil, err
		}
		err = s.AddIndex(i)
		if err != nil {
			return nil, err
		}
	}
	// Primary index
	if _, err := s.PrimaryIndex(); err != nil {
		return nil, err
	}
	return s, nil
}

// NewAlterCollectionWith creates an alter schema from the specified schema object.
func NewAlterCollectionWith(schema *query.Schema) (document.Schema, error) {
	s := document.NewSchema()
	return s, nil
}
