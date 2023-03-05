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
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
	"github.com/cybergarage/puzzledb-go/puzzledb/server/plugins"
)

// Serializer represents a CBOR erializer.
type Serializer struct {
	plugins.Service
}

// NewSerializer returns a new CBOR erializer instance.
func NewSerializer() *Serializer {
	return &Serializer{}

}

// Encode writes the specified object to the specified writer.
func (s *Serializer) Encode(w io.Writer, obj document.Object) error {
	cbor := cbor.NewEncoder(w)
	return cbor.Encode(obj)
}

// Decode returns the decorded object from the specified reader if available, otherwise returns an error.
func (s *Serializer) Decode(r io.Reader) (document.Object, error) {
	cbor := cbor.NewDecoder(r)
	return cbor.Decode()
}

// Start starts this serializer.
func (s *Serializer) Start() error {
	return nil
}

// Stop stops this serializer.
func (s Serializer) Stop() error {
	return nil
}
