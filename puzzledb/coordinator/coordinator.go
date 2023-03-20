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

// Value represents a value for a key-value object.
type Value any

// Coordinator represents a coordination service.
type Coordinator interface {
	// Set sets the value for the specified key.
	Set(key Key, value Value) error
	// Get gets the value for the specified key.
	Get(key Key) (Value, error)
}
