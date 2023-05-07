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

package cache

import (
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/store/kv"
)

// Store represents a cache store service instance.
type Store struct {
	kv.Service
}

// NewStore returns a new FoundationDB store instance.
func NewStore() kv.Service {
	return NewStoreWith(nil)
}

// NewStoreWith returns a new FoundationDB store instance with the specified key coder.
func NewStoreWith(service kv.Service) kv.Service {
	return &Store{
		Service: service,
	}
}