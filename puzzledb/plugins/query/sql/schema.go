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
	"github.com/cybergarage/go-sqlparser/sql"
	"github.com/cybergarage/go-sqlparser/sql/query"
	"github.com/cybergarage/go-sqlparser/sql/query/response/resultset"
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
)

// NewDocumentSchemaFrom creates a new schema from the specified schema object.
func NewDocumentSchemaFrom(stmt sql.CreateTable) (document.Schema, error) {
	s := document.NewSchema()
	s.SetName(stmt.TableName())
	// Add elements
	for _, col := range stmt.Schema().Columns() {
		e, err := NewDocumentElementFromColumn(col)
		if err != nil {
			return nil, err
		}
		err = s.AddElement(e)
		if err != nil {
			return nil, err
		}
	}
	// Add indexes
	for _, idx := range stmt.Schema().Indexes() {
		i, err := NewDocumentIndexFrom(s, idx)
		if err != nil {
			return nil, err
		}
		err = s.AddIndex(i)
		if err != nil {
			return nil, err
		}
	}
	// NOTE: Disable this check because the primary index is not set when the schema is created.
	// Check the primary index
	// if _, err := s.PrimaryIndex(); err != nil {
	// 	return nil, err
	// }
	return s, nil
}

// NewQuerySchemaFrom creates a new schema from the specified schema object.
func NewQuerySchemaFrom(doc document.Schema) (query.Schema, error) {
	columns := query.NewColumns()
	for _, elem := range doc.Elements() {
		column, err := NewQueryColumnFromElement(elem)
		if err != nil {
			return nil, err
		}
		columns = append(columns, column)
	}
	indexes := query.NewIndexes()
	for _, docIdx := range doc.Indexes() {
		idxType := query.SecondaryIndex
		if docIdx.Type() == document.PrimaryIndex {
			idxType = query.PrimaryIndex
		}
		idxColumns := query.NewColumns()
		for _, elem := range docIdx.Elements() {
			idxColumn, err := columns.LookupColumn(elem.Name())
			if err != nil {
				return nil, err
			}
			idxColumns = append(idxColumns, idxColumn)
		}
		indexes = append(indexes, query.NewIndexWith(docIdx.Name(), idxType, idxColumns))
	}
	return query.NewSchemaWith(
		doc.Name(),
		query.WithSchemaColumns(columns),
		query.WithSchemaIndexes(indexes),
	), nil
}

// NewResultSetSchemaFrom creates a resultset new schema from the specified schema object.
func NewResultSetSchemaFrom(dbName string, docSchema document.Schema, selectors query.Selectors) (resultset.Schema, error) {
	schema, err := NewQuerySchemaFrom(docSchema)
	if err != nil {
		return nil, err
	}
	return resultset.NewSchema(
		resultset.WithSchemaDatabaseName(dbName),
		resultset.WithSchemaTableSchema(schema),
		resultset.WithSchemaSelectors(selectors),
	), nil
}
