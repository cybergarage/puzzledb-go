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

package coordinator

import (
	"github.com/cybergarage/puzzledb-go/puzzledb/coordinator"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/coordinator/core"
)

type Service interface {
	coordinator.Coordinator
	plugins.Service
}

type serviceImpl struct {
	core.CoordinatorService
}

func NewServiceWith(c core.CoordinatorService) Service {
	return &serviceImpl{
		CoordinatorService: c,
	}
}

// Start starts this service.
func (service *serviceImpl) Start() error {
	return service.CoordinatorService.Start()
}

// Stop stops this service.
func (service *serviceImpl) Stop() error {
	return service.CoordinatorService.Stop()
}
