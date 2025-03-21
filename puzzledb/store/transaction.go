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
	"time"

	"github.com/cybergarage/puzzledb-go/puzzledb/context"
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
)

type Collection = document.Collection
type Schema = document.Schema

// DatabaseOperation represents database operations.
type DatabaseOperation interface {
	// ListCollections returns the all collection in the database.
	ListCollections(ctx context.Context) ([]Collection, error)
	// CreateCollection creates a new collection into the database.
	CreateCollection(ctx context.Context, col Collection) error
	// UpdateCollection updates the specified collection in the database.
	UpdateCollection(ctx context.Context, col Collection) error
	// LookupCollection returns the specified collection in the database.
	LookupCollection(ctx context.Context, name string) (Collection, error)
	// RemoveCollection removes the specified collection in the database.
	RemoveCollection(ctx context.Context, name string) error
	// TruncateCollections removes all collections in the database.
	TruncateCollections(ctx context.Context) error
}

// CollectionOperation represents collection operations.
type CollectionOperation interface {
	// InsertObject puts a document object with the specified primary key.
	InsertObject(ctx context.Context, docKey Key, obj Object) error
	// FindObjects returns a result set matching the specified key.
	FindObjects(ctx context.Context, docKey Key, opts ...Option) (ResultSet, error)
	// UpdateObject updates a document object with the specified primary key.
	UpdateObject(ctx context.Context, docKey Key, obj Object) error
	// RemoveObject removes a document object with the specified primary key.
	RemoveObject(ctx context.Context, docKey Key) error
	// RemoveObjects removes document objects with the specified primary key.
	RemoveObjects(ctx context.Context, docKey Key) error
	// TruncateObjects removes all document objects.
	TruncateObjects(ctx context.Context) error
}

// IndexOperation represents a secondary index operation.
type IndexOperation interface {
	// InsertIndex puts a secondary index with the primary key.
	InsertIndex(ctx context.Context, idxKey Key) error
	// RemoveIndex removes the specified secondary index.
	RemoveIndex(ctx context.Context, idxKey Key) error
	// FindObjectsByIndex returns a result set matching the specified index key.
	FindObjectsByIndex(ctx context.Context, idxKey Key, opts ...Option) (ResultSet, error)
	// TruncateIndexes removes all secondary indexes.
	TruncateIndexes(ctx context.Context) error
}

// StoreOperation represents a transaction operation.
type StoreOperation interface {
	DatabaseOperation
	CollectionOperation
	IndexOperation
}

// TransactionOption represents a transaction option.
type TransactionOption interface {
	// SetAutoCommit sets the auto commit flag.
	SetAutoCommit(bool)
	// IsAutoCommit returns true whether the auto commit flag is set.
	IsAutoCommit() bool
	// SetTimeout sets the timeout of this transaction.
	SetTimeout(t time.Duration) error
}

// TransactionOperation represents a transaction operation.
type TransactionOperation interface {
	TransactionOption
	// Commit commits this transaction.
	Commit(ctx context.Context) error
	// Cancel cancels this transaction.
	Cancel(ctx context.Context) error
}

// Transaction represents a transaction interface.
type Transaction interface {
	TransactionOperation
	StoreOperation
	// Database returns the transaction database.
	Database() Database
}
