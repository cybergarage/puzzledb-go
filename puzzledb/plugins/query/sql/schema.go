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

package sql

import (
	"github.com/cybergarage/go-sqlparser/sql/query"
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
)

// NewDocumentSchemaFrom creates a new schema from the specified schema object.
func NewDocumentSchemaFrom(stmt *query.CreateTable) (document.Schema, error) {
	s := document.NewSchema()
	s.SetName(stmt.TableName())
	// Add elements
	for _, col := range stmt.Schema().Columns() {
		e, err := NewDocumentElementFrom(col)
		if err != nil {
			return nil, err
		}
		s.AddElement(e)
	}
	// Add indexes
	for _, idx := range stmt.Schema().Indexes() {
		i, err := NewDocumentIndexWith(s, idx)
		if err != nil {
			return nil, err
		}
		s.AddIndex(i)
	}
	// Check the primary index
	if _, err := s.PrimaryIndex(); err != nil {
		return nil, err
	}
	return s, nil
}

// NewQuerySchemaFrom creates a new schema from the specified schema object.
func NewQuerySchemaFrom(col document.Schema) (*query.Schema, error) {
	columns := query.NewColumns()
	for _, elem := range col.Elements() {
		column, err := NewQueryColumnFrom(elem)
		if err != nil {
			return nil, err
		}
		columns = append(columns, column)
	}
	// indexes := query.NewIndexes()
	// for _, idx := range col.Indexes() {
	// 	idx, err := NewIndexFrom(idx)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	s.AddIndex(idx)
	// }
	return query.NewSchemaWith(
		col.Name(),
		query.WithSchemaColumns(columns),
		// query.WithSchemaIndexes(indexes),
	), nil
}
