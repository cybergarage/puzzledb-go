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
	CBOR = DocumentFormat(1)
)

const (
	StateHeaderObject   = ObjectCategory('S')
	MessageHeaderObject = ObjectCategory('M')
	JobHeaderObject     = ObjectCategory('J')
)

const (
	PrimaryIndex   = IndexFormat(1)
	SecondaryIndex = IndexFormat(2)
)

var (
	StateObjectKeyHeader   = [2]byte{byte(StateHeaderObject), byte(byte(CBOR) | V1.HeaderByte())}
	MessageObjectKeyHeader = [2]byte{byte(MessageHeaderObject), byte(byte(CBOR) | V1.HeaderByte())}
	JobObjectKeyHeader     = [2]byte{byte(JobHeaderObject), byte(byte(CBOR) | V1.HeaderByte())}
)

// HeaderTypes returns all header types.
func HeaderTypes() []ObjectCategory {
	return []ObjectCategory{
		StateHeaderObject,
		MessageHeaderObject,
		JobHeaderObject,
	}
}

// HeaderPrefixes returns all header prefixes.
func HeaderPrefixes() [][]byte {
	return [][]byte{
		StateObjectKeyHeader[:],
		MessageObjectKeyHeader[:],
		JobObjectKeyHeader[:],
	}
}
