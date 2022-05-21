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

// Timestamp represents a timestamp value.
type Timestamp struct {
	value time.Time
}

// NewTimestamp returns a timestamp instance.
func NewTimestamp() *Timestamp {
	return NewTimestampWithValue(time.Now())
}

// NewTimestampWithValue returns a timestamp instance with the specified value.
func NewTimestampWithValue(v time.Time) *Timestamp {
	return &Timestamp{value: v}
}

// NewTimestampWithBytes returns a timestamp instance with the specified bytes.
func NewTimestampWithBytes(src []byte) (*Timestamp, []byte, error) {
	v, src, err := ReadTimestampBytes(src)
	if err != nil {
		return nil, src, err
	}
	return &Timestamp{value: v}, src, nil
}

// Type returns the object type.
func (v *Timestamp) Type() Type {
	return TIMESTAMP
}

// Value returns the object value.
func (v *Timestamp) Value() interface{} {
	return v.value
}

// SetValue sets a specified value.
func (v *Timestamp) SetValue(value time.Time) {
	v.value = value
}

// GetValue returns the stored value.
func (v *Timestamp) GetValue() time.Time {
	return v.value
}

// Equals returns true when the specified value is s the same as this value, otherwise false.
func (v *Timestamp) Equals(other Object) bool {
	if _, ok := other.(*Timestamp); !ok {
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
func (v *Timestamp) Bytes() []byte {
	return AppendTimestampBytes(nil, v.value)
}

// ReadTimestampBytes reads the specified bytes as a long Float.
func ReadTimestampBytes(src []byte) (time.Time, []byte, error) {
	v, src, err := ReadInt64Bytes(src)
	if err != nil {
		return time.Now(), src, err
	}
	return time.Unix(v/1e3, (v%1e3)*1e6), src, nil
}

// AppendTimestampBytes appends a value to the specified buffer.
func AppendTimestampBytes(buf []byte, v time.Time) []byte {
	return AppendInt64Bytes(buf, (v.Unix()*1000)+int64(v.Nanosecond()/1e6))
}
