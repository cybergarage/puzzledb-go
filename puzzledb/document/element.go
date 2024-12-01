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

package document

type ElementType int8

const (
	// 0x0x (Reserved).
	// 0x1x (Reserved - Collection).
	ArrayType ElementType = 0x10
	MapType   ElementType = 0x11
	// 0x20 Integer.
	Int8Type  ElementType = 0x20
	Int16Type ElementType = 0x21
	Int32Type ElementType = 0x22
	Int64Type ElementType = 0x23
	// 0x30 StringType.
	StringType ElementType = 0x30
	BinaryType ElementType = 0x31
	// 0x40 Floating-point.
	Float32Type ElementType = 0x40
	Float64Type ElementType = 0x41
	// 0x70 Special.
	DatetimeType ElementType = 0x70
	BoolType     ElementType = 0x71
)

type Element interface {
	// Name returns the unique name.
	Name() string
	// Type returns the index type.
	Type() ElementType
	// Data returns the raw representation data in memory.
	Data() any
	// SetName sets the specified name to the element.
	SetName(name string) Element
	// SetType sets the specified type to the element.
	SetType(t ElementType) Element
}

// NewElementTypeWith returns an element type from the specified parameters.
func NewElementTypeWith(v any) (ElementType, error) {
	switch et := v.(type) {
	case ElementType:
		return et, nil
	case int8:
		return ElementType(et), nil
	case uint8:
		return ElementType(et), nil
	}
	return 0, newErrElementTypeInvalid(v)
}

// String represents the string representation.
func (et ElementType) String() string {
	switch et {
	case ArrayType:
		return "array"
	case MapType:
		return "map"
	case Int8Type:
		return "int8"
	case Int16Type:
		return "int16"
	case Int32Type:
		return "int32"
	case Int64Type:
		return "int64"
	case StringType:
		return "string"
	case BinaryType:
		return "binary"
	case Float32Type:
		return "float32"
	case Float64Type:
		return "float64"
	case DatetimeType:
		return "datetime"
	case BoolType:
		return "bool"
	}
	return ""
}
