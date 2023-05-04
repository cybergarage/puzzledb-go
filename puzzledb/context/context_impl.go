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
	tracer.Context
}

// NewContext returns a new context.
func NewContext() Context {
	return &ctx{
		Context: tracer.NewNullTracer().StartSpan(""),
	}
}

// NewContextWith returns a new context with the specified tracer context.
func NewContextWith(t tracer.Context) Context {
	return &ctx{
		Context: t,
	}
}
