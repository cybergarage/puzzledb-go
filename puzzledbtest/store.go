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

package puzzledbtest

import (
	"strings"

	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/store"
	"github.com/cybergarage/puzzledb-go/puzzledb/util"
)

// Store represents a store.
type Store struct {
	store.Service
}

// NewStoreWith creates a new store with the given service.
func NewStoreWith(service store.Service) *Store {
	return &Store{
		Service: service,
	}
}

// Dump returns a string representation of the store.
func (store *Store) Dump() []string {
	dumpStore, err := util.NewStoreWith(store.Service)
	if err != nil {
		return []string{}
	}
	lines, err := dumpStore.Dump()
	if err != nil {
		return []string{}
	}
	return lines
}

// String returns a string representation of the store.
func (store *Store) String() string {
	return strings.Join(store.Dump(), "\n")
}
