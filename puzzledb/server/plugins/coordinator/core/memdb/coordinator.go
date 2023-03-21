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
	"github.com/cybergarage/puzzledb-go/puzzledb/server/plugins/coordinator/core"
	"github.com/hashicorp/go-memdb"
)

const (
	tableName   = "coordinator"
	idFieldName = "id"
	prefix      = "_prefix"
)

type memdbCoordinator struct {
	*memdb.MemDB
}

// NewCoordinator returns a new etcd coordinator instance.
func NewCoordinator() core.CoordinatorService {
	return &memdbCoordinator{
		MemDB: nil,
	}
}

// AddObserver adds the observer to the coordinator.
func (coord *memdbCoordinator) AddObserver(key coordinator.Key, observer coordinator.Observer) error {
	return nil
}

func (coord *memdbCoordinator) Transact() (coordinator.Transaction, error) {
	return newTransactionWith(coord.MemDB.Txn(true)), nil
}

// Start starts this etcd coordinator.
func (coord *memdbCoordinator) Start() error {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			tableName: {
				Name: tableName,
				Indexes: map[string]*memdb.IndexSchema{
					idFieldName: {
						Name:   idFieldName,
						Unique: true,
						Indexer: &memdb.StringFieldIndex{Field: "Key",
							Lowercase: true},
					},
				},
			},
		},
	}
	memDB, err := memdb.NewMemDB(schema)
	if err != nil {
		return err
	}
	coord.MemDB = memDB
	return nil
}

// Stop stops this etcd coordinator.
func (coord *memdbCoordinator) Stop() error {
	return nil
}
