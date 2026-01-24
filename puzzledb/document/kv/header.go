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

package kv

import (
	"fmt"
)

// Category represents an object category.
type Category byte

// Format represents an object format.
type Format byte

// KeyHeader represents a header for all keys.
type KeyHeader [2]byte

// Version represents a format version.
type Version byte

// IndexFormat represents an index format.
type IndexFormat byte

// NewKeyHeader creates a new key header from the specified bytes.
func NewKeyHeaderFrom(b []byte) KeyHeader {
	var header KeyHeader
	copy(header[:], b)
	return header
}

// Category returns an object category.
func (header KeyHeader) Category() Category {
	return Category(header[0])
}

// Version returns a version.
func (header KeyHeader) Version() Version {
	return VertionFromHeaderByte(header[1])
}

// DocumentFormat returns a document format.
func (header KeyHeader) Format() Format {
	return Format(TypeFromHeaderByte(header[1]))
}

// IndexFormat returns an index format.
func (header KeyHeader) IndexFormat() IndexFormat {
	return IndexFormat(header.Format())
}

// Bytes returns a byte array.
func (header KeyHeader) Bytes() []byte {
	return header[:]
}

// String returns a string.
func (header KeyHeader) String() string {
	return fmt.Sprintf("%c %02x", header.Category(), header[1])
}
