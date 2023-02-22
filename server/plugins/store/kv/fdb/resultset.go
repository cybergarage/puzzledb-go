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

// Memdb represents a Memdb instance.
type resultSet struct {
	kv.Key
	fdb.RangeResult
	obj *kv.Object
	*fdb.RangeIterator
}

func newResultSet(key kv.Key, rs fdb.RangeResult) kv.ResultSet {
	return &resultSet{
		Key:           key,
		RangeResult:   rs,
		RangeIterator: rs.Iterator(),
		obj:           nil}
}

// Next moves the cursor forward next object from its current position.
func (rs *resultSet) Next() bool {
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

// Object returns an object in the current position.
func (rs *resultSet) Object() *kv.Object {
	return rs.obj
}
