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

package context

import (
	"github.com/cybergarage/go-tracing/tracer"
)

// Context represents a context interface.
type Context interface {
	// Span returns a current span context.
	Span() tracer.SpanContext
	// SetSpan sets a new root span context.
	SetSpan(span tracer.SpanContext) Context
	// StartSpan starts a new child span.
	StartSpan(name string) bool
	// FinishSpan ends the current span.
	FinishSpan() bool
}
