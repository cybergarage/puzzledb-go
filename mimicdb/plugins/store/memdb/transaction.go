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
	"github.com/cybergarage/mimicdb/mimicdb/errors"
	"github.com/cybergarage/mimicdb/mimicdb/plugins/store"
	"github.com/hashicorp/go-memdb"
)

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
	return txn.Txn.Insert(tableName, obj)
}

// Select gets an key-value object of the specified key.
func (txn *Transaction) Select(key Key) (*Object, error) {
	it, err := txn.Get(tableName, idFieldName, string(key))
	if err != nil {
		return nil, err
	}
	obj := it.Next()
	if obj == nil {
		return nil, errors.ObjectNotFound
	}
	return obj.(*Object), nil
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
