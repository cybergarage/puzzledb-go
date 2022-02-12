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

package obj

import (
	"fmt"
	"time"
)

// Data represents a primitive value.
type Data interface {
	// GetType returns the primitive type.
	GetType() Type
	// GetData returns the value.
	GetData() interface{}
	// Equals returns true when the specified value is the same as this value, otherwise false.
	Equals(other Data) bool
	// Bytes returns the binary representation.
	Bytes() []byte
}

// NewDataWithValue reads a value from the specified value.
func NewDataWithValue(val interface{}) (Data, error) {
	var data Data

	switch v := val.(type) {
	case nil:
		data = NewNull()
	case bool:
		data = NewBoolWithValue(v)
	case string:
		data = NewStringWithValue(v)
	case int16:
		data = NewShortWithValue(v)
	case int32:
		data = NewIntWithValue(v)
	case int:
		data = NewLongWithValue(int64(v))
	case int64:
		data = NewLongWithValue(v)
	case float32:
		data = NewFloatWithValue(v)
	case float64:
		data = NewDoubleWithValue(v)
	case time.Time:
		// NOTE: Time is created as only a Timestamp. Datetime is not supported.
		data = NewTimestampWithValue(v)
	case []byte:
		data = NewBinaryWithValue(v)
	default:
		return nil, fmt.Errorf(errorInvalidValueType, val, val)
	}

	return data, nil
}

// NewDataWithBytes returns a data from the specified bytes.
func NewDataWithBytes(src []byte) (Data, []byte, error) {
	var err error
	var Type Type

	Type, src, err = ReadTypeBytes(src)
	if err != nil {
		return nil, nil, err
	}

	var data Data
	switch Type {
	case NULL:
		src, err = ReadNullBytes(src)
		if err != nil {
			return nil, src, err
		}
		data = NewNull()
	case BOOL:
		var v bool
		v, src, err = ReadBoolBytes(src)
		if err != nil {
			return nil, src, err
		}
		data = NewBoolWithValue(v)
	case STRING:
		var v string
		v, src, err = ReadStringBytes(src)
		if err != nil {
			return nil, src, err
		}
		data = NewStringWithValue(v)
	case SHORT:
		var v int16
		v, src, err = ReadInt16Bytes(src)
		if err != nil {
			return nil, src, err
		}
		data = NewShortWithValue(v)
	case INT:
		var v int32
		v, src, err = ReadInt32Bytes(src)
		if err != nil {
			return nil, src, err
		}
		data = NewIntWithValue(v)
	case LONG:
		var v int64
		v, src, err = ReadInt64Bytes(src)
		if err != nil {
			return nil, src, err
		}
		data = NewLongWithValue(v)
	case FLOAT:
		var v float32
		v, src, err = ReadFloat32Bytes(src)
		if err != nil {
			return nil, src, err
		}
		data = NewFloatWithValue(v)
	case DOUBLE:
		var v float64
		v, src, err = ReadFloat64Bytes(src)
		if err != nil {
			return nil, src, err
		}
		data = NewDoubleWithValue(v)
	case TIMESTAMP:
		var v time.Time
		v, src, err = ReadTimestampBytes(src)
		if err != nil {
			return nil, src, err
		}
		data = NewTimestampWithValue(v)
	case DATETIME:
		var v time.Time
		v, src, err = ReadDatetimeBytes(src)
		if err != nil {
			return nil, src, err
		}
		data = NewDatetimeWithValue(v)
	case BINARY:
		var v []byte
		v, src, err = ReadBinaryBytes(src)
		if err != nil {
			return nil, src, err
		}
		data = NewBinaryWithValue(v)
	}

	if data == nil {
		return nil, src, fmt.Errorf(errorUnknownType, Type)
	}

	return data, src, nil
}

// AppendDataBytes appends a value to the specified buffer.
func AppendDataBytes(buf []byte, val Data) []byte {
	buf = AppendTypeBytes(buf, val.GetType())
	return append(buf, val.Bytes()...)
}
