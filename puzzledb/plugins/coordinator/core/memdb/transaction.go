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
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/coordinator/core"
	"github.com/hashicorp/go-memdb"
)

type document struct {
	Key   string
	Value []byte
}

type memdbTransaction struct {
	*core.NotifyManager
	*memdb.Txn
}

// NewTransaction returns a new transaction.
func newTransactionWith(mgr *core.NotifyManager, txn *memdb.Txn) coordinator.Transaction {
	return &memdbTransaction{
		NotifyManager: mgr,
		Txn:           txn,
	}
}

// Set sets the object for the specified key.
func (txn *memdbTransaction) Set(obj coordinator.Object) error {
	hasObj := false
	coordObj, err := txn.Get(obj.Key())
	if err != nil && coordObj != nil {
		hasObj = true
	}

	keyStr, err := obj.Key().Encode()
	if err != nil {
		return err
	}
	objBytes, err := obj.Encode()
	if err != nil {
		return err
	}
	doc := &document{
		Key:   keyStr,
		Value: objBytes,
	}
	err = txn.Txn.Insert(tableName, doc)
	if err != nil {
		return err
	}

	var evt coordinator.Event
	if hasObj {
		evt = coordinator.NewEventWith(coordinator.ObjectUpdated, obj)
	} else {
		evt = coordinator.NewEventWith(coordinator.ObjectCreated, obj)
	}
	err = txn.NotifyManager.NofifyEvent(evt)
	if err != nil {
		return err
	}
	return nil
}

// Get gets the object for the specified key.
func (txn *memdbTransaction) Get(key coordinator.Key) (coordinator.Object, error) {
	rs, err := txn.Range(key)
	if err != nil {
		return nil, err
	}
	if !rs.Next() {
		return nil, nil
	}
	return rs.Object(), nil
}

// Range gets the resultset for the specified key range.
func (txn *memdbTransaction) Range(key coordinator.Key) (coordinator.ResultSet, error) {
	keyStr, err := key.Encode()
	if err != nil {
		return nil, err
	}
	it, err := txn.Txn.Get(tableName, idFieldName+prefix, keyStr)
	if err != nil {
		return nil, err
	}
	return newResultSet(key, it), nil
}

// Delete deletes the object for the specified key.
func (txn *memdbTransaction) Delete(key coordinator.Key) error {
	return nil
}

// Commit commits this transaction.
func (txn *memdbTransaction) Commit() error {
	txn.Txn.Commit()
	return nil
}

// Cancel cancels this transaction.
func (txn *memdbTransaction) Cancel() error {
	txn.Txn.Abort()
	return nil
}
