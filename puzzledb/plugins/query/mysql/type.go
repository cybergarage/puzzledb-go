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
)

// nolint: exhaustive
func elementTypeFromSQLType(sqlType query.ValType) (document.ElementType, error) {
	switch sqlType {
	case query.Int8:
		return document.Int8Type, nil
	case query.Int16:
		return document.Int16Type, nil
	case query.Int32:
		return document.Int32Type, nil
	case query.Int64:
		return document.Int64Type, nil
	case query.Float32:
		return document.Float32Type, nil
	case query.Float64:
		return document.Float64Type, nil
	case query.Text, query.VarChar:
		return document.StringType, nil
	case query.Blob:
		return document.BinaryType, nil
	default:
		return 0, newNotSupportedError(sqlType.String())
	}
}

func sqlTypeFromElementType(elemType document.ElementType) (query.ValType, error) {
	switch elemType {
	case document.Int8Type:
		return query.Int8, nil
	case document.Int16Type:
		return query.Int16, nil
	case document.Int32Type:
		return query.Int32, nil
	case document.Int64Type:
		return query.Int64, nil
	case document.Float32Type:
		return query.Float32, nil
	case document.Float64Type:
		return query.Float64, nil
	case document.StringType:
		return query.Text /* query.VarChar*/, nil
	case document.BinaryType:
		return query.Blob, nil
	case document.ArrayType, document.MapType, document.DateTimeType, document.BoolType:
		return 0, newNotSupportedError(elemType)
	default:
		return 0, newNotSupportedError(elemType)
	}
}
