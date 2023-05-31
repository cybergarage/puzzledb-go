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

package ristretto

import (
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/store/kvcache"
	"github.com/cybergarage/puzzledb-go/puzzledb/store/kv"
)

type Transaction struct {
	kv.Transaction
	*kvcache.CacheConfig
}

func NewTransaction(txn kv.Transaction, config *kvcache.CacheConfig) kv.Transaction {
	return &Transaction{
		Transaction: txn,
		CacheConfig: config,
	}
}

// Set stores a key-value object. If the key already holds some value, it is overwritten.
func (txn *Transaction) Set(obj *kv.Object) error {
	return txn.Transaction.Set(obj)
}

// Get returns a key-value object of the specified key.
func (txn *Transaction) Get(key kv.Key) (*kv.Object, error) {
	return txn.Transaction.Get(key)
}

// GetRange returns a result set of the specified key.
func (txn *Transaction) GetRange(key kv.Key, opts ...kv.Option) (kv.ResultSet, error) {
	return txn.Transaction.GetRange(key, opts...)
}
