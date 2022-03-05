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
	"github.com/hashicorp/go-memdb"
)

// Memdb represents a Memdb instance.
type Memdb struct {
	*memdb.MemDB
}

type document struct {
	Key   []byte
	Value []byte
}

// New returns a new memdb service nstance.
func New() *Memdb {
	return &Memdb{
		MemDB: nil,
	}
}

// Start starts this memdb.
func (srvs *Memdb) Start() error {
	return nil
}

// Stop stops this memdb.
func (srvs Memdb) Stop() error {
	return nil
}
