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
	"fmt"
	"strconv"
	"strings"

	"github.com/cybergarage/puzzledb-go/puzzledb"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var gRPCHost string
var gRPCPort int

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{ // nolint:exhaustruct
	Use:               "puzzledb-cli",
	Version:           puzzledb.Version,
	Short:             "",
	Long:              "",
	DisableAutoGenTag: true,
}

// GetRootCommand returns the root command.
func GetRootCommand() *cobra.Command {
	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
// nolint:forbidigo
func Execute() error {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			return err
		}

		// Search config in home directory with name ".puzzledb-cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(puzzledb.ProductName + "-cli")
	}

	viper.SetEnvPrefix(strings.ToUpper(puzzledb.ProductName))
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	// Setup the gRPC client

	client := puzzledb.NewClient()
	client.SetHost(gRPCHost)
	client.SetPort(gRPCPort)

	if host := viper.GetString("host"); 0 < len(host) {
		client.SetHost(host)
	}
	if port := viper.GetString("port"); 0 < len(port) {
		i, err := strconv.Atoi(port)
		if err != nil {
			return err
		}
		client.SetPort(i)
	}

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
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.puzzledb-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().StringVar(&gRPCHost, "host", "localhost", fmt.Sprintf("gRPC host or address for a %v instance", puzzledb.ProductName))
	rootCmd.PersistentFlags().IntVar(&gRPCPort, "port", puzzledb.DefaultGrpcPort, fmt.Sprintf("gRPC port number for a %v instance", puzzledb.ProductName))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() { // nolint:gochecknoinits
}
