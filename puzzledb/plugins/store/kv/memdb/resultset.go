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

package memdb

import (
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
	"github.com/cybergarage/puzzledb-go/puzzledb/store/kv"
	"github.com/hashicorp/go-memdb"
)

// Memdb represents a Memdb instance.
type resultSet struct {
	document.KeyCoder
	it     memdb.ResultIterator
	obj    kv.Object
	offset uint
	limit  int
	nRead  uint
}

func newResultSetWith(coder document.KeyCoder, it memdb.ResultIterator, offset uint, limit int) kv.ResultSet {
	return &resultSet{
		KeyCoder: coder,
		it:       it,
		obj:      nil,
		offset:   offset,
		limit:    limit,
		nRead:    0,
	}
}

// Next moves the cursor forward next object from its current position.
func (rs *resultSet) Next() bool {
	if kv.NoLimit < rs.limit && uint(rs.limit) <= rs.nRead {
		return false
	}

	for rs.nRead < rs.offset {
		elem := rs.it.Next()
		if elem == nil {
			return false
		}
		rs.nRead++
	}

	elem := rs.it.Next()
	if elem == nil {
		return false
	}
	rs.nRead++

	doc, ok := elem.(*Document)
	if !ok {
		return false
	}
	key, err := rs.DecodeKey([]byte(doc.Key))
	if err != nil {
		return false
	}
	rs.obj = kv.NewObject(key, doc.Value)
	return true
}

// Object returns an object in the current position.
func (rs *resultSet) Object() (kv.Object, error) {
	return rs.obj, nil
}
