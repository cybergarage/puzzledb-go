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

package puzzledbtest

import (
	_ "embed"

	"github.com/cybergarage/puzzledb-go/puzzledb"
)

//go:embed puzzledb.yaml
var testConfigString string

type Config = puzzledb.Config

// NewConfig returns a new configuration.
func NewConfig() Config {
	conf, err := puzzledb.NewConfigWithString(testConfigString)
	if err != nil {
		panic(err)
	}
	return conf
}

// NewConfigWithString returns a new configuration with the specified string.
func NewConfigWithString(conString string) (Config, error) {
	return puzzledb.NewConfigWithString(conString)
}
