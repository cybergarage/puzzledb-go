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

package memdb

import (
	"testing"

	"github.com/cybergarage/mimicdb/mimicdb/plugins/store"
)

func TestStores(t *testing.T) {
	stores := []store.Store{
		NewStore(),
	}

	for _, store := range stores {
		testStore(t, store)
	}
}

func testStore(t *testing.T, store store.Store) {
	if err := store.Start(); err != nil {
		t.Error(err)
	}
	if err := store.Open("testdb"); err != nil {
		t.Error(err)
	}
	if err := store.Close(); err != nil {
		t.Error(err)
	}
	if err := store.Stop(); err != nil {
		t.Error(err)
	}
}
