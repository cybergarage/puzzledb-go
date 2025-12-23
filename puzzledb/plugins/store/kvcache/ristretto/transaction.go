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
	"github.com/cybergarage/puzzledb-go/puzzledb/store/kv"
)

type Transaction struct {
	kv.Transaction
	*Store
}

func NewTransaction(txn kv.Transaction, store *Store) kv.Transaction {
	return &Transaction{
		Transaction: txn,
		Store:       store,
	}
}

// Set stores a key-value object. If the key already holds some value, it is overwritten.
func (txn *Transaction) Set(obj kv.Object) error {
	key := obj.Key()
	if txn.IsRegisteredCacheKey(key) {
		err := txn.SetCache(obj)
		if err != nil {
			return err
		}
	}
	return txn.Transaction.Set(obj)
}

// Get returns a key-value object of the specified key.
func (txn *Transaction) Get(key kv.Key) (kv.Object, error) {
	if txn.IsRegisteredCacheKey(key) {
		txn.Store.IncrementRequestCount()
		mRequestTotal.Inc()
		kb, err := txn.EncodeKey(key)
		if err != nil {
			return nil, err
		}
		v, ok := txn.Cache.Get(kb)
		if ok {
			vb, ok := v.([]byte)
			if ok {
				return kv.NewObject(key, vb), nil
			}
			txn.Store.IncrementHitCount()
			mHitTotal.Inc()
			mHitRate.Set(txn.Store.CacheHitRate())
		}
		mHitRate.Set(txn.Store.CacheHitRate())
	}

	obj, err := txn.Transaction.Get(key)
	if err != nil {
		return nil, err
	}

	if txn.IsRegisteredCacheKey(key) {
		err := txn.SetCache(obj)
		if err != nil {
			return nil, err
		}
	}

	return obj, nil
}

// Scan returns the result set for the specified key.
func (txn *Transaction) Scan(key kv.Key, opts ...kv.Option) (kv.ResultSet, error) {
	return txn.Transaction.Scan(key, opts...)
}
