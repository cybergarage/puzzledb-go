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
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
)

type Schema = document.Schema

type DatabaseOperation interface {
	// CreateSchema creates a new schema.
	CreateSchema(schema Schema) error
	// GetSchema returns the specified schema.
	GetSchema(name string) (Schema, error)
	// RemoveSchema removes the specified schema.
	RemoveSchema(name string) error
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
	// RemoveDocuments removes document objects with the specified primary key.
	RemoveDocuments(docKey Key) error
}

type IndexOperation interface {
	// InsertIndex puts a secondary index with the primary key.
	InsertIndex(idxKey Key, key Key) error
	// RemoveIndex removes the specified secondary index.
	RemoveIndex(idxKey Key) error
	// FindDocumentsByIndex returns a result set matching the specified index key.
	FindDocumentsByIndex(indexKey Key) (ResultSet, error)
}

type TransactionOperation interface {
	DatabaseOperation
	DocumentOperation
	IndexOperation
}

// Transaction represents a transaction interface.
type Transaction interface {
	TransactionOperation
	// Database returns the transaction database.
	Database() Database
	// Commit commits this transaction.
	Commit() error
	// Cancel cancels this transaction.
	Cancel() error
}
