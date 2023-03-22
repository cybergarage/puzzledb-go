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

package etcd

import (
	"github.com/cybergarage/puzzledb-go/puzzledb/coordinator"
	"github.com/cybergarage/puzzledb-go/puzzledb/server/plugins/coordinator/core"
)

type etcdCoordinator struct {
	*core.BaseCoordinator
}

// NewCoordinator returns a new etcd coordinator instance.
func NewCoordinator() core.CoordinatorService {
	return &etcdCoordinator{
		BaseCoordinator: core.NewBaseCoordinator(),
	}
}

func (coord *etcdCoordinator) Transact() (coordinator.Transaction, error) {
	return NewTransaction(), nil
}

// Start starts this etcd coordinator.
func (coord *etcdCoordinator) Start() error {
	return nil
}

// Stop stops this etcd coordinator.
func (coord *etcdCoordinator) Stop() error {
	return nil
}
