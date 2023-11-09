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

package kvcache

import (
	"testing"

	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/coder/key/tuple"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/store/kv/memdb"
	plugin "github.com/cybergarage/puzzledb-go/puzzledb/plugins/store/kvcache"
	"github.com/cybergarage/puzzledb-go/puzzledbtest/plugins/store/kv"
)

//nolint:gosec,cyclop,gocognit,gocyclo,maintidx
func CacheStoreTest(t *testing.T, kvCacheStore plugin.Service) {
	t.Helper()

	kvStore := memdb.NewStore()
	kvStore.SetKeyCoder(tuple.NewCoder())
	if err := kvStore.Start(); err != nil {
		t.Error(err)
		return
	}

	kvCacheStore.SetStore(kvStore)
	kv.StoreTest(t, kvCacheStore)

	if err := kvStore.Stop(); err != nil {
		t.Error(err)
	}
}
