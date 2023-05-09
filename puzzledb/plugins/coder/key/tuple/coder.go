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

package tuple

import (
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
	"github.com/cybergarage/puzzledb-go/puzzledb/config"
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

// SetConfig sets the specified configuration.
func (s *Coder) SetConfig(conf config.Config) {
}

// ServiceType returns the plug-in service type.
func (s *Coder) ServiceType() plugins.ServiceType {
	return plugins.CoderKeyService
}

// ServiceName returns the plug-in service name.
func (s *Coder) ServiceName() string {
	return "tuple"
}

// EncodeKey returns the encoded bytes from the specified key if available, otherwise returns an error.
func (s *Coder) EncodeKey(key document.Key) ([]byte, error) {
	tpl, err := newTupleWith(key)
	if err != nil {
		return nil, err
	}
	return tpl.Pack(), nil
}

// DecodeKey returns the decoded key from the specified bytes if available, otherwise returns an error.
func (s *Coder) DecodeKey(b []byte) (document.Key, error) {
	tpl, err := tuple.Unpack(b)
	if err != nil {
		return nil, err
	}
	return newKeyWith(tpl), nil
}

// Start starts this coder.
func (s *Coder) Start() error {
	return nil
}

// Stop stops this coder.
func (s Coder) Stop() error {
	return nil
}
