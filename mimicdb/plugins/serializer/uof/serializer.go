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

package uof

import (
	"github.com/cybergarage/mimicdb/mimicdb/obj"
	"github.com/cybergarage/mimicdb/mimicdb/plugins/serializer"
)

// Memdb represents a Memdb instance.
type UOF struct {
	serializer.Serializer
}

// NewSerializer returns a new serializer instance.
func NewSerializer() *UOF {
	return &UOF{}
}

// Encode dumps a specified object to the byte array.
func (s *UOF) Encode(obj obj.Object) ([]byte, error) {
	return obj.Bytes(), nil
}

// Decode creates an object from the specified byte array.
func (s *UOF) Decode(b []byte) (obj.Object, error) {
	obj, _, err := obj.NewObjectWithBytes(b)
	return obj, err
}

// Start starts this serializer.
func (s *UOF) Start() error {
	return nil
}

// Stop stops this serializer.
func (s *UOF) Stop() error {
	return nil
}
