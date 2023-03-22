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
	"strings"

	"github.com/cybergarage/puzzledb-go/puzzledb/coordinator"
)

type observerMap = map[string]coordinator.Observer

type Manager struct {
	observers observerMap
}

func NewManager() *Manager {
	return &Manager{
		observers: observerMap{},
	}
}

// AddObserver adds the observer to the coordinator.
func (mgr *Manager) AddObserver(key coordinator.Key, observer coordinator.Observer) error {
	keyStr, err := key.Encode()
	if err != nil {
		return err
	}
	mgr.observers[keyStr] = observer
	return nil
}

func (mgr *Manager) NofifyEvent(e coordinator.Event) error {
	eKeyStr, err := e.Object().Key().Encode()
	if err != nil {
		return err
	}
	for key, observer := range mgr.observers {
		if !strings.HasPrefix(key, eKeyStr) {
			continue
		}
		observer.ProcessEvent(e)
	}
	return nil
}
