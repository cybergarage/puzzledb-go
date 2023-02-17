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

type HeaderType uint8

const (
	DatabaseObject = HeaderType('D')
	SchemaObject   = HeaderType('S')
	DocumentObject = HeaderType('O')
	IndexObject    = HeaderType('I')
)

type Version uint8

const (
	V1 = Version(1)
)

type BinaryType uint8

const (
	CBOR = BinaryType(1)
)

type IndexType uint8

const (
	PrimaryIndex   = BinaryType(1)
	SecondaryIndex = BinaryType(2)
)

func headerByteFromVersion(v Version) uint8 {
	return (uint8(v<<4) & 0x70)
}

func vertionFromHeaderByte(b uint8) Version {
	return Version((b >> 4) & 0x07)
}

func typeFromHeaderByte(b uint8) uint8 {
	return (b & 0x07)
}

var latestObjectKeyHeader = [2]uint8{uint8(DocumentObject), uint8(uint8(CBOR) | headerByteFromVersion(V1))}
var latestPrimaryIndexHeader = [2]uint8{uint8(IndexObject), uint8(uint8(PrimaryIndex) | headerByteFromVersion(V1))}
var latestSecondaryIndexHeader = [2]uint8{uint8(IndexObject), uint8(uint8(SecondaryIndex) | headerByteFromVersion(V1))}

// KeyHeader represents a header for any keys.
type KeyHeader [2]uint8

func NewDocumentKeyHeader() KeyHeader {
	return latestObjectKeyHeader
}

func (header KeyHeader) Type() HeaderType {
	return HeaderType(header[0])
}

func (header KeyHeader) Version() Version {
	return vertionFromHeaderByte(header[1])
}

func (header KeyHeader) IndexType() IndexType {
	return IndexType(typeFromHeaderByte(header[1]))
}
