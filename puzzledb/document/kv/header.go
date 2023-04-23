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

import "fmt"

type HeaderType byte

// DocumentType represents a document type.
type DocumentType byte

// KeyHeader represents a header for all keys.
type KeyHeader [2]byte

// Version represents a version.
type Version byte

// IndexType represents an index type.
type IndexType byte

// NewKeyHeader creates a new key header from the specified bytes.
func NewKeyHeaderFrom(b []byte) KeyHeader {
	var header KeyHeader
	copy(header[:], b)
	return header
}

// Type returns a header type.
func (header KeyHeader) Type() HeaderType {
	return HeaderType(header[0])
}

// Version returns a version.
func (header KeyHeader) Version() Version {
	return VertionFromHeaderByte(header[1])
}

// DocumentType returns a document type.
func (header KeyHeader) DocumentType() DocumentType {
	return DocumentType(TypeFromHeaderByte(header[1]))
}

// IndexType returns an index type.
func (header KeyHeader) IndexType() IndexType {
	return IndexType(TypeFromHeaderByte(header[1]))
}

// Bytes returns a byte array.
func (header KeyHeader) Bytes() []byte {
	return header[:]
}

// String returns a string.
func (header KeyHeader) String() string {
	return fmt.Sprintf("%c %02x", header.Type(), header[1])
}
