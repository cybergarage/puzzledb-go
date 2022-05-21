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
	"time"
)

// Datetime represents a datetime value.
type Datetime struct {
	value time.Time
}

// NewDatetime returns a datetime.
func NewDatetime() *Datetime {
	return NewDatetimeWithValue(time.Now())
}

// NewDatetimeWithValue returns a datetime instance with the specified value.
func NewDatetimeWithValue(v time.Time) *Datetime {
	return &Datetime{value: v}
}

// NewDatetimeWithBytes returns a datetime instance with the specified bytes.
func NewDatetimeWithBytes(src []byte) (*Datetime, []byte, error) {
	v, src, err := ReadDatetimeBytes(src)
	if err != nil {
		return nil, src, err
	}
	return &Datetime{value: v}, src, nil
}

// Type returns the object type.
func (v *Datetime) Type() Type {
	return DATETIME
}

// Value returns the object value.
func (v *Datetime) Value() interface{} {
	return v.value
}

// SetValue sets a specified value.
func (v *Datetime) SetValue(value time.Time) {
	v.value = value
}

// GetValue returns the stored value.
func (v *Datetime) GetValue() time.Time {
	return v.value
}

// Equals returns true when the specified value is s the same as this value, otherwise false.
func (v *Datetime) Equals(other Object) bool {
	if _, ok := other.(*Datetime); !ok {
		return false
	}
	otherValue, ok := other.Value().(time.Time)
	if !ok {
		return false
	}
	if v.value != otherValue {
		return false
	}
	return true
}

// Bytes returns the binary representation.
func (v *Datetime) Bytes() []byte {
	return AppendDatetimeBytes(nil, v.value)
}

// ReadDatetimeBytes reads the specified bytes as a Float.
func ReadDatetimeBytes(src []byte) (time.Time, []byte, error) {
	v, src, err := ReadInt64Bytes(src)
	if err != nil {
		return time.Now(), src, err
	}
	return time.Unix(v, 0), src, nil
}

// AppendDatetimeBytes appends a value to the specified buffer.
func AppendDatetimeBytes(buf []byte, val time.Time) []byte {
	return AppendInt64Bytes(buf, val.Unix())
}
