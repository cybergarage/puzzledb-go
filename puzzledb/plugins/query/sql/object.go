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

// NewObjectFromInsert returns a new object from the specified schema and columns.
func NewObjectFromInsert(dbName string, schema document.Schema, stmt *query.Insert) (document.Key, document.MapObject, error) {
	obj := document.MapObject{}
	for _, col := range stmt.Columns() {
		colName := col.Name()
		elem, err := schema.FindElement(colName)
		if err != nil {
			return nil, nil, err
		}
		v, err := document.NewValueForType(elem.Type(), col.Value())
		if err != nil {
			return nil, nil, err
		}
		obj[colName] = v
	}

	objKey, err := NewKeyFromObject(dbName, schema, obj)
	if err != nil {
		return nil, nil, err
	}

	return objKey, obj, nil
}
