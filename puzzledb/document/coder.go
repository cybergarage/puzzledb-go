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

import "io"

// An Decoder reads encorded objects from the specified input stream.
type Decoder interface {
	// Decode returns the decorded object from the specified reader if available, otherwise returns an error.
	Decode(r io.Reader) (Object, error)
}

// An Encoder writes the specified object to the specified output stream.
type Encoder interface {
	// Encode writes the specified object to the specified writer.
	Encode(w io.Writer, obj Object) error
}

// A Coder includes decoder and encoder interfaces.
type Coder interface {
	Decoder
	Encoder
}
