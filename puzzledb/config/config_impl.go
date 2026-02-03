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
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	return &viperConfig{}
}

// UseConfigFile uses the specified file as the configuration.
func (conf *viperConfig) UsedConfigFile() string {
	return viper.ConfigFileUsed()
}

// LookupConfigObject returns a object value for the specified path.
func (conf *viperConfig) LookupConfigObject(paths ...string) (any, error) {
	path := NewPathWith(paths...)
	v := viper.Get(path)
	if v == nil {
		return nil, newErrNotFound(path)
	}
	return v, nil
}

// LookupConfigString returns a string value for the specified path.
func (conf *viperConfig) LookupConfigString(paths ...string) (string, error) {
	path := NewPathWith(paths...)
	v := viper.GetString(path)
	if len(v) == 0 {
		return "", newErrNotFound(path)
	}
	return v, nil
}

// LookupConfigInt returns an integer value for the specified path.
func (conf *viperConfig) LookupConfigInt(paths ...string) (int, error) {
	path := NewPathWith(paths...)
	v := viper.GetInt(path)
	if v == 0 {
		return 0, newErrNotFound(path)
	}
	return v, nil
}

// LookupConfigBool returns a boolean value for the specified path.
func (conf *viperConfig) LookupConfigBool(paths ...string) (bool, error) {
	path := NewPathWith(paths...)
	v := viper.GetString(path)
	if len(v) == 0 {
		return false, newErrNotFound(path)
	}
	return strconv.ParseBool(v)
}

// UnmarshallConfig unmarshalls the specified path object to the specified object.
func (conf *viperConfig) UnmarshallConfig(paths []string, v any) error {
	path := NewPathWith(paths...)
	return viper.UnmarshalKey(path, v)
}

// SetConfigObject sets a object value for the specified path.
func (conf *viperConfig) SetConfigObject(paths []string, v any) error {
	path := NewPathWith(paths...)
	viper.Set(path, v)
	return nil
}

// SetConfigString sets a string value for the specified path.
func (conf *viperConfig) SetConfigString(paths []string, v string) error {
	path := NewPathWith(paths...)
	viper.Set(path, v)
	return nil
}

// SetConfigInt sets an integer value for the specified path.
func (conf *viperConfig) SetConfigInt(paths []string, v int) error {
	path := NewPathWith(paths...)
	viper.Set(path, v)
	return nil
}

// String returns a string representation of the configuration.
func (conf *viperConfig) String() string {
	var s strings.Builder
	keys := viper.AllKeys()
	sort.Strings(keys)
	for _, key := range keys {
		value := viper.Get(key)
		s.WriteString(fmt.Sprintf("%s: %v\n", key, value))
	}
	return strings.TrimSuffix(s.String(), "\n")
}
