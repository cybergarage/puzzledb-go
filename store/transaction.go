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

type DatabaseOperation interface {
	// CreateDatabase creates a database object with the specified name.
	CreateDatabase(name string) error
	// FindDocuments returns a result set matching the specified key.
	GetDatabase(name string) (Database, error)
}

type DocumentOperation interface {
	// InsertDocument puts a document object with the specified primary key.
	InsertDocument(docKey Key, obj Object) error
	// FindDocuments returns a result set matching the specified key.
	FindDocuments(docKey Key) (ResultSet, error)
	// UpdateDocument updates a document object with the specified primary key.
	UpdateDocument(docKey Key, obj Object) error
	// RemoveDocument removes a document object with the specified primary key.
	RemoveDocument(docKey Key) error
}

type IndexOperation interface {
	// InsertIndex puts a secondary index with the primary key.
	InsertIndex(idxKey Key, key Key) error
	// RemoveIndex removes the specified secondary index.
	RemoveIndex(idxKey Key) error
	// FindDocumentsByIndex returns a result set matching the specified index key.
	FindDocumentsByIndex(indexKey Key) (ResultSet, error)
}

// Transaction represents a transaction interface.
type Transaction interface {
	DocumentOperation
	IndexOperation
	// Commit commits this transaction.
	Commit() error
	// Cancel cancels this transaction.
	Cancel() error
}
