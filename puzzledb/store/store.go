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
	"github.com/cybergarage/puzzledb-go/puzzledb/context"
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
)

// Store represents a store interface.
type Store interface {
	// SetDocumentCoder sets the document coder.
	SetDocumentCoder(coder document.Coder)
	// SetKeyCoder sets the key coder.
	SetKeyCoder(coder document.KeyCoder)
	// CreateDatabase creates a new database.
	CreateDatabase(ctx context.Context, name string) error
	// GetDatabase retruns the specified database.
	GetDatabase(ctx context.Context, name string) (Database, error)
	// RemoveDatabase removes the specified database.
	RemoveDatabase(ctx context.Context, name string) error
	// ListDatabases returns the all databases.
	ListDatabases(ctx context.Context) ([]Database, error)
}
