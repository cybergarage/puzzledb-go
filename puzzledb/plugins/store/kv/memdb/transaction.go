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
func (txn *transaction) Set(obj kv.Object) error {
	now := time.Now()
	keyBytes, err := txn.EncodeKey(obj.Key())
	if err != nil {
		return err
	}
	doc := &Document{
		Key:   keyBytes,
		Value: obj.Value(),
	}
	mWriteLatency.Observe(float64(time.Since(now).Milliseconds()))
	return txn.Txn.Insert(tableName, doc)
}

// Get returns a key-value object of the specified key.
func (txn *transaction) Get(key kv.Key) (kv.Object, error) {
	now := time.Now()
	keyBytes, err := txn.EncodeKey(key)
	if err != nil {
		return nil, err
	}
	it, err := txn.Txn.Get(tableName, idName, keyBytes)
	if err != nil {
		return nil, err
	}
	rs := newResultSetWith(txn.KeyCoder, it, 0, kv.NoLimit)
	if !rs.Next() {
		return nil, kv.NewErrObjectNotExist(key)
	}
	mReadLatency.Observe(float64(time.Since(now).Milliseconds()))
	return rs.Object()
}

// GetRange returns a result set of the specified key.
func (txn *transaction) GetRange(key kv.Key, opts ...kv.Option) (kv.ResultSet, error) {
	now := time.Now()

	var err error
	keyBytes, err := txn.EncodeKey(key)
	if err != nil {
		return nil, err
	}

	offset := uint(0)
	limit := uint(0)
	order := kv.OrderNone
	for _, opt := range opts {
		switch v := opt.(type) {
		case kv.Offset:
			offset = uint(v)
		case kv.Limit:
			limit = uint(v)
		case kv.Order:
			order = v
		}
	}

	var it memdb.ResultIterator
	if order != kv.OrderDesc {
		it, err = txn.Txn.Get(tableName, idName+prefix, keyBytes)
	} else {
		it, err = txn.Txn.GetReverse(tableName, idName+prefix, keyBytes)
	}
	if err != nil {
		return nil, err
	}

	mRangeReadLatency.Observe(float64(time.Since(now).Milliseconds()))

	return newResultSetWith(txn.KeyCoder, it, offset, limit), nil
}

// Remove removes the specified key-value object.
func (txn *transaction) Remove(key kv.Key) error {
	obj, err := txn.Get(key)
	if err != nil {
		return err
	}
	keyBytes, err := txn.EncodeKey(obj.Key())
	if err != nil {
		return err
	}
	doc := &Document{
		Key:   keyBytes,
		Value: obj.Value(),
	}
	err = txn.Txn.Delete(tableName, doc)
	if err != nil {
		if errors.Is(err, memdb.ErrNotFound) {
			return kv.NewErrObjectNotExist(key)
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
	_, err = txn.Txn.DeleteAll(tableName, idName+prefix, keyBytes)
	if err != nil {
		if errors.Is(err, memdb.ErrNotFound) {
			return kv.NewErrObjectNotExist(key)
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

// SetTimeout sets the timeout of this transaction.
func (txn *transaction) SetTimeout(t time.Duration) error {
	return nil
}
