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

package store

import (
	"fmt"
	"strings"
	"testing"

	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/store"
	"github.com/cybergarage/puzzledb-go/puzzledbtest"
)

func TestDocumentStore(t *testing.T) {
	docStore := store.NewStore()
	mgr := puzzledbtest.NewPluginManager()
	for _, kvStore := range mgr.EnabledKvStoreServices() {
		for _, keyCoder := range mgr.EnabledKeyCoderServices() {
			for _, docCoder := range mgr.EnabledDocumentCoderServices() {
				docStore.SetKvStore(kvStore)
				docStore.SetKeyCoder(keyCoder)
				docStore.SetDocumentCoder(docCoder)
				serviceNames := []string{keyCoder.ServiceName(), docCoder.ServiceName(), kvStore.ServiceName()}
				testName := fmt.Sprintf("%s(%s)", docStore.ServiceName(), strings.Join(serviceNames, ","))
				t.Run(testName, func(t *testing.T) {
					if err := kvStore.Start(); err != nil {
						t.Skip(err)
						return
					}
					defer func() {
						if err := kvStore.Stop(); err != nil {
							t.Error(err)
						}
					}()
					DocumentStoreCRUDTest(t, docStore)
				})
			}
		}
	}
}
