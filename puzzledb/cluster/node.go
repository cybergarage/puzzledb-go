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

package cluster

import (
	"time"

	"github.com/google/uuid"
)

// Node represents a cluster node.
type Node interface {
	// SetID sets a UUID to the cluster node.
	SetID(uuid uuid.UUID)
	// ID returns a UUID of the cluster node.
	ID() uuid.UUID
	// SetHost sets a host name to the cluster node.
	SetHost(host string)
	// Host returns a host name of the cluster node.
	Host() string
	// SetClock sets a logical clock to the cluster node.
	SetClock(clock Clock)
	// Clock returns a logical clock of the cluster node.
	Clock() Clock
	// SetReceivedClock sets a received logical clock to the cluster node.
	SetReceivedClock(clock Clock) Clock
	// IncrementClock increments a logical clock of the cluster node.
	IncrementClock() Clock
	// SetTimestamp sets a phisical timestamp to the cluster node.
	SetTimestamp(ts time.Time)
	// Timestamp returns a phisical timestamp of the cluster node.
	Timestamp() time.Time
	// SetStatus sets a status to the cluster node.
	SetStatus(state NodeStatus)
	// Status returns a status of the cluster node.
	Status() NodeStatus
	// Lock locks the cluster node.
	Lock()
	// Unlock unlocks the cluster node.
	Unlock()
}
