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

/*
	puzzledb-server is the server process for PuzzleDB.
	NAME
	 puzzledb-server

	SYNOPSIS
	 puzzledb-server [OPTIONS]

	OPTIONS
	-v      : Enable verbose output.
	-p      : Enable profiling.

	RETURN VALUE
	  Return EXIT_SUCCESS or EXIT_FAILURE
*/

package main

import (
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/cybergarage/go-logger/log"
	"github.com/cybergarage/puzzledb-go/puzzledb"
)

const (
	prgName = " puzzledb-server"
)

func main() {
	isDebugEnabled := flag.Bool("d", false, "enable debugging log output")
	isProfileEnabled := flag.Bool("p", false, "enable profiling server")
	flag.Parse()

	logLevel := log.LevelTrace
	if *isDebugEnabled {
		logLevel = log.LevelDebug
	}
	log.SetSharedLogger(log.NewStdoutLogger(logLevel))

	if *isProfileEnabled {
		go func() {
			// nolint: gosec
			http.ListenAndServe("localhost:6060", nil)
		}()
	}

	var server *puzzledb.Server

	conf, err := puzzledb.NewConfigWithPath(".")
	if err == nil {
		log.Infof("%s couldn't load the configuration (%s)", prgName, err.Error())
	}

	server = puzzledb.NewServerWithConfig(conf)
	if err := server.Start(); err != nil {
		log.Errorf("%s couldn't be started (%s)", prgName, err.Error())
		os.Exit(1)
	}
	log.Infof("%s started", prgName)

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
					log.Errorf("%s couldn't be restarted (%s)", prgName, err.Error())
					os.Exit(1)
				}
			case syscall.SIGINT, syscall.SIGTERM:
				log.Infof("caught %s, terminating...", s.String())
				if err := server.Stop(); err != nil {
					log.Errorf("%s couldn't be terminated (%s)", prgName, err.Error())
					os.Exit(1)
				}
				exitCh <- 0
			}
		}
	}()

	code := <-exitCh

	log.Infof("%s terminated", prgName)

	os.Exit(code)
}
