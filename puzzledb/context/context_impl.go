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

type ctx struct {
	spans []tracer.SpanContext
}

// NewContext returns a new context.
func NewContext() Context {
	return &ctx{
		spans: nil,
	}
}

// Span returns a current span context.
func (ctx *ctx) Span() tracer.SpanContext {
	if len(ctx.spans) == 0 {
		return nil
	}
	return ctx.spans[len(ctx.spans)-1]
}

// SetSpan sets a new root span context.
func (ctx *ctx) SetSpan(span tracer.SpanContext) Context {
	ctx.spans = make([]tracer.SpanContext, 0)
	ctx.pushSpan(span)
	return ctx
}

// StartSpan starts a new child span.
func (ctx *ctx) StartSpan(name string) bool {
	if len(ctx.spans) == 0 {
		return false
	}
	span := ctx.spans[len(ctx.spans)-1]
	childSpan := span.Span().StartSpan(name)
	ctx.pushSpan(childSpan)
	return true
}

// FinishSpan ends the current span.
func (ctx *ctx) FinishSpan() bool {
	span := ctx.popSpan()
	if span == nil {
		return false
	}
	span.Span().Finish()
	return true
}

func (ctx *ctx) popSpan() tracer.SpanContext {
	if len(ctx.spans) == 0 {
		return nil
	}
	lastSpanIndex := len(ctx.spans) - 1
	span := ctx.spans[lastSpanIndex]
	ctx.spans = ctx.spans[:lastSpanIndex]
	return span
}

func (ctx *ctx) pushSpan(span tracer.SpanContext) {
	ctx.spans = append(ctx.spans, span)
}
