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
	"bufio"
	"bytes"
	_ "embed"
	"os"

	"github.com/cybergarage/puzzledb-go/puzzledb/config"
	"github.com/spf13/viper"
)

type configImpl struct {
	config.Config
}

// NewConfig returns a new configuration.
func NewConfig() (Config, error) {
	conf := config.NewConfigWith(ProductName)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return NewConfigWith(conf), nil
}

// NewConfigWith returns a new configuration with the specified configuration.
func NewConfigWith(conf config.Config) Config {
	return &configImpl{
		Config: conf,
	}
}

// NewConfigWithPath returns a new configuration with the specified path.
func NewConfigWithPath(path string) (Config, error) {
	conf := config.NewConfigWith(ProductName)
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return NewConfigWith(conf), nil
}

func NewConfigWithPaths(paths ...string) (Config, error) {
	conf := config.NewConfigWith(ProductName)
	for _, path := range paths {
		viper.AddConfigPath(path)
	}
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return NewConfigWith(conf), nil
}

// NewConfigWithString returns a new configuration with the specified string.
func NewConfigWithString(conString string) (Config, error) {
	conf := config.NewConfigWith(ProductName)
	if err := viper.ReadConfig(bytes.NewBufferString(conString)); err != nil {
		return nil, err
	}
	return NewConfigWith(conf), nil
}

// NewConfigWithFile returns a new configuration with the specified file.
func NewConfigWithFile(confFile string) (Config, error) {
	f, err := os.Open(confFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	conf := config.NewConfigWith(ProductName)
	if err := viper.ReadConfig(bufio.NewReader(f)); err != nil {
		return nil, err
	}
	return NewConfigWith(conf), nil
}

// SetConfig sets a manager configuration.
func (conf *configImpl) SetConfig(c config.Config) {
	conf.Config = c
}

// Object returns a raw configuration object.
func (conf *configImpl) Object() config.Config {
	return conf.Config
}

// String returns a string representation of the configuration.
func (conf *configImpl) String() string {
	if conf.Config == nil {
		return ""
	}
	return conf.Config.String()
}
