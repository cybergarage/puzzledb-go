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

type transaction struct {
}

// NewTransaction returns a new transaction.
func NewTransaction() coordinator.Transaction {
	return &transaction{}
}

// Set sets the object for the specified key.
func (txn *transaction) Set(obj coordinator.Object) error {
	return nil
}

// Get gets the object for the specified key.
func (txn *transaction) Get(key coordinator.Key) (coordinator.Object, error) {
	return nil, nil
}

// GetRange gets the result set for the specified key.
func (txn *transaction) GetRange(key coordinator.Key, opts ...coordinator.Option) (coordinator.ResultSet, error) {
	return nil, nil
}

// Remove removes the object for the specified key.
func (txn *transaction) Remove(key coordinator.Key) error {
	return nil
}

// Truncate removes all objects.
func (txn *transaction) Truncate() error {
	return nil
}

// Commit commits this transaction.
func (txn *transaction) Commit() error {
	return nil
}

// Cancel cancels this transaction.
func (txn *transaction) Cancel() error {
	return nil
}
