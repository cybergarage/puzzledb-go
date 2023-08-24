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

// nolint: exhaustive
func NewElementTypeFrom(sqlType query.DataType) (document.ElementType, error) {
	switch sqlType {
	case query.TinyIntData:
		return document.Int8Type, nil
	case query.SmallIntData:
		return document.Int16Type, nil
	case query.IntData:
		return document.Int32Type, nil
	case query.BigIntData:
		return document.Int64Type, nil
	case query.FloatData:
		return document.Float32Type, nil
	case query.DoubleData:
		return document.Float64Type, nil
	case query.TextData, query.VarCharData:
		return document.StringType, nil
	case query.BlobData, query.VarBinaryData:
		return document.BinaryType, nil
	default:
		return 0, newErrNotSupported(sqlType.String())
	}
}

func NewDataTypeFrom(elemType document.ElementType) (query.DataType, error) {
	switch elemType {
	case document.Int8Type:
		return query.TinyIntData, nil
	case document.Int16Type:
		return query.SmallIntData, nil
	case document.Int32Type:
		return query.IntData, nil
	case document.Int64Type:
		return query.BigIntData, nil
	case document.Float32Type:
		return query.FloatData, nil
	case document.Float64Type:
		return query.DoubleData, nil
	case document.StringType:
		return query.TextData, nil
	case document.BinaryType:
		return query.BlobData, nil
	case document.ArrayType, document.MapType, document.DateTimeType, document.BoolType:
		return 0, newErrNotSupported(elemType)
	default:
		return 0, newErrNotSupported(elemType)
	}
}
