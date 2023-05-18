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
	"errors"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/cybergarage/puzzledb-go/puzzledb/coordinator"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/coordinator/core"
)

const RequiredAPIVersion int = 630

type Coordinator struct {
	*core.BaseCoordinator
	fdb.Database
}

// NewCoordinator returns a new etcd coordinator instance.
func NewCoordinator() core.CoordinatorService {
	return &Coordinator{ //nolint:all
		BaseCoordinator: core.NewBaseCoordinator(),
	}
}

// ServiceName returns the plug-in service name.
func (coord *Coordinator) ServiceName() string {
	return "fdb"
}

func (coord *Coordinator) Transact() (coordinator.Transaction, error) {
	txn, err := coord.Database.CreateTransaction()
	if err != nil {
		return nil, err
	}
	err = txn.Options().SetAccessSystemKeys()
	if err != nil {
		return nil, err
	}

	return newTransaction(coord.KeyCoder, txn), nil
}

// Start starts this etcd coordinator.
func (coord *Coordinator) Start() error {
	err := fdb.APIVersion(RequiredAPIVersion)
	if err != nil {
		return errors.Join(err, coord.Stop())
	}
	db, err := fdb.OpenDefault()
	if err != nil {
		return errors.Join(err, coord.Stop())
	}
	coord.Database = db
	return nil
}

// Stop stops this etcd coordinator.
func (coord *Coordinator) Stop() error {
	return nil
}
