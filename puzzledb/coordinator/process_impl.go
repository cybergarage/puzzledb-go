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
	"math"
	"sync"

	"github.com/google/uuid"
)

type processImpl struct {
	sync.Mutex
	clock Clock
	uuid  uuid.UUID
}

func NewProcess() Process {
	return &processImpl{
		Mutex: sync.Mutex{},
		clock: 0,
		uuid:  uuid.New(),
	}
}

// ID returns a UUID of the coordinator process.
func (process *processImpl) ID() uuid.UUID {
	return process.uuid
}

// SetClock sets a logical clock to the coordinator process.
func (process *processImpl) SetClock(newClock Clock) {
	if newClock < process.clock {
		newClock = process.clock
	}
	if (math.MaxUint64 - 1) <= newClock {
		newClock = 0
	}
	process.clock = newClock + 1
}

// Clock returns a logical clock of the coordinator process.
func (process *processImpl) Clock() Clock {
	return process.clock
}
