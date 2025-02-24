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
	"errors"

	"github.com/cybergarage/go-sqlparser/sql/query"
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
)

// NewDocumentKeyForSchema returns a key for the specified schema.
func NewDocumentKeyForSchema(dbName string, schema document.Schema, colName string, colVal any) (document.Key, error) {
	keyVal, err := document.NewValueForSchema(schema, colName, colVal)
	if err != nil {
		return nil, err
	}
	return document.NewKeyWith(dbName, schema.Name(), colName, keyVal), nil
}

// NewDocumentKeyFromIndexes returns a key for the specified index.
func NewDocumentKeyFromIndexes(dbName string, colName string, objMap document.MapObject, idxes ...document.Index) (document.Key, error) {
	objKey := document.NewKey()
	objKey = append(objKey, dbName)
	objKey = append(objKey, colName)
	for _, idx := range idxes {
		for _, elem := range idx.Elements() {
			name := elem.Name()
			v, ok := objMap[name]
			if !ok {
				return nil, newErrObjectInvalid(objMap)
			}
			objKey = append(objKey, name)
			objKey = append(objKey, v)
		}
	}
	return objKey, nil
}

// NewDocumentKeyFromObject returns a key from the specified object.
func NewDocumentKeyFromObject(dbName string, schema document.Schema, obj document.MapObject) (document.Key, error) {
	firstElementAsIndex := func(schema document.Schema) (document.Index, error) {
		elems := schema.Elements()
		if len(elems) < 1 {
			return nil, document.NewErrPrimaryIndexNotExist()
		}
		idx := document.NewIndex()
		idx.SetType(document.PrimaryIndex)
		idx.AddElement(elems[0])
		return idx, nil
	}

	prIdx, err := schema.PrimaryIndex()
	if err != nil {
		if !errors.Is(err, document.ErrNotExist) {
			return nil, err
		}
		// Use the first element as the primary index
		firstElemIdx, err := firstElementAsIndex(schema)
		if err != nil {
			return nil, err
		}
		prIdx = firstElemIdx
	}

	return NewDocumentKeyFromIndexes(dbName, schema.Name(), obj, prIdx)
}

// NewDocumentKeyFromCond returns a key for the specified condition.
func NewDocumentKeyFromCond(dbName string, schema document.Schema, cond query.Condition) (document.Key, document.IndexType, error) {
	if !cond.HasConditions() {
		return document.NewKeyWith(dbName, schema.Name()), document.PrimaryIndex, nil
	}
	prIdx, err := schema.PrimaryIndex()
	if err != nil {
		return nil, 0, err
	}

	expr := cond.Expr()
	switch expr := expr.(type) { //nolint: gocritic
	case *query.CmpExpr:
		colName := expr.Left().Name()
		colValue := expr.Right().Value()
		switch expr.Operator() { //nolint: exhaustive
		case query.EQ:
			prIdxType := document.SecondaryIndex
			if colName == prIdx.Name() {
				prIdxType = document.PrimaryIndex
			}
			key, err := NewDocumentKeyForSchema(dbName, schema, colName, colValue)
			if err != nil {
				return nil, 0, err
			}
			return key, prIdxType, nil
		default:
			return nil, 0, newErrQueryConditionNotSupported(cond)
		}
	}

	return nil, 0, newErrQueryConditionNotSupported(cond)
}
