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
	"github.com/cybergarage/puzzledb-go/puzzledb/server/plugins/store/kv/memdb"
)

type memdbCoordinator struct {
	*memdb.Database
}

// NewCoordinator returns a new etcd coordinator instance.
func NewCoordinator() core.CoordinatorService {
	return &memdbCoordinator{
		Database: nil,
	}
}

// AddObserver adds the observer to the coordinator.
func (coord *memdbCoordinator) AddObserver(key coordinator.Key, observer coordinator.Observer) error {
	return nil
}

func (coord *memdbCoordinator) Transact() (coordinator.Transaction, error) {
	txn, err := coord.Database.Transact(true)
	if err != nil {
		return nil, err
	}
	return newTransactionWith(txn), nil
}

// Start starts this etcd coordinator.
func (coord *memdbCoordinator) Start() error {
	db, err := memdb.NewDatabase()
	if err != nil {
		return err
	}
	coord.Database = db
	return nil
}

// Stop stops this etcd coordinator.
func (coord *memdbCoordinator) Stop() error {
	return nil
}
