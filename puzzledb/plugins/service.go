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

// Service represents a plugin service.
type Service interface {
	// SetConfig sets the specified config.
	SetConfig(config config.Config)
	// ServiceType returns the service type.
	ServiceType() ServiceType
	// ServiceName returns the service name.
	ServiceName() string
	// Start starts the service
	Start() error
	// Stop stops the service
	Stop() error
}
