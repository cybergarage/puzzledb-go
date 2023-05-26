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
	"os"
	"sync"
	"time"

	"github.com/google/uuid"
)

type processImpl struct {
	sync.Mutex
	clock  Clock
	uuid   uuid.UUID
	host   string
	ts     time.Time
	status ProcessStatus
}

func NewProcess() Process {
	p := &processImpl{
		host:   "",
		Mutex:  sync.Mutex{},
		clock:  NewClock(),
		uuid:   uuid.New(),
		ts:     time.Now(),
		status: ProcessIdle,
	}
	host, err := os.Hostname()
	if err == nil {
		p.host = host
	}
	return p
}

// SetID sets a UUID to the coordinator process.
func (process *processImpl) SetID(uuid uuid.UUID) {
	process.uuid = uuid
}

// ID returns a UUID of the coordinator process.
func (process *processImpl) ID() uuid.UUID {
	return process.uuid
}

// SetHost sets a host name to the coordinator process.
func (process *processImpl) SetHost(host string) {
	process.host = host
}

// Host returns a host name of the coordinator process.
func (process *processImpl) Host() string {
	return process.host
}

// SetClock sets a logical clock to the coordinator process.
func (process *processImpl) SetClock(clock Clock) {
	process.clock = clock
	process.ts = time.Now()
}

// SetReceivedClock sets a received logical clock to the coordinator process.
func (process *processImpl) SetReceivedClock(clock Clock) Clock {
	process.clock = MaxClock(clock, process.clock)
	process.IncrementClock()
	process.ts = time.Now()
	return process.clock
}

// IncrementClock increments a logical clock of the coordinator process.
func (process *processImpl) IncrementClock() Clock {
	if (ClockMax - ClockDiffrent) <= process.clock {
		process.clock = 0
	} else {
		process.clock += ClockDiffrent
	}
	process.ts = time.Now()
	return process.clock
}

// Clock returns a logical clock of the coordinator process.
func (process *processImpl) Clock() Clock {
	return process.clock
}

// SetTimestamp sets a phisical timestamp to the coordinator process.
func (process *processImpl) SetTimestamp(ts time.Time) {
	process.ts = ts
}

// Timestamp returns a phisical timestamp of the coordinator process.
func (process *processImpl) Timestamp() time.Time {
	return process.ts
}

// SetStatus sets a status to the coordinator process.
func (process *processImpl) SetStatus(status ProcessStatus) {
	process.status = status
}

// Status returns a status of the coordinator process.
func (process *processImpl) Status() ProcessStatus {
	return process.status
}
