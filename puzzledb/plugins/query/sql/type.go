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

// NewDocumentElementTypeFrom returns a document element type from the specified SQL data type.
func NewDocumentElementTypeFrom(sqlType query.DataType) (document.ElementType, error) {
	switch sqlType { // nolint: exhaustive
	case query.TinyIntType:
		return document.Int8Type, nil
	case query.SmallIntType:
		return document.Int16Type, nil
	case query.IntType:
		return document.Int32Type, nil
	case query.BigIntType:
		return document.Int64Type, nil
	case query.FloatType:
		return document.Float32Type, nil
	case query.DoubleType:
		return document.Float64Type, nil
	case query.TextType, query.VarCharType, query.CharType:
		return document.StringType, nil
	case query.BlobType, query.VarBinaryType:
		return document.BinaryType, nil
	case query.DateTimeType, query.TimeStampType:
		return document.DatetimeType, nil
	default:
		return 0, newErrNotSupported(sqlType.String())
	}
}

// NewQueryDataTypeFrom returns a new column with the specified element.
func NewQueryDataTypeFrom(elemType document.ElementType) (query.DataType, error) {
	switch elemType {
	case document.Int8Type:
		return query.TinyIntType, nil
	case document.Int16Type:
		return query.SmallIntType, nil
	case document.Int32Type:
		return query.IntType, nil
	case document.Int64Type:
		return query.BigIntType, nil
	case document.Float32Type:
		return query.FloatType, nil
	case document.Float64Type:
		return query.DoubleType, nil
	case document.StringType:
		return query.TextType, nil
	case document.BinaryType:
		return query.BlobType, nil
	case document.DatetimeType:
		return query.TimeStampType, nil
	case document.ArrayType, document.MapType, document.BoolType:
		return 0, newErrNotSupported(elemType)
	default:
		return 0, newErrNotSupported(elemType)
	}
}
