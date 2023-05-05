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
	"github.com/cybergarage/go-logger/log"
	"github.com/cybergarage/go-tracing/tracer"
	"github.com/cybergarage/go-tracing/tracer/ot"
	"github.com/cybergarage/go-tracing/tracer/otel"
)

const (
	opentelemetryConfig = "opentelemetry"
	opentracingConfig   = "opentracing"
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

func (t *Tracer) EnabledConfig(name string) (bool, error) {
	return t.Server.Config.GetBool(tracingConfig, tracerConfig, name, enabledConfig)
}

func (t *Tracer) DefaultConfig() (string, error) {
	return t.Server.Config.GetString(tracingConfig, defaultConfig)
}

func (t *Tracer) EndpointConfig(name string) (string, error) {
	return t.Server.Config.GetString(tracingConfig, tracerConfig, name, endpointConfig)
}

func (t *Tracer) Start() error {
	t.tracer = tracer.NewNullTracer()

	tracerName, err := t.DefaultConfig()
	if err != nil {
		return err
	}
	var tracer tracer.Tracer
	switch tracerName {
	case opentelemetryConfig:
		tracer = otel.NewTracer()
	case opentracingConfig:
		tracer = ot.NewTracer()
	default:
		return newErrNotFound(tracerName)
	}

	endpoint, err := t.EndpointConfig(tracerName)
	if err != nil {
		return err
	}

	tracer.SetServiceName(ProductName)
	tracer.SetEndpoint(endpoint)
	if err := tracer.Start(); err != nil {
		return err
	}

	t.tracer = tracer

	log.Infof("%s tracer (%s) started", tracerName, endpoint)

	return nil
}

func (t *Tracer) Stop() error {
	tracerName, err := t.DefaultConfig()
	if err != nil {
		return err
	}

	endpoint, err := t.EndpointConfig(tracerName)
	if err != nil {
		return err
	}

	if err := t.tracer.Stop(); err != nil {
		return err
	}

	log.Infof("%s tracer (%s) started", tracerName, endpoint)

	return nil
}
