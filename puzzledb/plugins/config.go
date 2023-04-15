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

const (
	configPlugins = "plugins"
	configDefault = "default"
)

// Config represents a configuration interface.
type Config interface {
	// Set sets a value to the specified path.
	Set(path string, v any) error
	// Get returns a value for the specified name.
	Get(path string) (any, error)
	// GetString returns a string value for the specified name.
	GetString(path string) (string, error)
	// GetInt returns an integer value for the specified name.
	GetInt(path string) (int, error)
	// String returns a string representation of the configuration.
	String() string
}
