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

type rangeResultSet struct {
	kv.Key
	fdb.RangeResult
	obj *kv.Object
	*fdb.RangeIterator
}

func newRangeResultSet(key kv.Key, rs fdb.RangeResult) kv.ResultSet {
	return &rangeResultSet{
		Key:           key,
		RangeResult:   rs,
		RangeIterator: rs.Iterator(),
		obj:           nil}
}

func (rs *rangeResultSet) Next() bool {
	if !rs.RangeIterator.Advance() {
		return false
	}
	irs, err := rs.RangeIterator.Get()
	if err != nil {
		return false
	}
	rs.obj = &kv.Object{
		Key:   rs.Key,
		Value: irs.Value,
	}
	return true
}

func (rs *rangeResultSet) Object() *kv.Object {
	return rs.obj
}

// Range returns a result set of the specified key.
func (txn *transaction) Range(key kv.Key) (kv.ResultSet, error) {
	keyBytes, err := key.Encode()
	if err != nil {
		return nil, err
	}
	r := fdb.SelectorRange{
		Begin: fdb.FirstGreaterOrEqual(fdb.Key(keyBytes)),
		End:   fdb.LastLessThan(fdb.Key(keyBytes)),
	}
	ro := fdb.RangeOptions{
		Limit:   0,
		Mode:    fdb.StreamingModeIterator,
		Reverse: false,
	}
	rs := txn.Transaction.GetRange(r, ro)
	return newRangeResultSet(key, rs), nil
}
