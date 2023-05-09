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

package metrics

import (
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins"
)

// Service represents a query service.
type Service interface {
	plugins.Service
	// SetPort sets the specified port number.
	SetPort(port int)
}

type BaseService struct {
	*plugins.Config
}

// NewBaseService returns a new tracer base service.
func NewBaseService() *BaseService {
	return &BaseService{
		Config: plugins.NewConfig(),
	}
}

// ServiceType returns the plug-in service type.
func (service *BaseService) ServiceType() plugins.ServiceType {
	return plugins.MetricsService
}
