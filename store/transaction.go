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

type DocumentOperation interface {
	// InsertDocument puts a document object with the specified primary key.
	InsertDocument(key Key, obj Object) error
	// SelectDocuments gets document objects matching the specified key.
	SelectDocuments(key Key) ([]Object, error)
	// UpdateDocument updates a document object with the specified primary key.
	UpdateDocument(key Key, obj Object) error
	// RemoveDocument removes a document object with the specified primary key.
	RemoveDocument(key Key) error
}

type IndexOperation interface {
	// InsertIndex puts a secondary index with the primary key.
	InsertIndex(indexKey Key, key Key) error
	// SelectDocumentsByIndex gets document objects matching the specified index key.
	SelectDocumentsByIndex(indexKey Key) ([]Object, error)
	// UpdateDocument updates a document object with the specified primary key.
	UpdateDocument(key Key, obj Object) error
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
