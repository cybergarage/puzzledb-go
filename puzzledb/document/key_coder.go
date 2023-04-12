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

// An KeyDecoder decodes the specified bytes.
type KeyDecoder interface {
	// Decode returns the decoded key from the specified bytes if available, otherwise returns an error.
	Decode([]byte) (Key, error)
}

// An KeyEncoder encodes the specified key.
type KeyEncoder interface {
	// Encode returns the encoded bytes from the specified key if available, otherwise returns an error.
	Encode(Key) ([]byte, error)
}

// A KeyCoder includes key decoder and encoder interfaces.
type KeyCoder interface {
	KeyDecoder
	KeyEncoder
}
