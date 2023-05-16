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
)

type MessageBox struct {
	observers []coordinator.Observer
}

func NewMessageBox() *MessageBox {
	return &MessageBox{
		observers: []coordinator.Observer{},
	}
}

// AddObserver adds the specified observer.
func (mgr *MessageBox) AddObserver(newObserver coordinator.Observer) error {
	for _, observer := range mgr.observers {
		if observer == newObserver {
			return nil
		}
	}
	mgr.observers = append(mgr.observers, newObserver)
	return nil
}

// PostMessage posts the specified message to the coordinator.
func (mgr *MessageBox) PostMessage(msg coordinator.Message) error {
	return nil
}

// NofityMessage posts the specified message to the observers.
func (mgr *MessageBox) NofityMessage(msg coordinator.Message) error {
	for _, observer := range mgr.observers {
		observer.MessageReceived(msg)
	}
	return nil
}
