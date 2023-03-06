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
	"sync"

	"github.com/cybergarage/puzzledb-go/puzzledb/store"
	"github.com/cybergarage/puzzledb-go/puzzledb/store/kv"
)

// Databases represents a database map.
type Databases struct {
	sync.Map
}

func NewDatabases() *Databases {
	return &Databases{
		Map: sync.Map{},
	}
}

// CreateDatabase creates a new database.
func (dbs *Databases) CreateDatabase(name string) error {
	db, err := NewDatabaseWithID(name)
	if err != nil {
		return err
	}
	dbs.Map.Store(name, db)
	return nil
}

// GetDatabase retruns the specified database.
func (dbs *Databases) GetDatabase(id string) (kv.Database, error) {
	v, ok := dbs.Load(id)
	if !ok {
		return nil, store.NewDatabaseNotExistError(id)
	}
	db, ok := v.(*Database)
	if !ok {
		return nil, store.NewDatabaseNotExistError(id)
	}
	return db, nil
}

// RemoveDatabase removes the specified database.
func (dbs *Databases) RemoveDatabase(id string) error {
	_, ok := dbs.Load(id)
	if !ok {
		return nil
	}
	dbs.Delete(id)
	return nil
}
