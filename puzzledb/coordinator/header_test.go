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

package coordinator

import (
	"testing"
)

func TestKeyHeader(t *testing.T) {
	type expected struct {
		tp  Category
		ver Version
		doc Format
		idx IndexFormat
	}
	testKeyHeaders := []struct {
		header   KeyHeader
		expected expected
	}{
		{
			header: StateObjectKeyHeader,
			expected: expected{
				tp:  StateHeaderObject,
				ver: V1,
				doc: CBOR,
				idx: IndexFormat(0),
			},
		},
		{
			header: MessageObjectKeyHeader,
			expected: expected{
				tp:  MessageHeaderObject,
				ver: V1,
				doc: CBOR,
				idx: IndexFormat(0),
			},
		},
	}
	for _, key := range testKeyHeaders {
		if key.header.Category() != key.expected.tp {
			t.Errorf("%v != %v", key.header.Category(), key.expected.tp)
		}
		if key.header.Version() != key.expected.ver {
			t.Errorf("%v != %v", key.header.Version(), key.expected.ver)
		}
		if key.expected.doc != Format(0) {
			if key.header.Format() != key.expected.doc {
				t.Errorf("%v != %v", key.header.Format(), key.expected.doc)
			}
		}
		if key.expected.idx != IndexFormat(0) {
			if key.header.IndexFormat() != key.expected.idx {
				t.Errorf("%v != %v", key.header.IndexFormat(), key.expected.idx)
			}
		}
	}
}
