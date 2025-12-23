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

package coordinator

// Transaction represents a transaction interface.
type Transaction interface {
	// Set sets the object for the specified key.
	Set(obj Object) error
	// Get gets the object for the specified key.
	Get(key Key) (Object, error)
	// Scan returns the result set for the specified key.
	Scan(key Key, opts ...Option) (ResultSet, error)
	// Remove removes the object for the specified key.
	Remove(key Key) error
	// Commit commits this transaction.
	Commit() error
	// Cancel cancels this transaction.
	Cancel() error
	// Truncate removes all objects.
	Truncate() error
}
