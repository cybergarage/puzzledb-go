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

package fdb

import (
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/cybergarage/puzzledb-go/puzzledb/store/kv"
)

// transaction represents a transaction instance.
type transaction struct {
	fdb.Transaction
}

func newTransaction(txn fdb.Transaction) kv.Transaction {
	return &transaction{
		Transaction: txn,
	}
}

// Set stores a key-value object. If the key already holds some value, it is overwritten.
func (txn *transaction) Set(obj *kv.Object) error {
	keyBytes, err := obj.KeyBytes()
	if err != nil {
		return err
	}
	txn.Transaction.Set(fdb.Key(keyBytes), obj.Value)
	return nil
}

// Get returns a result set of the specified key.
func (txn *transaction) Get(key kv.Key) (kv.ResultSet, error) {
	return txn.getRange((key))
}

// Remove removes the specified key-value object.
func (txn *transaction) Remove(key kv.Key) error {
	keyBytes, err := key.Encode()
	if err != nil {
		return err
	}
	txn.Transaction.Clear(fdb.Key(keyBytes))
	return nil
}

// Commit commits this transaction.
func (txn *transaction) Commit() error {
	err := txn.Transaction.Commit().Get()
	if err != nil {
		return err
	}
	return nil
}

// Cancel cancels this transaction.
func (txn *transaction) Cancel() error {
	txn.Transaction.Cancel()
	return nil
}
