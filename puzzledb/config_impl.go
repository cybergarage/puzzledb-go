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
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type viperConfig struct {
}

func newConfig() Config {
	viper.SetConfigName(ProductName)
	viper.SetConfigType("yaml")
	return &viperConfig{}
}

// NewConfig returns a new configuration.
func NewConfig() (Config, error) {
	conf := newConfig()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return conf, nil
}

// NewConfigWithPath returns a new configuration with the specified path.
func NewConfigWithPath(path string) (Config, error) {
	conf := newConfig()
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return conf, nil
}

// NewConfigWithString returns a new configuration with the specified string.
func NewConfigWithString(config string) (Config, error) {
	conf := newConfig()
	if err := viper.ReadConfig(bytes.NewBuffer([]byte(config))); err != nil {
		return nil, err
	}
	return conf, nil
}

// Set sets a value to the specified path.
func (conf *viperConfig) Set(path []string, v any) error {
	return nil
}

func (conf *viperConfig) Get(name ...string) (any, error) {
	path := strings.Join(name, ".")
	v := viper.Get(path)
	if v == nil {
		return nil, newErrNotFound(path)
	}
	return v, nil
}

func (conf *viperConfig) GetString(name ...string) (string, error) {
	path := strings.Join(name, ".")
	v := viper.GetString(path)
	if v == "" {
		return "", newErrNotFound(path)
	}
	return v, nil
}

func (conf *viperConfig) GetInt(name ...string) (int, error) {
	path := strings.Join(name, ".")
	v := viper.GetInt(path)
	if v == 0 {
		return 0, newErrNotFound(path)
	}
	return v, nil
}

// String returns a string representation of the configuration.
func (conf *viperConfig) String() string {
	var s string
	for _, key := range viper.AllKeys() {
		value := viper.Get(key)
		s += fmt.Sprintf("%s: %v\n", key, value)
	}
	return strings.TrimSuffix(s, "\n")
}
