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

// Clock represents a logical clock.
type Clock uint64

// Process represents a coordinator process.
type Process interface {
	// SetClock sets a coordinator logical clock.
	SetClock(clock Clock)
	// Clock returns a coordinator logical clock.
	Clock() Clock
}

// Store represents a coordination store inteface.
type Store interface {
	// Transact begin a new transaction.
	Transact() (Transaction, error)
	// Watch adds a watcher to the coordinator.
	Watch(key Key, w Watcher) error
}

// Coordinator represents a coordination service.
type Coordinator interface {
	Store
	Process
}
