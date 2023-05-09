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

// Config represents a plug-in configuration interface.
type Config interface {
	// Get returns a value for the specified name.
	Get(paths ...string) (any, error)
	// GetString returns a string value for the specified name.
	GetString(paths ...string) (string, error)
	// GetInt returns an integer value for the specified name.
	GetInt(paths ...string) (int, error)
	// GetBool returns a boolean value for the specified name.
	GetBool(paths ...string) (bool, error)
	// String returns a string representation of the configuration.
	SetConfig(c config.Config)
	// Object returns a raw configuration object.
	Object() config.Config
}
