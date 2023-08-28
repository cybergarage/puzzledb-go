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
	"github.com/cybergarage/puzzledb-go/puzzledb/coordinator"
	"github.com/hashicorp/go-memdb"
)

type Document struct {
	Key   string
	Value []byte
}

type transaction struct {
	coordinator.KeyCoder
	*memdb.Txn
}

// NewTransaction returns a new transaction.
func newTransactionWith(coder coordinator.KeyCoder, txn *memdb.Txn) coordinator.Transaction {
	return &transaction{
		KeyCoder: coder,
		Txn:      txn,
	}
}

// Exists returns true if the object for the specified key exists.
func (txn *transaction) Exists(key coordinator.Key) (coordinator.Object, bool) {
	obj, err := txn.Get(key)
	if err != nil || obj == nil {
		return nil, false
	}
	return obj, true
}

// Set sets the object for the specified key.
func (txn *transaction) Set(obj coordinator.Object) error {
	keyBytes, err := txn.KeyCoder.EncodeKey(obj.Key())
	if err != nil {
		return err
	}
	doc := &Document{
		Key:   string(keyBytes),
		Value: obj.Bytes(),
	}
	err = txn.Txn.Insert(tableName, doc)
	if err != nil {
		return err
	}

	return nil
}

// Get gets the object for the specified key.
func (txn *transaction) Get(key coordinator.Key) (coordinator.Object, error) {
	rs, err := txn.GetRange(key)
	if err != nil {
		return nil, err
	}
	if !rs.Next() {
		return nil, coordinator.NewKeyNotExistError(key)
	}
	return rs.Object(), nil
}

// GetRange gets the result set for the specified key.
func (txn *transaction) GetRange(key coordinator.Key, opts ...coordinator.Option) (coordinator.ResultSet, error) {
	keyBytes, err := txn.KeyCoder.EncodeKey(key)
	if err != nil {
		return nil, err
	}

	offset := uint(0)
	limit := int(-1)
	order := coordinator.OrderNone
	for _, opt := range opts {
		switch v := opt.(type) {
		case *coordinator.OffsetOption:
			offset = v.Offset
		case *coordinator.LimitOption:
			limit = v.Limit
		case *coordinator.OrderOption:
			order = v.Order
		}
	}

	var it memdb.ResultIterator
	if order != coordinator.OrderDesc {
		it, err = txn.Txn.Get(tableName, idName+prefix, string(keyBytes))
	} else {
		it, err = txn.Txn.GetReverse(tableName, idName+prefix, string(keyBytes))
	}
	if err != nil {
		return nil, err
	}

	return newResultSet(txn.KeyCoder, key, it, offset, limit), nil
}

// Remove removes the object for the specified key.
func (txn *transaction) Remove(key coordinator.Key) error {
	keyBytes, err := txn.KeyCoder.EncodeKey(key)
	if err != nil {
		return err
	}
	_, err = txn.Txn.DeleteAll(tableName, idName, string(keyBytes))
	if err != nil {
		return err
	}
	return nil
}

// Truncate removes all objects.
func (txn *transaction) Truncate() error {
	_, err := txn.Txn.DeleteAll(tableName, idName)
	if err != nil {
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
