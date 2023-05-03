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

package config

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

type viperConfig struct {
}

// NewConfigWith creates a new configuration with the specified product name.
func NewConfigWith(productName string) Config {
	viper.SetConfigName(productName)
	viper.SetEnvPrefix(strings.ToUpper(productName))
	viper.AutomaticEnv()
	return &viperConfig{}
}

// Set sets a value to the specified path.
func (conf *viperConfig) Set(path string, v any) error {
	return nil
}

func (conf *viperConfig) Get(path string) (any, error) {
	v := viper.Get(path)
	if v == nil {
		return nil, newErrNotFound(path)
	}
	return v, nil
}

func (conf *viperConfig) GetString(path string) (string, error) {
	v := viper.GetString(path)
	if len(v) == 0 {
		return "", newErrNotFound(path)
	}
	return v, nil
}

func (conf *viperConfig) GetInt(path string) (int, error) {
	v := viper.GetInt(path)
	if v == 0 {
		return 0, newErrNotFound(path)
	}
	return v, nil
}

func (conf *viperConfig) GetBool(path string) (bool, error) {
	v := viper.GetString(path)
	if len(v) == 0 {
		return false, newErrNotFound(path)
	}
	return strconv.ParseBool(v)
}

// String returns a string representation of the configuration.
func (conf *viperConfig) String() string {
	var s string
	keys := viper.AllKeys()
	sort.Strings(keys)
	for _, key := range keys {
		value := viper.Get(key)
		s += fmt.Sprintf("%s: %v\n", key, value)
	}
	return strings.TrimSuffix(s, "\n")
}
