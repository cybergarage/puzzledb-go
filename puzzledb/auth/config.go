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

package auth

import (
	"github.com/cybergarage/puzzledb-go/puzzledb/config"
)

// Config represents a configuration for authenticator.
type Config struct {
	Type     string `mapstructure:"type"`
	Enabled  bool   `mapstructure:"enabled"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

// NewConfig returns a new configuration for authenticator.
func NewConfig() *Config {
	return &Config{}
}

// NewConfigWith returns a new configuration for authenticator with the specified configuration.
func NewConfigWith(config config.Config, path ...string) ([]Config, error) {
	var configs []Config
	err := config.UnmarshallConfig(path, &configs)
	return configs, err
}
