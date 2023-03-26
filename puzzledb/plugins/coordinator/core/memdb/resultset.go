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
	"github.com/cybergarage/puzzledb-go/puzzledb/coordinator"
	"github.com/hashicorp/go-memdb"
)

// Memdb represents a Memdb instance.
type resultSet struct {
	it  memdb.ResultIterator
	key coordinator.Key
	obj coordinator.Object
}

func newResultSet(key coordinator.Key, it memdb.ResultIterator) coordinator.ResultSet {
	return &resultSet{
		it:  it,
		key: key,
		obj: nil,
	}
}

// Next moves the cursor forward next object from its current position.
func (rs *resultSet) Next() bool {
	elem := rs.it.Next()
	if elem == nil {
		return false
	}
	doc, ok := elem.(*document)
	if !ok {
		return false
	}
	key, err := coordinator.NewKeyFrom(doc.id)
	if err != nil {
		return false
	}
	val, err := coordinator.NewValueFrom(doc.value)
	if err != nil {
		return false
	}
	rs.obj = coordinator.NewObjectWith(key, val)
	return true
}

// Object returns an object in the current position.
func (rs *resultSet) Object() coordinator.Object {
	return rs.obj
}

// Objects returns all objects in the resultset.
func (rs *resultSet) Objects() []coordinator.Object {
	objs := []coordinator.Object{}
	for rs.Next() {
		objs = append(objs, rs.Object())
	}
	return objs
}
