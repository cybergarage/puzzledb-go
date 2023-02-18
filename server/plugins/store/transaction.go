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

type transaction struct {
	kv kv.Transaction
	document.Serializer
}

// InsertDocument puts a document object with the primary key.
func (txn *transaction) InsertDocument(key store.Key, obj store.Object) error {
	var b bytes.Buffer
	err := txn.Encode(&b, obj)
	if err != nil {
		return err
	}
	kvObj := kv.Object{
		Key:   kv.NewKeyWith(kv.DocumentKeyHeader, key),
		Value: b.Bytes(),
	}
	return txn.kv.Set(&kvObj)
}

// SelectDocuments gets document objects matching the specified key.
func (txn *transaction) SelectDocuments(key store.Key) ([]store.Object, error) {
	kvObjs, err := txn.kv.Get(kv.NewKeyWith(kv.DocumentKeyHeader, key))
	if err != nil {
		return nil, err
	}
	objs := []store.Object{}
	for _, kvObj := range kvObjs {
		obj, err := txn.Decode(bytes.NewReader(kvObj.Value))
		if err != nil {
			return nil, err
		}
		objs = append(objs, obj)
	}
	return objs, nil
}

// InsertDocument puts a document object with the specified primary key.
func (txn *transaction) InsertIndex(key store.Key, primeryKey store.Key) error {
	primeryKeyBytes, err := key.Encode()
	if err != nil {
		return err
	}
	kvObj := kv.Object{
		Key:   kv.NewKeyWith(kv.SecondaryIndexHeader, key),
		Value: primeryKeyBytes,
	}
	return txn.kv.Set(&kvObj)
}

// SelectDocumentsByIndex gets document objects matching the specified index key.
func (txn *transaction) SelectDocumentsByIndex(indexKey store.Key) ([]store.Object, error) {
	kvIdxObjs, err := txn.kv.Get(kv.NewKeyWith(kv.SecondaryIndexHeader, indexKey))
	if err != nil {
		return nil, err
	}
	objs := []store.Object{}
	for _, kvIdxObj := range kvIdxObjs {
		kvIdx, err := txn.Decode(bytes.NewReader(kvIdxObj.Value))
		if err != nil {
			return nil, err
		}
		kvObjs, err := txn.kv.Get([]any{kvIdx}) // kvIdx is already encoded
		if err != nil {
			return nil, err
		}
		for _, kvObj := range kvObjs {
			obj, err := txn.Decode(bytes.NewReader(kvObj.Value))
			if err != nil {
				return nil, err
			}
			objs = append(objs, obj)
		}
	}
	return objs, nil
}

// Commit commits this transaction.
func (txn *transaction) Commit() error {
	return txn.kv.Commit()
}

// Cancel cancels this transaction.
func (txn *transaction) Cancel() error {
	return txn.kv.Cancel()
}
