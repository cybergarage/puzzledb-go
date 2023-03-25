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
	"fmt"

	"github.com/cybergarage/go-mysql/mysql"
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
)

func NewValueFrom(elem document.Element, val any) (mysql.Value, error) {
	et := elem.Type()
	var eb []byte
	switch et {
	case document.String:
		v, ok := val.(string)
		if !ok {
			return mysql.NewNullValue(), newDataTypeNotEqualError(val, et)
		}
		eb = []byte(v)
	case document.Binary:
		v, ok := val.([]byte)
		if !ok {
			return mysql.NewNullValue(), newDataTypeNotEqualError(val, et)
		}
		eb = v
	case document.Bool:
		v, ok := val.(bool)
		if !ok {
			return mysql.NewNullValue(), newDataTypeNotEqualError(val, et)
		}
		if v {
			eb = []byte("1")
		} else {
			eb = []byte("0")
		}
	case document.Int8, document.Int16, document.Int32, document.Int64, document.Float32, document.Float64:
		eb = []byte(fmt.Sprintf("%v", val))
	case document.DateTime:
		// TODO: Converts binary date format of MySQL protocol
		eb = []byte(fmt.Sprintf("%v", val))
	case document.Array, document.Map:
		return mysql.NewNullValue(), newNotSupportedError(et)
	default:
		return mysql.NewNullValue(), newNotSupportedError(et)
	}
	st, err := sqlTypeFromElementType(et)
	if err != nil {
		return mysql.NewNullValue(), err
	}
	// TODO: Consider simplifying sqltypes.NewValue() of vitess because the check is redundant.
	return mysql.NewValueWith(st, eb)
}
