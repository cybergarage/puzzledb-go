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
	"github.com/cybergarage/mimicdb/mimicdb/plugins"
)

// Store represents a store interface.
type Store interface {
	plugins.Service
	// Open opens the specified store.
	Open(name string) error
	// Transact opens a transaction.
	Transact(write bool) (Transaction, error)
	// Close closes this store.
	Close() error
}

// Key represents an object key.
type Key = string

// Transaction represents a transaction interface.
type Transaction interface {
	// Insert puts a key-value object.
	Insert(obj *Object) error
	// Select gets an key-value object of the specified key.
	Select(key Key) (*Object, error)
	// Commit commits this transaction.
	Commit() error
	// Cancel cancels this transaction.
	Cancel() error
}

// Object represents a key-value object.
type Object struct {
	Key   Key
	Value []byte
}
