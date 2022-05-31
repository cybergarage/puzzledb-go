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

// Object represents a primitive object.
type Object interface {
	// Type returns the object type.
	Type() Type
	// Value returns the object value.
	Value() interface{}
	// Equals returns true when the specified value is the same as this value, otherwise false.
	Equals(other Object) bool
	// Bytes returns the binary representation.
	Bytes() []byte
}

// NewObjectWithBytes returns a new object from the specified bytes.
func NewObjectWithBytes(src []byte) (Object, []byte, error) {
	var err error
	var objType Type

	objType, src, err = ReadTypeBytes(src)
	if err != nil {
		return nil, nil, err
	}

	return NewObjectWithTypeAndBytes(objType, src)
}

// NewObjectWithTypeAndBytes returns a new object from the specified obj type and bytes.
func NewObjectWithTypeAndBytes(objType Type, src []byte) (Object, []byte, error) {
	var err error

	var obj Object
	switch objType {
	case DICTIONARY:
		obj, src, err = ReadDictionaryBytes(src)
		if err != nil {
			return nil, src, err
		}
	case ARRAY:
		obj, src, err = NewArrayWithBytes(src)
		if err != nil {
			return nil, src, err
		}
	case NULL:
		src, err = ReadNullBytes(src)
		if err != nil {
			return nil, src, err
		}
		obj = NewNull()
	case BOOL:
		var v bool
		v, src, err = ReadBoolBytes(src)
		if err != nil {
			return nil, src, err
		}
		obj = NewBoolWithValue(v)
	case STRING:
		var v string
		v, src, err = ReadStringBytes(src)
		if err != nil {
			return nil, src, err
		}
		obj = NewStringWithValue(v)
	case TINY:
		var v int8
		v, src, err = ReadInt8Bytes(src)
		if err != nil {
			return nil, src, err
		}
		obj = NewTinyWithValue(v)
	case SHORT:
		var v int16
		v, src, err = ReadInt16Bytes(src)
		if err != nil {
			return nil, src, err
		}
		obj = NewShortWithValue(v)
	case INT:
		var v int32
		v, src, err = ReadInt32Bytes(src)
		if err != nil {
			return nil, src, err
		}
		obj = NewIntWithValue(v)
	case LONG:
		var v int64
		v, src, err = ReadInt64Bytes(src)
		if err != nil {
			return nil, src, err
		}
		obj = NewLongWithValue(v)
	case FLOAT:
		var v float32
		v, src, err = ReadFloat32Bytes(src)
		if err != nil {
			return nil, src, err
		}
		obj = NewFloatWithValue(v)
	case DOUBLE:
		var v float64
		v, src, err = ReadFloat64Bytes(src)
		if err != nil {
			return nil, src, err
		}
		obj = NewDoubleWithValue(v)
	case TIMESTAMP:
		var v time.Time
		v, src, err = ReadTimestampBytes(src)
		if err != nil {
			return nil, src, err
		}
		obj = NewTimestampWithValue(v)
	case DATETIME:
		var v time.Time
		v, src, err = ReadDatetimeBytes(src)
		if err != nil {
			return nil, src, err
		}
		obj = NewDatetimeWithValue(v)
	case BINARY:
		var v []byte
		v, src, err = ReadBinaryBytes(src)
		if err != nil {
			return nil, src, err
		}
		obj = NewBinaryWithValue(v)
	}

	if obj == nil {
		return nil, src, fmt.Errorf(errorUnknownType, objType)
	}

	return obj, src, nil
}

// AppendObjectBytes appends a value to the specified buffer.
func AppendObjectBytes(buf []byte, val Object) []byte {
	buf = AppendTypeBytes(buf, val.Type())
	return append(buf, val.Bytes()...)
}
