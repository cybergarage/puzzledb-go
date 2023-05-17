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

package coordinator

const (
	V1 = Version(1)
)

const (
	CBOR = DocumentType(1)
)

const (
	StateObject   = HeaderType('S')
	MessageObject = HeaderType('M')
)

const (
	PrimaryIndex   = IndexType(1)
	SecondaryIndex = IndexType(2)
)

var StateObjectKeyHeader = [2]byte{byte(StateObject), byte(byte(CBOR) | HeaderByteFromVersion(V1))}
var MessageObjectKeyHeader = [2]byte{byte(MessageObject), byte(byte(CBOR) | HeaderByteFromVersion(V1))}

func HeaderByteFromVersion(v Version) byte {
	return (byte(v<<4) & 0x70)
}

func VertionFromHeaderByte(b byte) Version {
	return Version((b >> 4) & 0x07)
}

func TypeFromHeaderByte(b byte) byte {
	return (b & 0x07)
}
