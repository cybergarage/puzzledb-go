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
	"github.com/cybergarage/puzzledb-go/puzzledb/store/kv"
)

// Memdb represents a Memdb instance.
type resultSet struct {
}

func newResultSet() kv.ResultSet {
	return &resultSet{}
}

// Next moves the cursor forward next object from its current position.
func (rs *resultSet) Next() bool {
	return false
}

// Object returns an object in the current position.
func (rs *resultSet) Object() *kv.Object {
	return nil
}
