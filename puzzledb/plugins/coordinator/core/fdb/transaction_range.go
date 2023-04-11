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
	"github.com/cybergarage/puzzledb-go/puzzledb/coordinator"
)

type rangeResultSet struct {
	coordinator.Key
	fdb.RangeResult
	obj coordinator.Object
	*fdb.RangeIterator
}

func newRangeResultSet(key coordinator.Key, rs fdb.RangeResult) coordinator.ResultSet {
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
	rs.obj = coordinator.NewObjectWith(rs.Key, irs.Value)
	return true
}

func (rs *rangeResultSet) Object() coordinator.Object {
	return rs.obj
}

// Objects returns all objects in the resultset.
func (rs *rangeResultSet) Objects() []coordinator.Object {
	objs := []coordinator.Object{}
	for rs.Next() {
		objs = append(objs, rs.Object())
	}
	return objs
}

// GetRange returns a result set of the specified key.
func (txn *transaction) Range(key coordinator.Key) (coordinator.ResultSet, error) {
	keyBytes, err := key.Encode()
	if err != nil {
		return nil, err
	}
	r, err := fdb.PrefixRange(fdb.Key(keyBytes))
	if err != nil {
		return nil, err
	}
	ro := fdb.RangeOptions{
		Limit:   0,
		Mode:    fdb.StreamingModeIterator,
		Reverse: false,
	}
	rs := txn.Transaction.GetRange(r, ro)
	return newRangeResultSet(key, rs), nil
}
