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
	"fmt"

	"github.com/cybergarage/puzzledb-go/puzzledb/context"
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
	"github.com/cybergarage/puzzledb-go/puzzledb/store/kv"
)

type indexResultSet struct {
	txn      *transaction
	kvIdxKey store.Key
	kvIdxRs  kv.ResultSet
	document.KeyDecoder
	document.Decoder
}

func newIndexResultSet(txn *transaction, keyDecoder document.KeyDecoder, docDecoder document.Decoder, idxKey store.Key, rs kv.ResultSet) store.ResultSet {
	return &indexResultSet{
		txn:        txn,
		kvIdxKey:   idxKey,
		kvIdxRs:    rs,
		KeyDecoder: keyDecoder,
		Decoder:    docDecoder,
	}
}

// Next moves the cursor forward next object from its current resultset position.
func (rs *indexResultSet) Next() bool {
	if rs.kvIdxRs == nil {
		return false
	}
	if !rs.kvIdxRs.Next() {
		rs.kvIdxRs = nil
		return false
	}
	return true
}

// Document returns the current object in the result set.
func (rs *indexResultSet) Document() (store.Document, error) {
	kvIdxObj, err := rs.kvIdxRs.Object()
	if err != nil {
		return nil, err
	}
	kvIdxKey := kvIdxObj.Key()
	if len(kvIdxKey) < (len(rs.kvIdxKey) + (1 /* header */)) {
		return nil, fmt.Errorf("invalid index key: %v", kvIdxKey)
	}

	// Compose the document primary key from the kv index key.
	docPrKey := document.NewKeyWith(kvIdxKey[1], kvIdxKey[2])
	docPrKey = append(docPrKey, kvIdxKey[(len(rs.kvIdxKey)+(1 /* header */)):]...)

	// Find the document object by the primary key.
	docPrRs, err := rs.txn.FindObjects(context.NewContext(), docPrKey)
	if err != nil {
		return nil, err
	}
	if !docPrRs.Next() {
		return nil, fmt.Errorf("not found: %v", docPrKey)
	}
	docPrObj, err := docPrRs.Document()
	if err != nil {
		return nil, err
	}
	return docPrObj, nil
}

// Err returns the error, if any, that was encountered during iteration.
func (rs *indexResultSet) Err() error {
	if rs.kvIdxRs != nil {
		return rs.kvIdxRs.Err()
	}
	return nil
}

// Close closes the result set and releases any resources.
func (rs *indexResultSet) Close() error {
	if rs.kvIdxRs != nil {
		if err := rs.kvIdxRs.Close(); err != nil {
			return err
		}
		rs.kvIdxRs = nil
	}
	return nil
}
