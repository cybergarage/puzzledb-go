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

func init() {
	listCmd.AddCommand(listConfigCmd)
	getCmd.AddCommand(getConfigCmd)
}

var listConfigCmd = &cobra.Command{ // nolint:exhaustruct
	Use:   "config",
	Short: "List configurations",
	Long:  "List all the configurations.",
	Run: func(cmd *cobra.Command, args []string) {
		dbs, err := GetClient().ListConfig()
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
var getConfigCmd = &cobra.Command{ // nolint:exhaustruct
	Use:   "config [name]",
	Short: "Get configuration",
	Long:  "Get the specified configurateion.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("No name specified")
			return
		}
		name := args[0]
		v, err := GetClient().GetConfig(name)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(v)
	},
}
