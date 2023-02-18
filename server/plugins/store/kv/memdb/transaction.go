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
	"github.com/cybergarage/puzzledb-go/puzzledb/store/errors"
	store "github.com/cybergarage/puzzledb-go/puzzledb/store/kv"
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

// Set stores a key-value object. If the key already holds some value, it is overwritten
func (txn *Transaction) Set(obj *store.Object) error {
	keyBytes, err := obj.KeyBytes()
	if err != nil {
		return err
	}
	doc := &document{
		Key:   string(keyBytes),
		Value: obj.Value,
	}
	return txn.Txn.Insert(tableName, doc)
}

// Get return a key-value object of the specified key.
func (txn *Transaction) Get(key store.Key) (*store.Object, error) {
	keyBytes, err := key.Encode()
	if err != nil {
		return nil, err
	}
	it, err := txn.Txn.Get(tableName, idFieldName, string(keyBytes))
	if err != nil {
		return nil, err
	}
	elem := it.Next()
	if elem == nil {
		return nil, errors.ObjectNotFound
	}
	doc, ok := elem.(*document)
	if !ok {
		return nil, errors.ObjectNotFound
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
