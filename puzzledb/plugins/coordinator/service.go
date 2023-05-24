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
	"context"
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
	*MessageQueue
	ctx       context.Context
	ctxCancel context.CancelFunc
}

// NewService returns a new coordinator service with the specified core coordinator service.
func NewServiceWith(service core.CoordinatorService) Service {
	ctx, cancel := context.WithCancel(context.Background())
	return &serviceImpl{
		CoordinatorService: service,
		Process:            coordinator.NewProcess(),
		observers:          make([]coordinator.Observer, 0),
		MessageQueue:       NewMessageQueue(),
		ctx:                ctx,
		ctxCancel:          cancel,
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

// nofityMessage posts the specified message to the observers.
func (coord *serviceImpl) nofityMessage(msg coordinator.Message) {
	for _, observer := range coord.observers {
		observer.MessageReceived(msg)
	}
}

func (coord *serviceImpl) notifyUpdateMessages(txn coordinator.Transaction) error {
	key := NewScanMessageKey()
	rs, err := txn.GetRange(
		key,
		coordinator.NewOrderOptionWith(coordinator.OrderDesc))
	if err != nil {
		return err
	}

	localClock := coord.Clock()

	msgs := []coordinator.Message{}
	for rs.Next() {
		msgObj := NewMessageObject()
		obj := rs.Object()
		err = obj.Unmarshal(msgObj)
		if err != nil {
			return err
		}

		// Skip the message if the message clock is older than the local clock
		if 0 < coordinator.CompareClocks(localClock, msgObj.Clock) {
			break
		}

		// Skip the self messages
		if msgObj.Host == coord.Host() {
			continue
		}

		msg := NewMessageWith(obj.Key(), msgObj)
		msgs = append([]coordinator.Message{msg}, msgs...)

		coord.SetReceivedClock(msgObj.Clock)
	}

	for _, msg := range msgs {
		log.Infof("Received message: %s %s (%d)", msg.From().Host(), msg.Type().String(), msg.Clock())
		coord.nofityMessage(msg)
	}

	return nil
}

// PostMessage posts the specified message to the coordinator.
func (coord *serviceImpl) PostMessage(msg coordinator.Message) error {
	coord.Lock()
	defer coord.Unlock()

	coord.EnqueueMessage(msg)

	return nil
}

// postMessage posts the specified message to the coordinator.
func (coord *serviceImpl) postMessage(txn coordinator.Transaction, msg coordinator.Message) error {
	localClock := coord.IncrementClock()

	key := NewMessageKeyWith(msg, localClock)
	obj, err := NewMessageObjectWith(msg, coord, localClock)
	if err != nil {
		return errors.Join(err, txn.Cancel())
	}

	objBytes, err := cbor.Marshal(obj)
	if err != nil {
		return err
	}

	err = txn.Set(coordinator.NewObjectWith(key, objBytes))
	if err != nil {
		return err
	}

	return nil
}

// Start starts this etcd coordinator.
func (coord *serviceImpl) Start() error {
	if err := coord.CoordinatorService.Start(); err != nil {
		return err
	}
	go func() {
		logError := func(err error) {
			log.Errorf("coordinator worker: %s", err)
		}

		for {
			jitter := time.Duration(rand.Intn(int(DefaultStoreScanInterval/time.Millisecond/2))) * time.Millisecond //nolint:gosec
			select {
			case <-time.After(DefaultStoreScanInterval + jitter):
				var err error
				coord.Lock()

				txn, err := coord.Transact()
				if err != nil {
					logError(err)
					coord.Unlock()
					continue
				}

				// Receive update messages and update local clock

				err = coord.notifyUpdateMessages(txn)
				if err != nil {
					if err != nil {
						logError(errors.Join(err, txn.Cancel()))
					}
					coord.Unlock()
					continue
				}

				// Post message if there is no message in the queue

				msg, err := coord.PopMessage()
				for msg != nil {
					err = coord.postMessage(txn, msg)
					if err != nil {
						coord.PushMessage(msg)
						logError(err)
						break
					}
					msg, err = coord.PopMessage()
				}

				if err != nil && !errors.Is(err, coordinator.ErrNoMessage) {
					logError(err)
				}

				err = txn.Commit()
				if err != nil {
					logError(err)
				}

				coord.Unlock()
			case <-coord.ctx.Done():
				return
			}
		}
	}()
	return nil
}

// Stop stops this etcd coordinator.
func (coord *serviceImpl) Stop() error {
	coord.ctxCancel()
	<-coord.ctx.Done()

	if err := coord.CoordinatorService.Stop(); err != nil {
		return err
	}

	return nil
}
