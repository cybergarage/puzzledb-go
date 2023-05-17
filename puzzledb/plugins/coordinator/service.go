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
	"errors"
	"math/rand"
	"time"

	"github.com/cybergarage/go-logger/log"
	"github.com/cybergarage/puzzledb-go/puzzledb/coordinator"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/coordinator/core"
)

const (
	DefaultStoreScanInterval = time.Second
)

type Service interface {
	coordinator.Coordinator
	plugins.Service
}

type serviceImpl struct {
	core.CoordinatorService
	observers []coordinator.Observer
	coordinator.Process
	*time.Ticker
}

func NewServiceWith(service core.CoordinatorService) Service {
	return &serviceImpl{
		CoordinatorService: service,
		Process:            coordinator.NewProcess(),
		observers:          make([]coordinator.Observer, 0),
		Ticker:             time.NewTicker(DefaultStoreScanInterval),
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

// SetObject sets the object for the specified key.
func (coord *serviceImpl) SetObject(obj coordinator.Object) error {
	txn, err := coord.Transact()
	if err != nil {
		return err
	}
	err = txn.Set(obj)
	if err != nil {
		return errors.Join(err, txn.Cancel())
	}
	return txn.Commit()
}

// GetObject gets the object for the specified key.
func (coord *serviceImpl) GetObject(key coordinator.Key) (coordinator.Object, error) {
	txn, err := coord.Transact()
	if err != nil {
		return nil, err
	}
	obj, err := txn.Get(key)
	if err != nil {
		return nil, errors.Join(err, txn.Cancel())
	}
	err = txn.Commit()
	return obj, err
}

// GetRangeObjects gets the result set for the specified key.
func (coord *serviceImpl) GetRangeObjects(key coordinator.Key) (coordinator.ResultSet, error) {
	txn, err := coord.Transact()
	if err != nil {
		return nil, err
	}
	rs, err := txn.GetRange(key)
	if err != nil {
		return nil, errors.Join(err, txn.Cancel())
	}
	err = txn.Commit()
	return rs, err
}

// PostMessage posts the specified message to the coordinator.
func (coord *serviceImpl) PostMessage(msg coordinator.Message) error {
	return nil
}

func (coord *serviceImpl) GetUpdateMessages() ([]coordinator.Message, error) {
	msgs := []coordinator.Message{}
	return msgs, nil
}

// NofityMessage posts the specified message to the observers.
func (coord *serviceImpl) NofityMessage(msg coordinator.Message) {
	for _, observer := range coord.observers {
		observer.MessageReceived(msg)
	}
}

// Start starts this etcd coordinator.
func (coord *serviceImpl) Start() error {
	go func() {
		for range coord.Ticker.C {
			msgs, err := coord.GetUpdateMessages()
			if err != nil {
				log.Errorf("Failed to get the update coordinator messages : %s", err)
				continue
			}
			for _, msg := range msgs {
				coord.NofityMessage(msg)
			}
			// Reset the timer with a random jitter.
			coord.Ticker.Reset(DefaultStoreScanInterval + time.Duration(rand.Intn(100))*time.Millisecond) //nolint:gosec
		}
	}()
	return nil
}

// Stop stops this etcd coordinator.
func (coord *serviceImpl) Stop() error {
	coord.Ticker.Stop()
	return nil
}
