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

import (
	"github.com/google/uuid"
)

// Process represents a coordinator process.
type Process interface {
	// ID returns a UUID of the coordinator process.
	ID() uuid.UUID
	// SetHost sets a host name to the coordinator process.
	SetHost(host string)
	// Host returns a host name of the coordinator process.
	Host() string
	// SetClock sets a logical clock to the coordinator process.
	SetClock(clock Clock)
	// Clock returns a logical clock of the coordinator process.
	Clock() Clock
}

// Store represents a coordination store inteface.
type Store interface {
	// Transact begin a new transaction.
	Transact() (Transaction, error)
}

// Coordinator represents a coordination service.
type Coordinator interface {
	Store
	Process
	// PostMessage posts the specified message to the coordinator.
	PostMessage(msg Message) error
	// AddObserver adds the specified observer to the coordinator.
	AddObserver(observer Observer) error
}
