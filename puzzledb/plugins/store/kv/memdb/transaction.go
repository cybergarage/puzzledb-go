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
	"errors"

	"github.com/cybergarage/puzzledb-go/puzzledb/store/kv"
	"github.com/hashicorp/go-memdb"
)

// transaction represents a Memdb transaction instance.
type transaction struct {
	kv.Transaction
	*memdb.Txn
}

func newTransaction(txn *memdb.Txn) *transaction {
	return &transaction{
		Txn:         txn,
		Transaction: nil,
	}
}

// Set stores a key-value object. If the key already holds some value, it is overwritten.
func (txn *transaction) Set(obj *kv.Object) error {
	// FIXME: memdb.Transaction::Set() inserts a duplicate record instead of updating it.
	err := txn.Remove(obj.Key)
	if err != nil && !errors.Is(err, kv.ErrNotExist) {
		return err
	}
	// Sets new record
	keyBytes, err := obj.KeyBytes()
	if err != nil {
		return err
	}
	doc := &Document{
		Key:   string(keyBytes),
		Value: obj.Value,
	}
	return txn.Txn.Insert(tableName, doc)
}

// Get returns a result set of the specified key.
func (txn *transaction) Get(key kv.Key) (kv.ResultSet, error) {
	keyBytes, err := key.Encode()
	if err != nil {
		return nil, err
	}
	it, err := txn.Txn.Get(tableName, idName+prefix, string(keyBytes))
	if err != nil {
		return nil, err
	}
	return newResultSet(key, it), nil
}

// Remove removes the specified key-value object.
func (txn *transaction) Remove(key kv.Key) error {
	keyBytes, err := key.Encode()
	if err != nil {
		return err
	}
	doc := &Document{
		Key:   string(keyBytes),
		Value: nil,
	}
	err = txn.Txn.Delete(tableName, doc)
	if err != nil {
		return kv.NewObjectNotExistError(key)
	}
	return nil
}

// Commit commits this transaction.
func (txn *transaction) Commit() error {
	txn.Txn.Commit()
	return nil
}

// Cancel cancels this transaction.
func (txn *transaction) Cancel() error {
	txn.Txn.Abort()
	return nil
}
