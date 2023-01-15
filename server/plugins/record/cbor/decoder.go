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

// Decoder represents a CBOR decoder instance.
type Decoder struct {
	record.Decoder
}

// NewDecoder returns a new encorder instance.
func NewDecoder() *Decoder {
	return &Decoder{}

}

// Decode returns the decorded object from the specified reader if available, otherwise returns an error.
func (dec *Decoder) Decode(r io.Reader) (record.Object, error) {
	cbor := cbor.NewDecoder(r)
	return cbor.Decode()
}
