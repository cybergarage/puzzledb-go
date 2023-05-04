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
	"fmt"
	"testing"

	"github.com/cybergarage/go-tracing/tracer"
)

func TestContext(t *testing.T) {
	loopCnt := 10

	ctx := NewContext().SetSpan(tracer.NullTracer.StartSpan("root"))

	for n := 0; n < loopCnt; n++ {
		name := fmt.Sprintf("span%d", n)
		if !ctx.StartSpan(name) {
			t.Errorf("ctx.StartSpan(%v)", name)
		}
	}

	for n := 0; n < loopCnt; n++ {
		name := fmt.Sprintf("span%d", n)
		if !ctx.FinishSpan() {
			t.Errorf("ctx.FinishSpan(%v)", name)
		}
	}
}
