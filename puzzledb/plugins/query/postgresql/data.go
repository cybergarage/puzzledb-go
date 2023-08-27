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
	"github.com/cybergarage/go-postgresql/postgresql/protocol/message"
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
)

// DataTypeFrom returns a data type of PostgreSQL from the specified query data type.
func DataTypeFrom(t document.ElementType) message.DataType {
	switch t { //nolint:exhaustive
	case document.Int8Type:
		return message.ByteaType
	case document.Int16Type:
		return message.Int2Type
	case document.Int32Type:
		return message.Int4Type
	case document.Int64Type:
		return message.Int8Type
	case document.BinaryType:
		return message.ByteaType
	case document.StringType:
		return message.TextType
	case document.Float64Type:
		return message.Float8Type
	case document.Float32Type:
		return message.Float4Type
	case document.DateTimeType:
		return message.TimestampType
	case document.BoolType:
		return message.BoolType
	}
	return 0
}
