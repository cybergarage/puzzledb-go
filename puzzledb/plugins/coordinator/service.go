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

	"github.com/cybergarage/go-cbor/cbor"
	"github.com/cybergarage/go-logger/log"
	"github.com/cybergarage/puzzledb-go/puzzledb/coordinator"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/coordinator/core"
)

const (
	DefaultStoreScanInterval = time.Second
)

// Service is an interface for the coordinator service.
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

// NewService returns a new coordinator service with the specified core coordinator service.
func NewServiceWith(service core.CoordinatorService) Service {
	return &serviceImpl{
		CoordinatorService: service,
		Process:            coordinator.NewProcess(),
		observers:          make([]coordinator.Observer, 0),
		Ticker:             time.NewTicker(DefaultStoreScanInterval),
	}
}

// SetProcess sets the coordinator process.
func (coord *serviceImpl) SetProcess(process coordinator.Process) {
	coord.Process = process
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

// SetStateObject sets the state object for the specified key.
func (coord *serviceImpl) SetStateObject(t coordinator.StateType, obj coordinator.Object) error {
	var err error
	var key coordinator.Key
	var objBytes []byte
	switch v := obj.(type) {
	case coordinator.Process:
		key = coordinator.NewStateKeyWith(t, v.Host())
		p := &ProcessObject{
			ID:    v.ID(),
			Host:  v.Host(),
			Clock: uint64(v.Clock()),
		}
		objBytes, err = cbor.Marshal(p)
		if err != nil {
			return err
		}
	default:
		return coordinator.NewErrObjectNotSupported(obj)
	}

	txn, err := coord.Transact()
	if err != nil {
		return err
	}
	err = txn.Set(coordinator.NewObjectWith(key, objBytes))
	if err != nil {
		return errors.Join(err, txn.Cancel())
	}
	return txn.Commit()
}

// GetObject gets the object for the specified key and state type.
func (coord *serviceImpl) GetStateObject(t coordinator.StateType, key coordinator.Key) (coordinator.Object, error) {
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

// GetRangeObjects gets the result set for the specified key and state type.
func (coord *serviceImpl) GetStateObjects(t coordinator.StateType) (coordinator.ResultSet, error) {
	txn, err := coord.Transact()
	if err != nil {
		return nil, err
	}
	rs, err := txn.GetRange(coordinator.NewScanStateKeyWith(t))
	if err != nil {
		return nil, errors.Join(err, txn.Cancel())
	}
	err = txn.Commit()
	return rs, err
}

// ScanLatestMessageClock returns the latest message clock.
func (coord *serviceImpl) ScanLatestMessageClock(txn coordinator.Transaction) (coordinator.Clock, error) {
	key := NewScanMessageKey()
	rs, err := txn.GetRange(
		key,
		coordinator.NewLimitOption(1),
		coordinator.NewOrderOptionWith(coordinator.OrderDesc))
	if err != nil {
		return 0, err
	}
	if !rs.Next() {
		return 0, nil
	}
	var msgObj MessageObject
	err = rs.Object().Unmarshal(&msgObj)
	if err != nil {
		return 0, err
	}
	return msgObj.Clock, nil
}

// GetLatestMessageClock returns the latest message clock.
func (coord *serviceImpl) GetLatestMessageClock() (coordinator.Clock, error) {
	txn, err := coord.Transact()
	if err != nil {
		return 0, err
	}
	clock, err := coord.ScanLatestMessageClock(txn)
	if err != nil {
		return 0, errors.Join(err, txn.Cancel())
	}
	return clock, txn.Commit()
}

// PostMessage posts the specified message to the coordinator.
func (coord *serviceImpl) PostMessage(msg coordinator.Message) error {
	coord.Lock()
	defer coord.Unlock()

	txn, err := coord.Transact()
	if err != nil {
		return err
	}

	// Receive the latest message clock

	localClock := coord.Clock()
	coordClock, err := coord.ScanLatestMessageClock(txn)
	if err != nil {
		return errors.Join(err, txn.Cancel())
	}
	coord.SetClock(coordinator.MaxClock(coordClock, localClock))
	coord.IncrementClock()

	// Post a new message

	localClock = coord.IncrementClock()

	key := NewMessageKeyWith(msg, localClock)
	obj, err := NewMessageObjectWith(msg, coord, localClock)
	if err != nil {
		return errors.Join(err, txn.Cancel())
	}

	objBytes, err := cbor.Marshal(obj)
	if err != nil {
		return errors.Join(err, txn.Cancel())
	}

	err = txn.Set(coordinator.NewObjectWith(key, objBytes))
	if err != nil {
		return errors.Join(err, txn.Cancel())
	}

	return txn.Commit()
}

func (coord *serviceImpl) GetUpdateMessages() ([]coordinator.Message, error) {
	msgs := []coordinator.Message{}
	txn, err := coord.Transact()
	if err != nil {
		return nil, err
	}

	localClock := coord.Clock()
	key := NewScanMessageKeyWith(localClock)
	rs, err := txn.GetRange(key)
	if err != nil {
		return msgs, errors.Join(err, txn.Cancel())
	}

	for rs.Next() {
		var msgObj MessageObject
		obj := rs.Object()
		err = obj.Unmarshal(&msgObj)
		if err != nil {
			return msgs, errors.Join(err, txn.Cancel())
		}
		msg := NewMessageWith(obj.Key(), &msgObj)
		msgs = append(msgs, msg)
	}

	return msgs, txn.Commit()
}

// NofityMessage posts the specified message to the observers.
func (coord *serviceImpl) NofityMessage(msg coordinator.Message) {
	for _, observer := range coord.observers {
		observer.MessageReceived(msg)
	}
}

// Start starts this etcd coordinator.
func (coord *serviceImpl) Start() error {
	if err := coord.CoordinatorService.Start(); err != nil {
		return err
	}
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
	if err := coord.CoordinatorService.Stop(); err != nil {
		return err
	}
	coord.Ticker.Stop()
	return nil
}
