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

type HeaderType byte

const (
	DatabaseObject = HeaderType('D')
	SchemaObject   = HeaderType('S')
	DocumentObject = HeaderType('O')
	IndexObject    = HeaderType('I')
)

type Version byte

const (
	V1 = Version(1)
)

type DocumentType byte

const (
	CBOR = DocumentType(1)
)

type IndexType byte

const (
	PrimaryIndex   = IndexType(1)
	SecondaryIndex = IndexType(2)
)

func headerByteFromVersion(v Version) byte {
	return (byte(v<<4) & 0x70)
}

func vertionFromHeaderByte(b byte) Version {
	return Version((b >> 4) & 0x07)
}

func typeFromHeaderByte(b byte) byte {
	return (b & 0x07)
}

var DatabaseKeyHeader = [2]byte{byte(DatabaseObject), byte(byte(CBOR) | headerByteFromVersion(V1))}
var SchemaKeyHeader = [2]byte{byte(SchemaObject), byte(byte(CBOR) | headerByteFromVersion(V1))}
var DocumentKeyHeader = [2]byte{byte(DocumentObject), byte(byte(CBOR) | headerByteFromVersion(V1))}
var PrimaryIndexHeader = [2]byte{byte(IndexObject), byte(byte(PrimaryIndex) | headerByteFromVersion(V1))}
var SecondaryIndexHeader = [2]byte{byte(IndexObject), byte(byte(SecondaryIndex) | headerByteFromVersion(V1))}

// KeyHeader represents a header for any keys.
type KeyHeader [2]byte

func (header KeyHeader) Type() HeaderType {
	return HeaderType(header[0])
}

func (header KeyHeader) Version() Version {
	return vertionFromHeaderByte(header[1])
}

func (header KeyHeader) DocumentType() DocumentType {
	return DocumentType(typeFromHeaderByte(header[1]))
}

func (header KeyHeader) IndexType() IndexType {
	return IndexType(typeFromHeaderByte(header[1]))
}

func (header KeyHeader) Bytes() []byte {
	return header[:]
}
