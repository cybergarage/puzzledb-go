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

type watchersMap = map[string][]coordinator.Watcher

type NotifyManager struct {
	watchersMap watchersMap
}

func NewNotifyManager() *NotifyManager {
	return &NotifyManager{
		watchersMap: watchersMap{},
	}
}

// Watch registers a watcher for the specified key.
func (mgr *NotifyManager) Watch(key coordinator.Key, watcher coordinator.Watcher) error {
	keyStr, err := key.Encode()
	if err != nil {
		return err
	}

	watchers, ok := mgr.watchersMap[keyStr]
	if !ok {
		watchers = []coordinator.Watcher{}
	}
	mgr.watchersMap[keyStr] = append(watchers, watcher)

	return nil
}

// NotifyMessage notifies the specified event to the watchers.
func (mgr *NotifyManager) NotifyMessage(e coordinator.Message) error {
	keyStr, err := e.Object().Key().Encode()
	if err != nil {
		return err
	}
	for key, watcheres := range mgr.watchersMap {
		if !strings.HasPrefix(key, keyStr) {
			continue
		}
		for _, w := range watcheres {
			w.ProcessEvent(e)
		}
	}
	return nil
}
