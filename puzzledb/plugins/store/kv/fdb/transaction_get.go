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
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/cybergarage/puzzledb-go/puzzledb/store/kv"
)

type resultSet struct {
	kv.Key
	fdb.FutureByteSlice
	obj *kv.Object
}

func newResultSet(key kv.Key, fbs fdb.FutureByteSlice) kv.ResultSet {
	return &resultSet{
		Key:             key,
		FutureByteSlice: fbs,
		obj:             nil}
}

func (rs *resultSet) Next() bool {
	if rs.FutureByteSlice == nil {
		return false
	}
	val, err := rs.FutureByteSlice.Get()
	if err != nil {
		return false
	}
	rs.obj = &kv.Object{
		Key:   rs.Key,
		Value: val,
	}
	rs.FutureByteSlice = nil
	return true
}

func (rs *resultSet) Object() *kv.Object {
	return rs.obj
}

func (txn *transaction) getone(key kv.Key) (kv.ResultSet, error) {
	keyBytes, err := key.Encode()
	if err != nil {
		return nil, err
	}
	fbs := txn.Transaction.Get(fdb.Key(keyBytes))
	return newResultSet(key, fbs), nil
}
