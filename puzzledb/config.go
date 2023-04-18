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

package puzzledb

import (
	"bytes"
	"strings"

	"github.com/cybergarage/puzzledb-go/puzzledb/config"
	"github.com/spf13/viper"
)

const (
	pluginsConfig = "plugins"
	queryConfig   = "query"
	portConfig    = "port"
)

type Config struct {
	config.Config
}

// NewConfig returns a new configuration.
func NewConfig() (config.Config, error) {
	conf := config.NewConfigWith(ProductName)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return conf, nil
}

// NewConfigWith returns a new configuration with the specified configuration.
func NewConfigWith(config config.Config) *Config {
	return &Config{
		Config: config,
	}
}

// NewConfigWithPath returns a new configuration with the specified path.
func NewConfigWithPath(path string) (config.Config, error) {
	conf := config.NewConfigWith(ProductName)
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return conf, nil
}

func NewConfigWithPaths(paths ...string) (config.Config, error) {
	conf := config.NewConfigWith(ProductName)
	for _, path := range paths {
		viper.AddConfigPath(path)
	}
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return conf, nil
}

// NewConfigWithString returns a new configuration with the specified string.
func NewConfigWithString(conString string) (config.Config, error) {
	conf := config.NewConfigWith(ProductName)
	if err := viper.ReadConfig(bytes.NewBuffer([]byte(conString))); err != nil {
		return nil, err
	}
	return conf, nil
}

// Port returns a port number for the specified name.
func (conf *Config) Port(name string) (int, error) {
	return conf.Config.GetInt(strings.Join([]string{pluginsConfig, queryConfig, name, portConfig}, "."))
}

func (conf *Config) String() string {
	if conf.Config == nil {
		return ""
	}
	return conf.Config.String()
}
