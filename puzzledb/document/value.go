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

import (
	"github.com/cybergarage/go-safecast/safecast"
)

// NewValueForSchema returns a value for the specified schema.
func NewValueForSchema(schema Schema, name string, av any) (any, error) {
	col, err := schema.FindElement(name)
	if err != nil {
		return nil, err
	}
	return NewValueForType(col.Type(), av)
}

// NewValueForType returns a value for the specified element type.
func NewValueForType(et ElementType, av any) (any, error) {
	switch et { //nolint:exhaustive
	case StringType:
		var v string
		err := safecast.ToString(av, &v)
		return v, err
	case Int8Type:
		var v int8
		err := safecast.ToInt8(av, &v)
		return v, err
	case Int16Type:
		var v int16
		err := safecast.ToInt16(av, &v)
		return v, err
	case Int32Type:
		var v int32
		err := safecast.ToInt32(av, &v)
		return v, err
	case Int64Type:
		var v int64
		err := safecast.ToInt64(av, &v)
		return v, err
	case Float32Type:
		var v float32
		err := safecast.ToFloat32(av, &v)
		return v, err
	case Float64Type:
		var v float64
		err := safecast.ToFloat64(av, &v)
		return v, err
	case BoolType:
		var v bool
		err := safecast.ToBool(av, &v)
		return v, err
	}
	return av, nil
}
