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

// Config represents a configuration interface.
type Config interface {
	// UseConfigFile uses the specified file as the configuration.
	UsedConfigFile() string
	// GetConfigObject returns a object value for the specified path.
	GetConfigObject(paths ...string) (any, error)
	// GetConfigString returns a string value for the specified path.
	GetConfigString(paths ...string) (string, error)
	// GetConfigInt returns an integer value for the specified path.
	GetConfigInt(paths ...string) (int, error)
	// GetConfigBool returns a boolean value for the specified path.
	GetConfigBool(paths ...string) (bool, error)
	// UnmarshallConfig unmarshalls the specified path object to the specified object.
	UnmarshallConfig(paths []string, v any) error
	// SetConfigObject sets a object value to the specified path.
	SetConfigObject(paths []string, v any) error
	// SetConfigString sets a string value to the specified path.
	SetConfigString(paths []string, v string) error
	// SetConfigInt sets an integer value to the specified path.
	SetConfigInt(paths []string, v int) error
	// String returns a string representation of the configuration.
	String() string
}
