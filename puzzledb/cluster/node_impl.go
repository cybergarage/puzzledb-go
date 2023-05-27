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
	"os"
	"sync"
	"time"

	"github.com/google/uuid"
)

type nodeImpl struct {
	sync.Mutex
	clock   Clock
	cluster string
	uuid    uuid.UUID
	host    string
	ts      time.Time
	status  NodeStatus
}

func NewNode() Node {
	p := &nodeImpl{
		host:    "",
		Mutex:   sync.Mutex{},
		clock:   NewClock(),
		cluster: DefaultClusterName,
		uuid:    uuid.New(),
		ts:      time.Now(),
		status:  NodeIdle,
	}
	host, err := os.Hostname()
	if err == nil {
		p.host = host
	}
	return p
}

// SetCluster sets a cluster name to the cluster node.
func (node *nodeImpl) SetCluster(cluster string) {
	node.cluster = cluster
}

// Cluster returns a cluster name of the cluster node.
func (node *nodeImpl) Cluster() string {
	return node.cluster
}

// SetID sets a UUID to the cluster node.
func (node *nodeImpl) SetID(uuid uuid.UUID) {
	node.uuid = uuid
}

// ID returns a UUID of the cluster node.
func (node *nodeImpl) ID() uuid.UUID {
	return node.uuid
}

// SetHost sets a host name to the cluster node.
func (node *nodeImpl) SetHost(host string) {
	node.host = host
}

// Host returns a host name of the cluster node.
func (node *nodeImpl) Host() string {
	return node.host
}

// SetClock sets a logical clock to the cluster node.
func (node *nodeImpl) SetClock(clock Clock) {
	node.clock = clock
	node.ts = time.Now()
}

// SetReceivedClock sets a received logical clock to the cluster node.
func (node *nodeImpl) SetReceivedClock(clock Clock) Clock {
	node.clock = MaxClock(clock, node.clock)
	node.IncrementClock()
	node.ts = time.Now()
	return node.clock
}

// IncrementClock increments a logical clock of the cluster node.
func (node *nodeImpl) IncrementClock() Clock {
	if (ClockMax - ClockDiffrent) <= node.clock {
		node.clock = 0
	} else {
		node.clock += ClockDiffrent
	}
	node.ts = time.Now()
	return node.clock
}

// Clock returns a logical clock of the cluster node.
func (node *nodeImpl) Clock() Clock {
	return node.clock
}

// SetTimestamp sets a phisical timestamp to the cluster node.
func (node *nodeImpl) SetTimestamp(ts time.Time) {
	node.ts = ts
}

// Timestamp returns a phisical timestamp of the cluster node.
func (node *nodeImpl) Timestamp() time.Time {
	return node.ts
}

// SetStatus sets a status to the cluster node.
func (node *nodeImpl) SetStatus(status NodeStatus) {
	node.status = status
}

// Status returns a status of the cluster node.
func (node *nodeImpl) Status() NodeStatus {
	return node.status
}
