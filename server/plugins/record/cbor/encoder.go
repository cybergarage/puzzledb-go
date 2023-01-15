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

package cbor

import (
	"io"

	"github.com/cybergarage/go-cbor/cbor"
	"github.com/cybergarage/puzzledb-go/puzzledb/record"
)

// Encoder represents a CBOR encoder instance.
type Encoder struct {
	record.Encoder
}

// NewEncoder returns a new encorder instance.
func NewEncoder() *Encoder {
	return &Encoder{}

}

// Encode writes the specified object to the specified writer.
func (enc *Encoder) Encode(w io.Writer, obj record.Object) error {
	cbor := cbor.NewEncoder(w)
	return cbor.Encode(obj)
}
