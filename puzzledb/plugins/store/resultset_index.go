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

package store

import (
	"bytes"

	"github.com/cybergarage/puzzledb-go/puzzledb/document"
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
	"github.com/cybergarage/puzzledb-go/puzzledb/store/kv"
)

type indexResultSet struct {
	txn     *transaction
	kvRs    kv.ResultSet
	obj     store.Object
	decoder document.Decoder
}

func newIndexResultSet(txn *transaction, decoder document.Decoder, rs kv.ResultSet) store.ResultSet {
	return &indexResultSet{
		txn:     txn,
		kvRs:    rs,
		obj:     nil,
		decoder: decoder,
	}
}

// Next moves the cursor forward next object from its current position.
func (rs *indexResultSet) Next() bool {
	if !rs.kvRs.Next() {
		return false
	}
	kvIdxObj := rs.kvRs.Object()
	kvIdx, err := rs.decoder.DecodeDocument(bytes.NewReader(kvIdxObj.Value))
	if err != nil {
		return false
	}
	kvRs, err := rs.txn.kv.GetRange([]any{kvIdx}) // kvIdx is already encoded
	if err != nil {
		return false
	}
	if !kvRs.Next() {
		return false
	}
	kvObj := kvRs.Object()
	if kvObj == nil {
		return false
	}
	obj, err := rs.txn.DecodeDocument(bytes.NewReader(kvObj.Value))
	if err != nil {
		return false
	}
	rs.obj = obj
	return true
}

// Object returns an object in the current position.
func (rs *indexResultSet) Object() store.Object {
	return rs.obj
}

// Objects returns all objects in the resultset.
func (rs *indexResultSet) Objects() []store.Object {
	objs := []store.Object{}
	for rs.Next() {
		objs = append(objs, rs.Object())
	}
	return objs
}
