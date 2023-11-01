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
	"time"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
	"github.com/cybergarage/puzzledb-go/puzzledb/store/kv"
)

// transaction represents a transaction instance.
type transaction struct {
	fdb.Transaction
	document.KeyCoder
}

func newTransaction(txn fdb.Transaction, coder document.KeyCoder) kv.Transaction {
	return &transaction{
		Transaction: txn,
		KeyCoder:    coder,
	}
}

// Set stores a key-value object. If the key already holds some value, it is overwritten.
func (txn *transaction) Set(obj *kv.Object) error {
	now := time.Now()
	keyBytes, err := txn.EncodeKey(obj.Key)
	if err != nil {
		return err
	}
	txn.Transaction.Set(fdb.Key(keyBytes), obj.Value)
	mWriteLatency.Observe(float64(time.Since(now).Milliseconds()))
	return nil
}

// Get returns a key-value object of the specified key.
func (txn *transaction) Get(key kv.Key) (*kv.Object, error) {
	now := time.Now()
	keyBytes, err := txn.EncodeKey(key)
	if err != nil {
		return nil, err
	}
	fbs := txn.Transaction.Get(fdb.Key(keyBytes))
	val, err := fbs.Get()
	if err != nil {
		return nil, err
	}
	// NOTE: FutureByteSlice::Get() doesn't return nil if the key doesn't exist.
	if len(val) == 0 {
		return nil, kv.NewObjectNotExistError(key)
	}
	mReadLatency.Observe(float64(time.Since(now).Milliseconds()))
	return &kv.Object{
		Key:   key,
		Value: val,
	}, nil
}

// Remove removes the specified key-value object.
func (txn *transaction) Remove(key kv.Key) error {
	keyBytes, err := txn.EncodeKey(key)
	if err != nil {
		return err
	}
	txn.Transaction.Clear(fdb.Key(keyBytes))
	return nil
}

// RemoveRange removes the specified key-value objects.
func (txn *transaction) RemoveRange(key kv.Key) error {
	keyBytes, err := txn.EncodeKey(key)
	if err != nil {
		return err
	}
	r, err := fdb.PrefixRange(fdb.Key(keyBytes))
	if err != nil {
		return err
	}
	txn.Transaction.ClearRange(r)
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

// SetTimeout sets the timeout of this transaction.
func (txn *transaction) SetTimeout(t time.Duration) error {
	return txn.Options().SetTimeout(int64(t / time.Millisecond))
}
