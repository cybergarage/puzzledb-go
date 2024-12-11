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

	"github.com/cybergarage/puzzledb-go/puzzledb/document"
	"github.com/cybergarage/puzzledb-go/puzzledb/document/kv"
)

func TestConfig(t *testing.T) {
	prefixes := []kv.KeyHeader{
		kv.DatabaseKeyHeader,
		kv.CollectionKeyHeader,
	}

	conf := NewCacheConfig()

	for _, prefix := range prefixes {
		conf.RegisterCacheKeyHeader(prefix)
	}

	testKeys := []kv.Key{
		kv.NewKeyWith(kv.DatabaseKeyHeader, document.NewKey()),
		kv.NewKeyWith(kv.DatabaseKeyHeader, document.NewKey()),
	}

	for _, testKey := range testKeys {
		if !conf.IsRegisteredCacheKey(testKey) {
			t.Errorf("The key (%s) must be cached", testKey.String())
		}
	}

	testKeys = []kv.Key{
		kv.NewKeyWith(kv.DocumentKeyHeader, document.NewKey()),
		kv.NewKeyWith(kv.SecondaryIndexHeader, document.NewKey()),
	}

	for _, testKey := range testKeys {
		if conf.IsRegisteredCacheKey(testKey) {
			t.Errorf("The key (%s) must not be cached", testKey.String())
		}
	}
}
