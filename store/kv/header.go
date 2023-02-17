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
	DatabaseType = HeaderType('D')
	SchemaType   = HeaderType('S')
	ObjectType   = HeaderType('O')
	IndexType    = HeaderType('I')
)

type Version uint8

const (
	V1 = Version(1)
)

type BinaryType uint8

const (
	CBOR = BinaryType(1)
)

var defaultObjectHeader = [2]uint8{uint8(ObjectType), uint8(uint8(CBOR) & uint8(V1<<4))}

// Header represents a header for any keys.
type Header [2]uint8

func NewObjectHeader() Header {
	return defaultObjectHeader
}
