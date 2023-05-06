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

package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() { // nolint:gochecknoinits
	listCmd.AddCommand(listMetricCmd)
	getCmd.AddCommand(getMetricCmd)
}

var listMetricCmd = &cobra.Command{ // nolint:exhaustruct
	Use:   "config",
	Short: "List metrics",
	Long:  "List all the metrics.",
	Run: func(cmd *cobra.Command, args []string) {
		dbs, err := GetClient().ListMetric()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		for _, db := range dbs {
			fmt.Println(db)
		}
	},
}

// nolint:forbidigo
var getMetricCmd = &cobra.Command{ // nolint:exhaustruct
	Use:   "config [name]",
	Short: "Get metric",
	Long:  "Get the specified metric.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("No name specified")
			return
		}
		name := args[0]
		v, err := GetClient().GetMetric(name)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(v)
	},
}
