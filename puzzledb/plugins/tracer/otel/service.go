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
	"github.com/cybergarage/go-logger/log"
	"github.com/cybergarage/go-tracing/tracer"
	"github.com/cybergarage/go-tracing/tracer/otel"
)

// Tracer represents a tracer.
type Tracer struct {
	tracer tracer.Tracer
}

// NewTracer returns a new tracer.
func NewTracerWith() *Tracer {
	return &Tracer{
		tracer: otel.NewTracer(),
	}
}

// ServiceName returns the plug-in service name.
func (t *Tracer) ServiceName() string {
	return "opentelemetry"
}

// Tracer returns an enabled tracer.
func (t *Tracer) Tracer() tracer.Tracer {
	return t.tracer
}

func (t *Tracer) Start() error {
	if err := t.tracer.Start(); err != nil {
		return err
	}

	log.Infof("%s tracer (%s) started", t.ServiceName(), t.tracer.Endpoint())

	return nil
}

func (t *Tracer) Stop() error {
	if err := t.tracer.Stop(); err != nil {
		return err
	}

	log.Infof("%s tracer (%s) terminated", t.ServiceName(), t.tracer.Endpoint())

	return nil
}
