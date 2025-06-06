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

// NewDocumentElementFromColumn returns a new element with the specified column.
func NewDocumentElementFromColumn(col query.Column) (document.Element, error) {
	t, err := NewDocumentElementTypeFrom(col.DataType())
	if err != nil {
		return nil, err
	}
	e := document.NewElement()
	e.SetName(col.Name())
	e.SetType(t)
	return e, nil
}

// NewQueryColumnFromElement returns a new column with the specified element.
func NewQueryColumnFromElement(elem document.Element) (query.Column, error) {
	dt, err := NewQueryDataTypeFrom(elem.Type())
	if err != nil {
		return nil, err
	}
	def := query.NewDataDef(dt, 0)
	return query.NewColumnWithOptions(
		query.WithColumnName(elem.Name()),
		query.WithColumnData(def),
	), nil
}
