// Copyright (C) 2024 The PuzzleDB Authors.
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

package memdb

import (
	"errors"
)

// BinaryFieldIndexer is a custom field indexer for binary keys.
type BinaryFieldIndexer struct {
}

// FromArgs extracts the binary key from the arguments.
func (indexer *BinaryFieldIndexer) FromArgs(args ...any) ([]byte, error) {
	if len(args) < 1 {
		return nil, errors.New("invalid arguments")
	}
	_, bytes, err := indexer.FromObject(args[0])
	return bytes, err
}

// FromObject extracts the binary key from the object.
func (indexer *BinaryFieldIndexer) FromObject(obj any) (bool, []byte, error) {
	binKey, ok := obj.([]byte)
	if ok {
		return true, binKey, nil
	}
	doc, ok := obj.(*Document)
	if ok {
		return true, doc.Key, nil
	}
	return false, nil, nil
}

// PrefixFromArgs returns the prefix of the key.
func (indexer *BinaryFieldIndexer) PrefixFromArgs(args ...any) ([]byte, error) {
	if len(args) < 1 {
		return nil, errors.New("invalid arguments")
	}
	obj := args[0]
	binKey, ok := obj.([]byte)
	if ok {
		return binKey, nil
	}
	return nil, errors.New("invalid object")
}
