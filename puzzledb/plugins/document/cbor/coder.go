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
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins"
)

// Coder represents a CBOR erializer.
type Coder struct {
}

// NewCoder returns a new CBOR erializer instance.
func NewCoder() *Coder {
	return &Coder{}
}

// ServiceType returns the plug-in service type.
func (s *Coder) ServiceType() plugins.ServiceType {
	return plugins.EncoderDocumentService
}

// ServiceName returns the plug-in service name.
func (s *Coder) ServiceName() string {
	return "cbor"
}

// Encode writes the specified object to the specified writer.
func (s *Coder) Encode(w io.Writer, obj document.Object) error {
	cbor := cbor.NewEncoder(w)
	return cbor.Encode(obj)
}

// Decode returns the decorded object from the specified reader if available, otherwise returns an error.
func (s *Coder) Decode(r io.Reader) (document.Object, error) {
	cbor := cbor.NewDecoder(r)
	return cbor.Decode()
}

// Start starts this coder.
func (s *Coder) Start() error {
	return nil
}

// Stop stops this coder.
func (s Coder) Stop() error {
	return nil
}
