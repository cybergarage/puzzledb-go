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
	"github.com/cybergarage/puzzledb-go/puzzledb/coordinator"
)

// transaction represents a transaction instance.
type transaction struct {
	coordinator.KeyCoder
	fdb.Transaction
}

func newTransaction(coder coordinator.KeyCoder, txn fdb.Transaction) coordinator.Transaction {
	return &transaction{
		KeyCoder:    coder,
		Transaction: txn,
	}
}

// Set stores a key-value object. If the key already holds some value, it is overwritten.
func (txn *transaction) Set(obj coordinator.Object) error {
	keyBytes, err := txn.KeyCoder.EncodeKey(obj.Key())
	if err != nil {
		return err
	}
	objBytes, err := obj.Encode()
	if err != nil {
		return err
	}
	txn.Transaction.Set(fdb.Key(keyBytes), objBytes)
	return nil
}

// Get returns a key-value object of the specified key.
func (txn *transaction) Get(key coordinator.Key) (coordinator.Object, error) {
	keyBytes, err := txn.KeyCoder.EncodeKey(key)
	if err != nil {
		return nil, err
	}
	fbs := txn.Transaction.Get(fdb.Key(keyBytes))
	v, err := fbs.Get()
	if err != nil {
		return nil, err
	}
	// NOTE: FutureByteSlice::Get() doesn't return nil if the key doesn't exist.
	if len(v) == 0 {
		return nil, coordinator.NewKeyNotExistError(key)
	}
	val, err := coordinator.NewValueFrom(v)
	if err != nil {
		return nil, err
	}
	return coordinator.NewObjectWith(key, val), nil
}

// Remove removes the specified key-value object.
func (txn *transaction) Remove(key coordinator.Key) error {
	keyBytes, err := txn.KeyCoder.EncodeKey(key)
	if err != nil {
		return err
	}
	txn.Transaction.Clear(fdb.Key(keyBytes))
	return nil
}

// Truncate removes all objects.
func (txn *transaction) Truncate() error {
	// Get all keys in the transaction
	rangeQuery := fdb.KeyRange{
		Begin: fdb.Key(txn.KeyCoder.EncodeKey(coordinator.NewKey(""))),
		End:   fdb.Key(txn.KeyCoder.EncodeKey(coordinator.NewKey("\xff"))),
	}
	kvs, err := txn.Transaction.GetRange(rangeQuery, fdb.RangeOptions{}).GetSliceWithError()
	if err != nil {
		return err
	}

	// Remove all keys in the transaction
	for _, kv := range kvs {
		txn.Transaction.Clear(kv.Key)
	}

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
