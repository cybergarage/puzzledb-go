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

// NewKeyWith returns a key from the specified parameters.
func NewKeyWith(dbName string, tblName string, keyName string, val any) (store.Key, error) {
	return document.NewKeyWith(dbName, tblName, keyName, val), nil
}

// NewKeyWithCond returns a key from the specified parameters.
func NewKeyWithCond(dbName string, schema document.Schema, cond *query.Condition) (store.Key, document.IndexType, error) {
	prIdx, err := schema.PrimaryIndex()
	if err != nil {
		return nil, 0, err
	}
	switch v := cond.Expr.(type) {
	case *query.ComparisonExpr:
		col, ok := v.Left.(*query.ColName)
		if !ok {
			return nil, 0, newQueryConditionNotSupportedError(cond)
		}
		val, ok := v.Right.(*query.Literal)
		if !ok {
			return nil, 0, newQueryConditionNotSupportedError(cond)
		}
		colName := col.Name.String()
		prIdxType := document.SecondaryIndex
		if colName == prIdx.Name() {
			prIdxType = document.PrimaryIndex
		}
		return document.NewKeyWith(dbName, schema.Name(), colName, val.Bytes()), prIdxType, nil
	case *query.RangeCond:
		return nil, 0, newQueryConditionNotSupportedError(cond)
	}
	return nil, 0, newQueryConditionNotSupportedError(cond)
}

// NewPrimaryKeyWith returns a key from the specified parameters.
func NewPrimaryKeyWith(dbName string, idx document.Index, obj document.Object) (store.Key, error) {
	return nil, nil
}
