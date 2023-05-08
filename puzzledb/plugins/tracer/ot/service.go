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

package ot

import (
	"github.com/cybergarage/go-tracing/tracer"
	"github.com/cybergarage/go-tracing/tracer/ot"
	tracer_plugin "github.com/cybergarage/puzzledb-go/puzzledb/plugins/tracer"
)

// Service represents a tracer service.
type Service struct {
	*tracer_plugin.BaseService
	tracer tracer.Tracer
}

// NewTracer returns a new tracer service.
func NewService() *Service {
	return &Service{
		BaseService: tracer_plugin.NewBaseService(),
		tracer:      ot.NewTracer(),
	}
}

// ServiceName returns the plug-in service name.
func (t *Service) ServiceName() string {
	return "opentracing"
}

// Tracer returns an enabled tracer.
func (t *Service) Tracer() tracer.Tracer {
	return t.tracer
}

func (t *Service) Start() error {
	if err := t.tracer.Start(); err != nil {
		return err
	}

	return nil
}

func (t *Service) Stop() error {
	if err := t.tracer.Stop(); err != nil {
		return err
	}

	return nil
}
