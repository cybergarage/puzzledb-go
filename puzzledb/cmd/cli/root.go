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
	"strings"

	"github.com/cybergarage/puzzledb-go/puzzledb"
	"github.com/spf13/cobra"
)

var gRPCHost string
var gRPCPort int

var rootCmd = &cobra.Command{ // nolint:exhaustruct
	Use:               strings.ToLower(puzzledb.PackageName) + "-cli",
	Version:           puzzledb.Version,
	Short:             "",
	Long:              "",
	DisableAutoGenTag: true,
}

func GetRootCommand() *cobra.Command {
	return rootCmd
}

func Execute() error {
	client := puzzledb.NewClient()
	client.SetHost(gRPCHost)
	client.SetPort(gRPCPort)

	if err := client.Open(); err != nil {
		return err
	}

	SetClient(client)

	defer func() {
		if err := client.Close(); err != nil {
			fmt.Println(err)
			return
		}
	}()

	return rootCmd.Execute()
}

func init() { // nolint:gochecknoinits
	rootCmd.PersistentFlags().StringVar(&gRPCHost, "host", "localhost", fmt.Sprintf("gRPC host or address for a %v instance", puzzledb.ProductName))
	rootCmd.PersistentFlags().IntVar(&gRPCPort, "port", puzzledb.DefaultGrpcPort, fmt.Sprintf("gRPC port number for a %v instance", puzzledb.ProductName))
}
