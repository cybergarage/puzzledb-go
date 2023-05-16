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

package core

import (
	"github.com/cybergarage/puzzledb-go/puzzledb/coordinator"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins"
)

type BaseCoordinator struct {
	observers []coordinator.Observer
	coordinator.Process
	plugins.Config
}

// NewBaseCoordinator returns a new base coordinator instance.
func NewBaseCoordinator() *BaseCoordinator {
	return &BaseCoordinator{
		Process:   coordinator.NewProcess(),
		Config:    plugins.NewConfig(),
		observers: []coordinator.Observer{},
	}
}

// AddObserver adds the specified observer.
func (coord *BaseCoordinator) AddObserver(newObserver coordinator.Observer) error {
	for _, observer := range coord.observers {
		if observer == newObserver {
			return nil
		}
	}
	coord.observers = append(coord.observers, newObserver)
	return nil
}

// PostMessage posts the specified message to the coordinator.
func (coord *BaseCoordinator) PostMessage(msg coordinator.Message) error {
	return nil
}

// NofityMessage posts the specified message to the observers.
func (coord *BaseCoordinator) NofityMessage(msg coordinator.Message) error {
	for _, observer := range coord.observers {
		observer.MessageReceived(msg)
	}
	return nil
}
