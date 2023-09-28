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
)

// NewAlterAddColumnSchemaWith creates a new schema with the specified alter index.
func NewAlterAddIndexSchemaWith(schema query.Schema, addIndex *query.Index) (query.Schema, error) {
	newIndexColums := query.NewColumns()
	for _, indexColumn := range addIndex.Columns() {
		schemaColum, err := schema.ColumnByName(indexColumn.Name())
		if err != nil {
			return schema, err
		}
		newIndexColums = append(newIndexColums, schemaColum)
	}

	newIndex := query.NewIndexWith(
		addIndex.Name(),
		addIndex.Type(),
		newIndexColums)

	err := schema.AddIndex(newIndex)
	if err != nil {
		return schema, err
	}

	return schema, nil
}
