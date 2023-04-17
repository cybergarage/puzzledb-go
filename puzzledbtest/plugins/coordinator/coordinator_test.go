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
	"strings"
	"testing"

	"github.com/cybergarage/puzzledb-go/puzzledbtest"
)

func TestCoordinator(t *testing.T) {
	mgr := puzzledbtest.NewPluginManager()
	for _, coodinator := range mgr.EnabledCoordinatorServices() {
		for _, keyCoder := range mgr.EnabledKeyCoderServices() {
			for _, docCoder := range mgr.EnabledDocumentCoderServices() {
				if err := coodinator.Start(); err != nil {
					t.Skip(err)
					return
				}
				defer func() {
					if err := coodinator.Stop(); err != nil {
						t.Error(err)
					}
				}()
				coodinator.SetKeyCoder(keyCoder)
				coodinator.SetDocumentCoder(docCoder)
				serviceNames := []string{keyCoder.ServiceName(), docCoder.ServiceName()}
				testName := fmt.Sprintf("%s(%s)", coodinator.ServiceName(), strings.Join(serviceNames, ","))
				t.Run(testName, func(t *testing.T) {
					CoordinatorTest(t, coodinator)
				})
			}
		}
	}
}
