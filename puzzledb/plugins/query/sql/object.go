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

// Object represents a database object.
type Object map[string]any

// NewObjectWith returns a new object from the specified object.
func NewObjectWith(anyObj any) (Object, error) {
	obj, ok := anyObj.(Object)
	if ok {
		return obj, nil
	}
	objMap, ok := anyObj.(map[any]any)
	if ok {
		obj := Object{}
		for key, val := range objMap {
			switch k := key.(type) {
			case string:
				obj[k] = val
			case []byte:
				obj[string(k)] = val
			default:
				return nil, newErrObjectInvalid(obj)
			}
		}
		return obj, nil
	}
	return nil, newErrObjectInvalid(obj)
}

// NewObjectFromInsert returns a new object from the specified schema and columns.
func NewObjectFromInsert(dbName string, schema document.Schema, stmt *query.Insert) (document.Key, Object, error) {
	obj := Object{}
	for _, col := range stmt.Columns() {
		colName := col.Name()
		// TODO: Checks data types
		_, err := schema.FindElement(colName)
		if err != nil {
			return nil, nil, err
		}
		obj[colName] = col.Value()
	}

	objKey, err := NewKeyFromObject(dbName, schema, obj)
	if err != nil {
		return nil, nil, err
	}

	return objKey, obj, nil
}