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
	store "github.com/cybergarage/puzzledb-go/puzzledb/store/kv"
)

// Transaction represents a transaction instance.
type Transaction struct {
	store.Transaction
}

// Set stores a key-value object. If the key already holds some value, it is overwritten.
func (txn *Transaction) Set(obj *store.Object) error {
	return nil
}

// Get returns a key-value object of the specified key.
func (txn *Transaction) Get(key store.Key) ([]*store.Object, error) {
	return nil, nil
}

// Remove removes the specified key-value object.
func (txn *Transaction) Remove(key store.Key) error {
	return nil
}

// Commit commits this transaction.
func (txn *Transaction) Commit() error {
	return nil
}

// Cancel cancels this transaction.
func (txn *Transaction) Cancel() error {
	return nil
}
