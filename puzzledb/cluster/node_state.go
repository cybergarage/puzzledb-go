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

// NodeStatus represents a cluster node status.
type NodeStatus int

const (
	// NodeUnknown represents an unknown node status.
	NodeUnknown NodeStatus = iota
	// NodeIdle represents an idle node status.
	NodeIdle
	// NodeStarting represents a starting node status.
	NodeStarting
	// NodeRunning represents a running node status.
	NodeRunning
	// NodeStopping represents a stopping node status.
	NodeStopping
	// NodeStopped represents a stopped node status.
	NodeStopped
	// NodeAborted represents an aborted node status.
	NodeAborted
)

var processStatuses = map[NodeStatus]string{
	NodeUnknown:  "unknown",
	NodeIdle:     "idle",
	NodeStarting: "starting",
	NodeRunning:  "running",
	NodeStopping: "stopping",
	NodeStopped:  "stopped",
	NodeAborted:  "aborted",
}

// NewNodeStatusWith returns a new node status with the specified string.
func NewNodeStatusWith(s string) NodeStatus {
	for status, statusString := range processStatuses {
		if statusString == s {
			return status
		}
	}
	return NodeUnknown
}

// String represents a string of the node status.
func (t NodeStatus) String() string {
	s, ok := processStatuses[t]
	if ok {
		return s
	}
	return processStatuses[NodeUnknown]
}