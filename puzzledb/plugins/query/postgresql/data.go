// Copyright (C) 2019 The PuzzleDB Authors.
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

package postgresql

import (
	"github.com/cybergarage/go-postgresql/postgresql/system"
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
)

// NewDataTypeFrom returns a data type from the specified query data type.
func NewDataTypeFrom(t document.ElementType) (*system.DataType, error) {
	return system.GetDataType(dataTypeOIDFrom(t))
}

func dataTypeOIDFrom(t document.ElementType) system.OID {
	switch t { //nolint:exhaustive
	case document.Int8Type:
		return system.Bytea
	case document.Int16Type:
		return system.Int2
	case document.Int32Type:
		return system.Int4
	case document.Int64Type:
		return system.Int8
	case document.BinaryType:
		return system.Bytea
	case document.StringType:
		return system.Text
	case document.Float64Type:
		return system.Float8
	case document.Float32Type:
		return system.Float4
	case document.DateTimeType:
		return system.Timestamp
	case document.BoolType:
		return system.Bool
	}
	return 0
}
