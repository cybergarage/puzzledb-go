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
	"fmt"
	"testing"

	"github.com/cybergarage/go-logger/log"
	"github.com/cybergarage/puzzledb-go/puzzledbtest"
)

func TestCoordinator(t *testing.T) {
	log.SetSharedLogger(log.NewStdoutLogger(log.LevelInfo))

	mgr := puzzledbtest.NewPluginManager()
	for _, keyCoder := range mgr.EnabledKeyCoderServices() {
		for _, coord := range mgr.EnabledCoordinatorServices() {
			coord.SetKeyCoder(keyCoder)
			testName := fmt.Sprintf("%s(%s)", coord.ServiceName(), keyCoder.ServiceName())
			t.Run(testName, func(t *testing.T) {
				if err := coord.Start(); err != nil {
					t.Skip(err)
					return
				}
				defer func() {
					if err := coord.Stop(); err != nil {
						t.Error(err)
					}
				}()
				t.Run("message", func(t *testing.T) {
					CoordinatorMessageTest(t, coord)
				})
				t.Run("process", func(t *testing.T) {
					// CoordinatorProcessTest(t, coord)
				})
			})
		}
	}
}
