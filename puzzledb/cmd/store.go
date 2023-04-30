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

package cmd

import (
	"github.com/spf13/cobra"
)

func init() { // nolint:gochecknoinits
	rootCmd.AddCommand(listCmd)
	listCmd.AddCommand(listDatabasesCmd)
	listCmd.AddCommand(listCollectionsCmd)
}

var listCmd = &cobra.Command{ // nolint:exhaustruct
	Use:   "list",
	Short: "List store resources",
	Long:  `List store resources such as databases and collections.`,
}

var listDatabasesCmd = &cobra.Command{ // nolint:exhaustruct
	Use:   "databases",
	Short: "List databases",
	Long:  `List all the databases.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Add logic to list databases
		// fmt.Println("Listing databases...")
	},
}

// nolint:forbidigo
var listCollectionsCmd = &cobra.Command{ // nolint:exhaustruct
	Use:   "collections [database]",
	Short: "List collections in a database",
	Long:  `List all the collections in the specified database.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Add logic to list collections in the specified database
		// fmt.Printf("Listing collections in database %s...\n", args[0])
	},
}
