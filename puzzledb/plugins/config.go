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

package plugins

import (
	"github.com/cybergarage/puzzledb-go/puzzledb/config"
)

const (
	configPlugins = "plugins"
	configDefault = "default"
	configEnabled = "enabled"
)

type Config struct {
	config.Config
}

func NewConfig() *Config {
	return NewConfigWith(nil)
}

func NewConfigWith(config config.Config) *Config {
	return &Config{
		Config: config,
	}
}

// SetConfig sets a manager configuration.
func (conf *Config) SetConfig(config config.Config) {
	conf.Config = conf
}

func (conf *Config) EnabledConfig(t ServiceType) (string, error) {
	return conf.GetString(config.NewPathWith(configPlugins, t.String()))
}
