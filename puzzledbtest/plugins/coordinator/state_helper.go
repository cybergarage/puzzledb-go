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
	"errors"
	"fmt"
	"testing"

	"github.com/cybergarage/puzzledb-go/puzzledb/cluster"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/coordinator"
)

func truncateCoordinatorStore(coord coordinator.Service) error {
	txn, err := coord.Transact()
	if err != nil {
		return err
	}
	err = txn.Truncate()
	if err != nil {
		return errors.Join(err, txn.Cancel())
	}
	return txn.Commit()
}

// CoordinatorClusterTest runs coordinator cluster conformance tests against the specified services.
func CoordinatorClusterTest(t *testing.T, coords []coordinator.Service) { //nolint:gocognit,gci,gocyclo,gosec,maintidx
	t.Helper()

	for _, coord := range coords {
		if err := truncateCoordinatorStore(coord); err != nil {
			t.Error(err)
			return
		}
	}

	testCluster := "test_cluster"

	for n, coord := range coords {
		coord.SetCluster(testCluster)
		coord.SetHost(fmt.Sprintf("coord%d", n))
	}

	for _, coord := range coords {
		coord.SetStatus(cluster.NodeUp)
		err := coord.SetNodeState(coord)
		if err != nil {
			t.Error(err)
			return
		}
	}

	for _, coord := range coords {
		cluster, err := coord.GetClusterState(testCluster)
		if err != nil {
			t.Error(err)
			return
		}

		clusterNodes := cluster.Nodes()
		if len(clusterNodes) != len(coords) {
			t.Errorf("cluster node count (%d) is invalid", len(clusterNodes))
			return
		}

		for _, clusterNode := range clusterNodes {
			if clusterNode.Cluster() != testCluster {
				t.Errorf("cluster node cluster (%s) is invalid", clusterNode.Cluster())
				return
			}
			foundNode := false
			for _, coord := range coords {
				if coord.Equals(clusterNode) {
					foundNode = true
					break
				}
			}
			if !foundNode {
				t.Errorf("cluster node (%v) is not found", clusterNode)
				return
			}
		}
	}
}
