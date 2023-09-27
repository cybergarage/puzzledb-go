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
	"time"

	"github.com/cybergarage/go-mysql/mysql"
	"github.com/cybergarage/go-safecast/safecast"
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
)

const (
	timestampFormat = "2006-01-02 15:04:05.999999"
)

func NewValueFrom(elem document.Element, val any) (mysql.Value, error) {
	et := elem.Type()
	var eb []byte
	switch et {
	case document.StringType:
		v, ok := val.(string)
		if !ok {
			return mysql.NewNullValue(), newDataTypeNotEqualError(val, et)
		}
		eb = []byte(v)
	case document.BinaryType:
		v, ok := val.([]byte)
		if !ok {
			return mysql.NewNullValue(), newDataTypeNotEqualError(val, et)
		}
		eb = v
	case document.BoolType:
		v, ok := val.(bool)
		if !ok {
			return mysql.NewNullValue(), newDataTypeNotEqualError(val, et)
		}
		if v {
			eb = []byte("1")
		} else {
			eb = []byte("0")
		}
	case document.Int8Type, document.Int16Type, document.Int32Type, document.Int64Type, document.Float32Type, document.Float64Type:
		eb = []byte(fmt.Sprintf("%v", val))
	case document.TimestampType:
		switch v := val.(type) {
		case time.Time:
			eb = []byte(v.Format(timestampFormat))
		default:
			var tv time.Time
			if err := safecast.ToTime(v, &tv); err == nil {
				eb = []byte(tv.Format(timestampFormat))
			} else {
				eb = []byte(fmt.Sprintf("%v", val))
			}
		}
	case document.ArrayType, document.MapType:
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
