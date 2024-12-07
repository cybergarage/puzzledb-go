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
	kvIdxRs kv.ResultSet
	doc     store.Document
	document.KeyDecoder
	document.Decoder
}

func newIndexResultSet(txn *transaction, keyDecoder document.KeyDecoder, docDecoder document.Decoder, rs kv.ResultSet) store.ResultSet {
	return &indexResultSet{
		txn:        txn,
		kvRs:       rs,
		kvIdxRs:    nil,
		doc:        nil,
		KeyDecoder: keyDecoder,
		Decoder:    docDecoder,
	}
}

// Next moves the cursor forward next object from its current resultset position.
func (rs *indexResultSet) Next() bool {
	// First, checks the current index resultset.
	if rs.kvIdxRs != nil {
		return rs.nextIndex()
	}

	// Next, checks the current resultset when the current index resultset is nil.
	if !rs.kvRs.Next() {
		return false
	}
	kvIdxObj := rs.kvRs.Object()
	kvIdx, err := rs.txn.DecodeKey(kvIdxObj.Value())
	if err != nil {
		return false
	}
	kvIdxRs, err := rs.txn.kv.GetRange(kvIdx)
	if err != nil {
		return false
	}
	rs.kvIdxRs = kvIdxRs
	if rs.nextIndex() {
		return true
	}
	return rs.Next()
}

// nextIndex moves the cursor forward next object from the current index resultset.
func (rs *indexResultSet) nextIndex() bool {
	if rs.kvIdxRs == nil {
		return false
	}
	if !rs.kvIdxRs.Next() {
		rs.kvIdxRs = nil
		return false
	}
	kvObj := rs.kvIdxRs.Object()
	if kvObj == nil {
		return false
	}
	obj, err := rs.txn.DecodeDocument(bytes.NewReader(kvObj.Value()))
	if err != nil {
		return false
	}
	rs.doc = store.NewDocument(kvObj.Key(), obj)
	return true
}

// Document returns the current object in the result set.
func (rs *indexResultSet) Document() (store.Document, error) {
	return rs.doc, nil
}
