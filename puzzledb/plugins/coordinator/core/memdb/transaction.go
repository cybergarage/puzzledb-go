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

type Document struct {
	Key   string
	Value []byte
}

type Txn struct {
	*core.NotifyManager
	*memdb.Txn
}

// NewTransaction returns a new transaction.
func newTransactionWith(mgr *core.NotifyManager, txn *memdb.Txn) coordinator.Transaction {
	return &Txn{
		NotifyManager: mgr,
		Txn:           txn,
	}
}

// Set sets the object for the specified key.
func (txn *Txn) Set(obj coordinator.Object) error {
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
	doc := &Document{
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
func (txn *Txn) Get(key coordinator.Key) (coordinator.Object, error) {
	rs, err := txn.Range(key)
	if err != nil {
		return nil, err
	}
	if !rs.Next() {
		return nil, coordinator.NewKeyNotExistError(key)
	}
	return rs.Object(), nil
}

// Range gets the resultset for the specified key range.
func (txn *Txn) Range(key coordinator.Key) (coordinator.ResultSet, error) {
	keyStr, err := key.Encode()
	if err != nil {
		return nil, err
	}
	it, err := txn.Txn.Get(tableName, idName+prefix, keyStr)
	if err != nil {
		return nil, err
	}
	return newResultSet(key, it), nil
}

// Delete deletes the object for the specified key.
func (txn *Txn) Delete(key coordinator.Key) error {
	keyBytes, err := key.Encode()
	if err != nil {
		return err
	}
	_, err = txn.Txn.DeleteAll(tableName, idName, string(keyBytes))
	if err != nil {
		return err
	}
	return nil
}

// Commit commits this transaction.
func (txn *Txn) Commit() error {
	txn.Txn.Commit()
	return nil
}

// Cancel cancels this transaction.
func (txn *Txn) Cancel() error {
	txn.Txn.Abort()
	return nil
}
