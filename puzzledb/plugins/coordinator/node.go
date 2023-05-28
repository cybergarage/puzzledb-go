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
	"time"

	"github.com/cybergarage/puzzledb-go/puzzledb/cluster"
	"github.com/cybergarage/puzzledb-go/puzzledb/coordinator"
	"github.com/google/uuid"
)

// NodeObject represents a store node state object.
type NodeObject struct {
	ID      string
	Cluster string
	Host    string
	Clock   uint64
	Time    time.Time
	Status  string
}

// NewNodeWith returns a new node with the specified node object.
func NewNodeWith(obj *NodeObject) (cluster.Node, error) {
	uuid, err := uuid.Parse(obj.ID)
	if err != nil {
		return nil, err
	}
	node := cluster.NewNode()
	node.SetCluster(obj.Cluster)
	node.SetID(uuid)
	node.SetHost(obj.Host)
	node.SetClock(obj.Clock)
	node.SetTimestamp(obj.Time)
	node.SetStatus(cluster.NewNodeStatusWith(obj.Status))
	return node, nil
}

// NewNodeScanKey returns a new scan node key to get all node states.
func NewNodeScanKey() coordinator.Key {
	return coordinator.NewKeyWith(coordinator.StateObjectKeyHeader[:], byte(NodeState))
}

// NewClusterScanKeyWith returns a new scan node key to get all node states with the specified cluster.
func NewClusterScanKeyWith(cluster string) coordinator.Key {
	return coordinator.NewKeyWith(coordinator.StateObjectKeyHeader[:], byte(NodeState), cluster)
}

// NewNodeKeyWith returns a new node key with the specified node.
func NewNodeKeyWith(node cluster.Node) coordinator.Key {
	return coordinator.NewKeyWith(coordinator.StateObjectKeyHeader[:], byte(NodeState), node.Cluster(), node.ID().String())
}

// NewNodeObject returns a new node object.
func NewNodeObject() *NodeObject {
	return &NodeObject{
		Cluster: "",
		ID:      "",
		Host:    "",
		Clock:   0,
		Time:    time.Now(),
		Status:  "",
	}
}

// NewNodeObjectWith returns a new node object with the specified node.
func NewNodeObjectWith(node cluster.Node) *NodeObject {
	return &NodeObject{
		Cluster: node.Cluster(),
		ID:      node.ID().String(),
		Host:    node.Host(),
		Clock:   uint64(node.Clock()),
		Time:    node.Timestamp(),
		Status:  node.Status().String(),
	}
}
