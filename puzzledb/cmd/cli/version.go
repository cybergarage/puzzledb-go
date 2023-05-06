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
	getCmd.AddCommand(getVersionCmd)
}

var getVersionCmd = &cobra.Command{ // nolint:exhaustruct
	Use:   "version",
	Short: "Get version",
	Long:  "Get version string.",
	Run: func(cmd *cobra.Command, args []string) {
		ver, err := GetClient().GetVersion()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(ver)
	},
}
