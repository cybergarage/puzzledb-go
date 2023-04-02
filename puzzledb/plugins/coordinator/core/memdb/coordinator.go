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
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/coordinator/core"
	"github.com/hashicorp/go-memdb"
)

const (
	tableName   = "coordinator"
	idName      = "id"
	idFieldName = "Key"
	prefix      = "_prefix"
)

type memdbCoordinator struct {
	*core.BaseCoordinator
	*memdb.MemDB
}

// NewCoordinator returns a new etcd coordinator instance.
func NewCoordinator() core.CoordinatorService {
	return &memdbCoordinator{
		BaseCoordinator: core.NewBaseCoordinator(),
		MemDB:           nil,
	}
}

// Type returns the plug-in service type.
func (coord *memdbCoordinator) Type() plugins.ServiceType {
	return plugins.CoordinatorService
}

// Name returns the plug-in service name.
func (coord *memdbCoordinator) Name() string {
	return "memdb"
}

func (coord *memdbCoordinator) Transact() (coordinator.Transaction, error) {
	return newTransactionWith(coord.NotifyManager, coord.MemDB.Txn(true)), nil
}

// Start starts this etcd coordinator.
func (coord *memdbCoordinator) Start() error {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			tableName: {
				Name: tableName,
				Indexes: map[string]*memdb.IndexSchema{
					idName: {
						Name:         idName,
						AllowMissing: false,
						Unique:       true,
						Indexer: &memdb.StringFieldIndex{
							Field:     idFieldName,
							Lowercase: true,
						},
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
