// Copyright (C) 2024 The PuzzleDB Authors.
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

package document

import (
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/store"
)

// Store represents a new document store utility instance.
type Store struct {
	store.DocumentStore
}

type DumpOptions struct {
	// DumpAll specifies whether to dump all the documents.
	DumpAll bool
}

// NewStoreWith returns a new document store utility instance with the specified store.
func NewStoreWith(store store.DocumentKvStore) *Store {
	return &Store{
		DocumentStore: store,
	}
}

// Dump returns a string array representation of the document store.
func (store *Store) Dump(opts DumpOptions) []string {
	line := []string{}
	return line
}
