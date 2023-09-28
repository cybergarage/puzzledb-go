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

// NewAlterAddColumnSchemaWith creates a new schema with the specified alter index.
func NewAlterAddIndexSchemaWith(schema document.Schema, addIndex *query.Index) (document.Schema, error) {
	newIndexType, err := NewDocumentIndexTypeFrom(addIndex.Type())
	if err != nil {
		return schema, err
	}

	newIndex := document.NewIndex()
	newIndex.SetName(addIndex.Name())
	newIndex.SetType(newIndexType)
	for _, indexColumn := range addIndex.Columns() {
		schemaElem, err := schema.FindElement(indexColumn.Name())
		if err != nil {
			return schema, err
		}
		newIndex.AddElement(schemaElem)
	}

	err = schema.AddIndex(newIndex)
	if err != nil {
		return schema, err
	}

	return schema, nil
}
