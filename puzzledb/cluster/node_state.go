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
	// NodeJoining represents a joining node status.
	NodeJoining
	// NodeUp represents a running node status.
	NodeUp
	// NodeDown represents a down node status.
	NodeDown
	// NodeLeaving represents a leaving node status.
	NodeLeaving
	// NodeExiting represents an exiting node status.
	NodeExiting
	// NodeRemoved represents a removed node status.
	NodeRemoved
	// NodeAborted represents an aborted node status.
	NodeAborted
	// NodeUnreachable represents an unreachable node status.
	NodeUnreachable
)

var processStatuses = map[NodeStatus]string{
	NodeUnknown:     "unknown",
	NodeIdle:        "idle",
	NodeJoining:     "joining",
	NodeUp:          "up",
	NodeDown:        "down",
	NodeLeaving:     "leaving",
	NodeExiting:     "exiting",
	NodeRemoved:     "removed",
	NodeAborted:     "aborted",
	NodeUnreachable: "unreachable",
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
