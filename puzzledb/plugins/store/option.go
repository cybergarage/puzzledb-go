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
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
	"github.com/cybergarage/puzzledb-go/puzzledb/store/kv"
)

// NewKvOptionsWith returns a new kv.options with the specified store options.
func NewKvOptionsWith(opts ...store.Option) []kv.Option {
	kvOpts := []kv.Option{}
	for _, opt := range opts {
		switch v := opt.(type) {
		case store.Offset:
			kvOpts = append(kvOpts, kv.Offset(v))
		case store.Limit:
			kvOpts = append(kvOpts, kv.Limit(v))
		case store.Order:
			kvOpts = append(kvOpts, kv.Order(v))
		}
	}
	return kvOpts
}
