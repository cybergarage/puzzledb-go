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

package server

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/cybergarage/go-logger/log"
	"github.com/cybergarage/puzzledb-go/puzzledb"
	"github.com/cybergarage/puzzledb-go/puzzledb/config"
	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{ // nolint:exhaustruct
	Use:               strings.ToLower(puzzledb.ProductName) + "-server",
	Version:           puzzledb.Version,
	Short:             "",
	Long:              "",
	DisableAutoGenTag: true,
}

var versionCmd = &cobra.Command{ // nolint:exhaustruct
	Use:               "version",
	Short:             "Print " + strings.ToLower(puzzledb.ProductName) + " version",
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(puzzledb.Version)
		os.Exit(0)
	},
}

func GetRootCommand() *cobra.Command {
	return rootCmd
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Errorf("%s couldn't be executed (%s)", puzzledb.ProductName, err.Error())
		os.Exit(1)
	}

	var conf config.Config
	if cfgFile != "" {
		fileConf, err := puzzledb.NewConfigWithFile(cfgFile)
		if err != nil {
			log.Errorf("%s couldn't be loaded (%s)", cfgFile, err.Error())
			os.Exit(1)
		}
		conf = fileConf
	} else {
		localConf, err := puzzledb.NewConfigWithPath(".")
		if err == nil {
			conf = localConf
		}
	}

	var server *puzzledb.Server
	if conf != nil {
		server = puzzledb.NewServerWithConfig(conf)
	} else {
		server = puzzledb.NewServer()
	}

	if err := server.Start(); err != nil {
		log.Errorf("%s couldn't be started (%s)", puzzledb.ProductName, err.Error())
		os.Exit(1)
	}

	sigCh := make(chan os.Signal, 1)

	signal.Notify(sigCh,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM)

	exitCh := make(chan int)

	go func() {
		for {
			s := <-sigCh
			switch s {
			case syscall.SIGHUP:
				log.Infof("caught %s, restarting...", s.String())
				if err := server.Restart(); err != nil {
					log.Errorf("%s couldn't be restarted (%s)", puzzledb.ProductName, err.Error())
					os.Exit(1)
				}
			case syscall.SIGINT, syscall.SIGTERM:
				log.Infof("caught %s, terminating...", s.String())
				if err := server.Stop(); err != nil {
					log.Errorf("%s couldn't be terminated (%s)", puzzledb.ProductName, err.Error())
					os.Exit(1)
				}
				exitCh <- 0
			}
		}
	}()

	code := <-exitCh

	os.Exit(code)
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is .puzzledb.yaml)")
}
