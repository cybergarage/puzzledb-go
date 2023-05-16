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
	observers []coordinator.Observer
	coordinator.Process
	core.CoordinatorService
}

func NewServiceWith(c core.CoordinatorService) Service {
	return &serviceImpl{
		Process:            coordinator.NewProcess(),
		observers:          make([]coordinator.Observer, 0),
		CoordinatorService: c,
	}
}

// AddObserver adds the specified observer.
func (coord *serviceImpl) AddObserver(newObserver coordinator.Observer) error {
	for _, observer := range coord.observers {
		if observer == newObserver {
			return nil
		}
	}
	coord.observers = append(coord.observers, newObserver)
	return nil
}

// PostMessage posts the specified message to the coordinator.
func (coord *serviceImpl) PostMessage(msg coordinator.Message) error {
	return nil
}

// NofityMessage posts the specified message to the observers.
func (coord *serviceImpl) NofityMessage(msg coordinator.Message) error {
	for _, observer := range coord.observers {
		observer.MessageReceived(msg)
	}
	return nil
}

// Start starts this etcd coordinator.
func (coord *serviceImpl) Start() error {
	return nil
}

// Stop stops this etcd coordinator.
func (coord *serviceImpl) Stop() error {
	return nil
}
