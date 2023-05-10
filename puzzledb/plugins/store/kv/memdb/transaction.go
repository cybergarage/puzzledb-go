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
	"time"

	"github.com/cybergarage/puzzledb-go/puzzledb/document"
	"github.com/cybergarage/puzzledb-go/puzzledb/store/kv"
	"github.com/hashicorp/go-memdb"
)

// transaction represents a Memdb transaction instance.
type transaction struct {
	kv.Transaction
	*memdb.Txn
	document.KeyCoder
}

func newTransaction(txn *memdb.Txn, coder document.KeyCoder) *transaction {
	return &transaction{
		Txn:         txn,
		Transaction: nil,
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
	doc := &Document{
		Key:   string(keyBytes),
		Value: obj.Value,
	}
	mWriteLatency.Observe(float64(time.Since(now).Milliseconds()))
	return txn.Txn.Insert(tableName, doc)
}

// Get returns a key-value object of the specified key.
func (txn *transaction) Get(key kv.Key) (*kv.Object, error) {
	now := time.Now()
	keyBytes, err := txn.EncodeKey(key)
	if err != nil {
		return nil, err
	}
	it, err := txn.Txn.Get(tableName, idName, string(keyBytes))
	if err != nil {
		return nil, err
	}
	rs := newResultSet(txn.KeyCoder, it)
	if !rs.Next() {
		return nil, kv.NewObjectNotExistError(key)
	}
	mReadLatency.Observe(float64(time.Since(now).Milliseconds()))
	return rs.Object(), nil
}

// GetRange returns a result set of the specified key.
func (txn *transaction) GetRange(key kv.Key, opts ...kv.Option) (kv.ResultSet, error) {
	now := time.Now()
	keyBytes, err := txn.EncodeKey(key)
	if err != nil {
		return nil, err
	}
	it, err := txn.Txn.Get(tableName, idName+prefix, string(keyBytes))
	if err != nil {
		return nil, err
	}
	mRangeReadLatency.Observe(float64(time.Since(now).Milliseconds()))
	return newResultSet(txn.KeyCoder, it), nil
}

// Remove removes the specified key-value object.
func (txn *transaction) Remove(key kv.Key) error {
	obj, err := txn.Get(key)
	if err != nil {
		return err
	}
	keyBytes, err := txn.EncodeKey(obj.Key)
	if err != nil {
		return err
	}
	doc := &Document{
		Key:   string(keyBytes),
		Value: obj.Value,
	}
	err = txn.Txn.Delete(tableName, doc)
	if err != nil {
		if errors.Is(err, memdb.ErrNotFound) {
			return kv.NewObjectNotExistError(key)
		}
		return err
	}
	return nil
}

// RemoveRange removes the specified key-value object.
func (txn *transaction) RemoveRange(key kv.Key) error {
	keyBytes, err := txn.EncodeKey(key)
	if err != nil {
		return err
	}
	_, err = txn.Txn.DeleteAll(tableName, idName+prefix, string(keyBytes))
	if err != nil {
		if errors.Is(err, memdb.ErrNotFound) {
			return kv.NewObjectNotExistError(key)
		}
		return err
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
