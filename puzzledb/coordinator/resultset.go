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

// ResultSet represents a result set which includes range operation results.
type ResultSet interface {
	// Next moves the cursor forward next object from its current position.
	Next() bool
	// Object returns an object in the current position.
	Object() Object
	// Err returns the error, if any, that was encountered during iteration.
	Err() error
	// Close closes the result set and releases any resources.
	Close() error
}
