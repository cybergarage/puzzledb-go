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

	"github.com/cybergarage/puzzledb-go/puzzledb/store"
	"github.com/cybergarage/puzzledb-go/puzzledb/store/kv"
)

// InsertIndex puts a secondary index with the primary key.
func (txn *transaction) InsertIndex(idxKey store.Key, docKey store.Key) error {
	kvDocKey := kv.NewKeyWith(kv.DocumentKeyHeader, docKey)
	kvDocKeyBytes, err := kvDocKey.Encode()
	if err != nil {
		return err
	}
	var encDocKey bytes.Buffer
	err = txn.Encode(&encDocKey, kvDocKeyBytes)
	if err != nil {
		return err
	}
	kvIdxKey := kv.NewKeyWith(kv.SecondaryIndexHeader, idxKey)
	kvObj := kv.Object{
		Key:   kvIdxKey,
		Value: encDocKey.Bytes(),
	}
	return txn.kv.Set(&kvObj)
}

// RemoveIndex removes the specified secondary index.
func (txn *transaction) RemoveIndex(idxKey store.Key) error {
	kvIdxKey := kv.NewKeyWith(kv.SecondaryIndexHeader, idxKey)
	return txn.kv.Remove(kvIdxKey)
}

// FindDocumentsByIndex gets document objects matching the specified index key.
func (txn *transaction) FindDocumentsByIndex(idxKey store.Key) (store.ResultSet, error) {
	kvIdxKey := kv.NewKeyWith(kv.SecondaryIndexHeader, idxKey)
	kvIdxRs, err := txn.kv.GetRange(kvIdxKey)
	if err != nil {
		return nil, err
	}
	return newIndexResultSet(txn, txn.Serializer, kvIdxRs), nil
}
