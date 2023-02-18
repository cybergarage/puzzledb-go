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

// Transaction represents a transaction interface.
type Transaction interface {
	// InsertDocument puts a document object with the primary key.
	InsertDocument(key Key, obj Object) error
	// SelectDocument gets a document object with the specified key.
	SelectDocument(key Key) (Object, error)
	// InsertIndex puts a secondary index with the primary key.
	InsertIndex(key Key, primeryKey Key) error
	// Commit commits this transaction.
	Commit() error
	// Cancel cancels this transaction.
	Cancel() error
}
