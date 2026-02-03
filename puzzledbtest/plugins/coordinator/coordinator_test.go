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
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/coordinator"
	"github.com/cybergarage/puzzledb-go/puzzledbtest"
)

func TestCoordinators(t *testing.T) {
	log.SetDefault(log.NewStdoutLogger(log.LevelInfo))

	mgr := puzzledbtest.NewPluginManager()

	mgr01 := puzzledbtest.NewPluginManager()
	mgr02 := puzzledbtest.NewPluginManager()

	coords01 := mgr01.EnabledCoordinatorServices()
	coords02 := mgr02.EnabledCoordinatorServices()

	for _, keyCoder := range mgr.EnabledKeyCoderServices() {
		for i, coords01 := range coords01 {
			coords := []coordinator.Service{coords01, coords02[i]}
			for j, coord := range coords {
				coord.SetHost(fmt.Sprintf("coodinator%02d", j))
				coord.SetKeyCoder(keyCoder)
				if err := coord.Start(); err != nil {
					t.Error(err)
					return
				}
			}

			testSuffix := fmt.Sprintf("(%s,%s)", coords[0].ServiceName(), keyCoder.ServiceName())

			tests := []struct {
				name string
				fn   func(t *testing.T, coords []coordinator.Service)
			}{
				{"messaging", CoordinatorsTest},
				{"clustring", CoordinatorClusterTest},
			}
			for _, test := range tests {
				t.Run(test.name+testSuffix, func(t *testing.T) {
					test.fn(t, coords)
				})
			}

			for _, coord := range coords {
				if err := coord.Stop(); err != nil {
					t.Error(err)
					return
				}
			}
		}
	}
}
