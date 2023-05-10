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
	"time"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
	"github.com/cybergarage/puzzledb-go/puzzledb/store/kv"
)

type rangeResultSet struct {
	fdb.RangeResult
	obj *kv.Object
	*fdb.RangeIterator
	document.KeyCoder
	limit int
	nRead int
}

func newRangeResultSetWith(coder document.KeyCoder, rs fdb.RangeResult, limit int) kv.ResultSet {
	return &rangeResultSet{
		KeyCoder:      coder,
		RangeResult:   rs,
		RangeIterator: rs.Iterator(),
		limit:         limit,
		nRead:         0,
		obj:           nil}
}

func (rs *rangeResultSet) Next() bool {
	if kv.NoLimit < rs.limit && rs.limit <= rs.nRead {
		return false
	}

	if !rs.RangeIterator.Advance() {
		return false
	}
	irs, err := rs.RangeIterator.Get()
	if err != nil {
		return false
	}
	rs.nRead++

	key, err := rs.DecodeKey(irs.Key)
	if err != nil {
		return false
	}
	rs.obj = &kv.Object{
		Key:   key,
		Value: irs.Value,
	}
	return true
}

func (rs *rangeResultSet) Object() *kv.Object {
	return rs.obj
}

// GetRange returns a result set of the specified key.
func (txn *transaction) GetRange(key kv.Key, opts ...kv.Option) (kv.ResultSet, error) {
	now := time.Now()

	keyBytes, err := txn.EncodeKey(key)
	if err != nil {
		return nil, err
	}
	r, err := fdb.PrefixRange(fdb.Key(keyBytes))
	if err != nil {
		return nil, err
	}

	limit := -1
	reverseOrder := false
	for _, opt := range opts {
		switch v := opt.(type) {
		case *kv.LimitOption:
			limit = v.Limit
		case *kv.OrderOption:
			if v.Order == kv.OrderDesc {
				reverseOrder = true
			}
		}
	}

	ro := fdb.RangeOptions{
		Limit:   0,
		Mode:    fdb.StreamingModeIterator,
		Reverse: reverseOrder,
	}
	rs := txn.Transaction.GetRange(r, ro)

	mReadLatency.Observe(float64(time.Since(now).Milliseconds()))

	return newRangeResultSetWith(txn.KeyCoder, rs, limit), nil
}
