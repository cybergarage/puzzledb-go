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
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
)

// NewKeyWith returns a key from the specified parameters.
func NewKeyWith(dbName string, tblName string, keyName string, val any) (store.Key, error) {
	return document.NewKeyWith(dbName, tblName, keyName, val), nil
}

// NewKeyFromIndex returns a key for the specified index.
func NewKeyFromIndex(dbName string, schema document.Schema, idx document.Index, objMap document.MapObject) (store.Key, error) {
	objKey := document.NewKey()
	objKey = append(objKey, dbName)
	objKey = append(objKey, schema.Name())
	objKey = append(objKey, idx.Name())
	for _, elem := range idx.Elements() {
		name := elem.Name()
		v, ok := objMap[name]
		if !ok {
			return nil, newErrObjectInvalid(objMap)
		}
		objKey = append(objKey, v)
	}
	return objKey, nil
}

// NewKeyFromObject returns a key from the specified object.
func NewKeyFromObject(dbName string, schema document.Schema, obj document.MapObject) (store.Key, error) {
	prIdx, err := schema.PrimaryIndex()
	if err != nil {
		return nil, err
	}
	return NewKeyFromIndex(dbName, schema, prIdx, obj)
}

// NewKeyFromCond returns a key for the specified condition.
func NewKeyFromCond(dbName string, schema document.Schema, cond *query.Condition) (store.Key, document.IndexType, error) {
	if cond == nil {
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
			return document.NewKeyWith(dbName, schema.Name(), colName, colValue), prIdxType, nil
		default:
			return nil, 0, newErrQueryConditionNotSupported(cond)
		}
	}

	return nil, 0, newErrQueryConditionNotSupported(cond)
}
