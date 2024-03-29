// Copyright (C) 2022 PuzzleDB Contributors.
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
	"github.com/cybergarage/puzzledb-go/puzzledb/config"
)

//go:embed puzzledb.yaml
var testConfigString string

type Config struct {
	puzzledb.Config
}

// NewConfigWith returns a new configuration with the specified configuration.
func NewConfigWith(conf config.Config) *Config {
	return &Config{
		Config: puzzledb.NewConfigWith(conf),
	}
}

// NewConfigWithString returns a new configuration with the specified string.
func NewConfigWithString(conString string) (*Config, error) {
	conf, err := puzzledb.NewConfigWithString(conString)
	if err != nil {
		return nil, err
	}
	return NewConfigWith(conf), nil
}

// GetConfigObject overrides the GetConfigObject method of the Config interface for testing.
func (conf *Config) GetConfigObject(paths ...string) (any, error) {
	return conf.Config.GetConfigObject(paths...)
}

// GetConfigString overrides the GetConfigString method of the Config interface for testing.
func (conf *Config) GetConfigString(paths ...string) (string, error) {
	return conf.Config.GetConfigString(paths...)
}

// GetConfigInt overrides the GetConfigInt method of the Config interface for testing.
func (conf *Config) GetConfigInt(paths ...string) (int, error) {
	return conf.Config.GetConfigInt(paths...)
}

// GetConfigBool overrides the GetConfigBool method of the Config interface for testing.
func (conf *Config) GetConfigBool(paths ...string) (bool, error) {
	return conf.Config.GetConfigBool(paths...)
}
