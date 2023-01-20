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

package memdb

import (
	"bytes"

	"github.com/cybergarage/puzzledb-go/puzzledb/store"
	"github.com/hashicorp/go-memdb"
)

type document struct {
	Key   string
	Value []byte
}

// Memdb represents a Memdb instance.
type Transaction struct {
	store.Transaction
	*memdb.Txn
}

func newTransaction(txn *memdb.Txn) *Transaction {
	return &Transaction{
		Txn: txn,
	}
}

// Insert puts a key-value object.
func (txn *Transaction) Insert(obj *store.Object) error {
	keyBytes, err := obj.KeyBytes()
	if err != nil {
		return err
	}

	var writer bytes.Buffer

	doc := &document{
		Key:   string(keyBytes),
		Value: obj.Value,
	}
	return txn.Txn.Insert(tableName, doc)
}

// Select gets an key-value object of the specified key.
func (txn *Transaction) Select(key store.Key) (*store.Object, error) {
	keyBytes, err := store.KeyToBytes(key)
	if err != nil {
		return nil, err
	}
	it, err := txn.Get(tableName, idFieldName, string(keyBytes))
	if err != nil {
		return nil, err
	}
	elem := it.Next()
	if elem == nil {
		return nil, store.ObjectNotFound
	}
	doc, ok := elem.(*document)
	if !ok {
		return nil, store.ObjectNotFound
	}
	return &store.Object{
		Key:   key,
		Value: doc.Value,
	}, nil
}

// Commit commits this transaction.
func (txn *Transaction) Commit() error {
	txn.Txn.Commit()
	return nil
}

// Cancel cancels this transaction.
func (txn *Transaction) Cancel() error {
	txn.Txn.Abort()
	return nil
}
