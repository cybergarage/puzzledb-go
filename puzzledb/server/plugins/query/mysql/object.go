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

package mysql

import (
	"github.com/cybergarage/go-mysql/mysql/query"
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
)

// Object represents a database object.
type Object map[string]any

// NewObjectWith returns a new object from the specified object.
func NewObjectWith(obj any) (store.Object, error) {
	obj, ok := obj.(Object)
	if ok {
		return obj, nil
	}
	objMap, ok := obj.(map[any]any)
	if ok {
		obj := Object{}
		for key, val := range objMap {
			switch k := key.(type) {
			case string:
				obj[k] = val
			case []byte:
				obj[string(k)] = val
			default:
				return nil, newObjectInvalidError(obj)
			}
		}
		return obj, nil
	}
	return nil, newObjectInvalidError(obj)
}

// NewObjectFromInsert returns a new object from the specified schema and columns.
func NewObjectFromInsert(dbName string, schema document.Schema, stmt *query.Insert) (store.Key, store.Object, error) {
	prIdx, err := schema.PrimaryIndex()
	if err != nil {
		return nil, nil, err
	}
	prIdxName := prIdx.Name()

	cols, err := stmt.Columns()
	if err != nil {
		return nil, nil, err
	}

	var docKey store.Key
	doc := Object{}
	for _, col := range cols.Columns() {
		colName := col.Name()
		// Docment key
		if colName == prIdxName {
			prKey, err := NewKeyWith(dbName, schema.Name(), prIdxName, col.Value())
			if err != nil {
				return nil, nil, err
			}
			docKey = prKey
		}

		// Document data
		// TODO: Checks data types
		_, err := schema.FindElement(colName)
		if err != nil {
			return nil, nil, err
		}
		doc[colName] = col.Value()
	}

	// Checks primary key data
	if docKey == nil {
		return nil, nil, newPrimaryKeyDataNotExistError(prIdxName, doc)
	}

	return docKey, doc, nil
}
