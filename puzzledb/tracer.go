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
	"github.com/cybergarage/go-tracing/tracer"
	"github.com/cybergarage/puzzledb-go/puzzledb/config"
)

// Tracer represents a tracer.
type Tracer struct {
	*Server
	tracer tracer.Tracer
}

// NewTracer returns a new tracer.
func NewTracerWith(server *Server) *Tracer {
	return &Tracer{
		Server: server,
		tracer: tracer.NewNullTracer(),
	}
}

// Tracer returns an enabled tracer.
func (t *Tracer) Tracer() tracer.Tracer {
	return t.tracer
}

func (t *Tracer) EnabledConfig() (bool, error) {
	return t.Server.Config.GetBool(config.NewPathWith(tracingConfig, enabledConfig))
}

func (t *Tracer) Start() error {
	enabled, err := t.EnabledConfig()
	if err != nil {
		return err
	}
	if enabled {

	}
	t.tracer.SetServiceName(ProductName)
	return nil
}

func (t *Tracer) Stop() error {
	return t.tracer.Stop()
}
