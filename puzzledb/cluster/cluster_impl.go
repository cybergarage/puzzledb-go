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

const (
	DefaultClusterName = "puzzledb"
)

type clusterImpl struct {
	name  string
	nodes []Node
}

// NewCluster returns a new cluster.
func NewCluster() Cluster {
	return &clusterImpl{
		name:  DefaultClusterName,
		nodes: make([]Node, 0),
	}
}

func NewClusterWith(name string, nodes []Node) Cluster {
	return &clusterImpl{
		name:  name,
		nodes: nodes,
	}
}

// Name returns a name of the cluster.
func (cluster *clusterImpl) Name() string {
	return cluster.name
}

// Nodes returns nodes in the cluster.
func (cluster *clusterImpl) Nodes() []Node {
	return cluster.nodes
}
