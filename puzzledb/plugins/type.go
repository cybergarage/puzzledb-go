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

package plugins

// ServiceType represents a service type.
type ServiceType uint8

const (
	// DocumentService represents a serializer service.
	DocumentService ServiceType = iota
	// QueryService represents a query service.
	QueryService
	// DocumentStoreService represents a document store service.
	DocumentStoreService
	// KvStoreService represents a key-value store service.
	KvStoreService
	// CoordinatorService represents a coordinator service.
	CoordinatorService
	// ExtendService represents an uncategorized service.
	ExtendService
)

func (t ServiceType) String() string {
	switch t {
	case DocumentService:
		return "Document"
	case QueryService:
		return "Query"
	case DocumentStoreService:
		return "DocumentStore"
	case KvStoreService:
		return "KvStore"
	case CoordinatorService:
		return "Coordinator"
	case ExtendService:
		return "Extend"
	default:
		return "Unknown"
	}
}
