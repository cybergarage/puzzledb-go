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

package kv

import (
	"testing"
)

func TestKeyHeader(t *testing.T) {
	type expected struct {
		tp  HeaderType
		ver Version
		doc DocumentType
		idx IndexType
	}
	testKeyHeaders := []struct {
		header   KeyHeader
		expected expected
	}{
		{
			header: DatabaseKeyHeader,
			expected: expected{
				tp:  DatabaseObject,
				ver: V1,
				doc: CBOR,
				idx: IndexType(0),
			},
		},
		{
			header: CollectionKeyHeader,
			expected: expected{
				tp:  CollectionObject,
				ver: V1,
				doc: CBOR,
				idx: IndexType(0),
			},
		},
		{
			header: DocumentKeyHeader,
			expected: expected{
				tp:  DocumentObject,
				ver: V1,
				doc: CBOR,
				idx: IndexType(0),
			},
		},
		{
			header: IndexKeyHeader,
			expected: expected{
				tp:  IndexObject,
				ver: V1,
				doc: DocumentType(0),
				idx: SecondaryIndex,
			},
		},
	}
	for _, key := range testKeyHeaders {
		if key.header.Type() != key.expected.tp {
			t.Errorf("%v != %v", key.header.Type(), key.expected.tp)
		}
		if key.header.Version() != key.expected.ver {
			t.Errorf("%v != %v", key.header.Version(), key.expected.ver)
		}
		if key.expected.doc != DocumentType(0) {
			if key.header.DocumentType() != key.expected.doc {
				t.Errorf("%v != %v", key.header.DocumentType(), key.expected.doc)
			}
		}
		if key.expected.idx != IndexType(0) {
			if key.header.IndexType() != key.expected.idx {
				t.Errorf("%v != %v", key.header.IndexType(), key.expected.idx)
			}
		}
	}
}
