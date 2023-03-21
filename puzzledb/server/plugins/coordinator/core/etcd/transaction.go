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

package etcd

import (
	"github.com/cybergarage/puzzledb-go/puzzledb/coordinator"
)

type etcdTransaction struct {
}

// NewTransaction returns a new transaction.
func NewTransaction() coordinator.Transaction {
	return &etcdTransaction{}
}

// Set sets the object for the specified key.
func (txn *etcdTransaction) Set(obj coordinator.Object) error {
	return nil
}

// Get gets the object for the specified key.
func (txn *etcdTransaction) Get(key coordinator.Key) (coordinator.Object, error) {
	return nil, nil
}

// Range gets the resultset for the specified key range.
func (txn *etcdTransaction) Range(key coordinator.Key) (coordinator.ResultSet, error) {
	return nil, nil
}

// Commit commits this transaction.
func (txn *etcdTransaction) Commit() error {
	return nil
}

// Cancel cancels this transaction.
func (txn *etcdTransaction) Cancel() error {
	return nil
}