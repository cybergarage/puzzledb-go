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

package memdb

// Database represents a database.
type Database struct {
	ID string
}

// NewDatabaseWithID returns a new database with the specified ID.
func NewDatabaseWithID(id string) *Database {
	return &Database{
		ID: id,
	}
}

// Name returns the unique name.
func (db *Database) Name() string {
	return db.ID
}
