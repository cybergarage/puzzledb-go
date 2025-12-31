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
	obj kv.Object
	*fdb.RangeIterator
	document.KeyCoder
	offset uint
	limit  uint
	nRead  uint
}

func newRangeResultSetWith(coder document.KeyCoder, rs fdb.RangeResult, offset uint, limit uint) kv.ResultSet {
	return &rangeResultSet{
		KeyCoder:      coder,
		RangeResult:   rs,
		RangeIterator: rs.Iterator(),
		offset:        offset,
		limit:         limit,
		nRead:         0,
		obj:           nil}
}

// Next moves the cursor forward next object from its current position.
func (rs *rangeResultSet) Next() bool {
	if kv.NoLimit < rs.limit && rs.limit <= rs.nRead {
		return false
	}

	for rs.nRead < rs.offset {
		if !rs.RangeIterator.Advance() {
			return false
		}
		_, err := rs.RangeIterator.Get()
		if err != nil {
			return false
		}
		rs.nRead++
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
	rs.obj = kv.NewObject(key, irs.Value)
	return true
}

// Object returns the current object in the result set.
func (rs *rangeResultSet) Object() (kv.Object, error) {
	return rs.obj, nil
}

// Scan returns the result set for the specified key.
func (txn *transaction) Scan(key kv.Key, opts ...kv.Option) (kv.ResultSet, error) {
	now := time.Now()

	keyBytes, err := txn.EncodeKey(key)
	if err != nil {
		return nil, err
	}
	r, err := fdb.PrefixRange(fdb.Key(keyBytes))
	if err != nil {
		return nil, err
	}

	offset := uint(0)
	limit := uint(0)
	reverseOrder := false
	for _, opt := range opts {
		switch v := opt.(type) {
		case kv.Offset:
			offset = uint(v)
		case kv.Limit:
			limit = uint(v)
		case kv.Order:
			if v == kv.OrderDesc {
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

	return newRangeResultSetWith(txn.KeyCoder, rs, offset, limit), nil
}

// Close closes the result set and releases any resources.
func (rs *rangeResultSet) Close() error {
	return nil
}
