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
	// ExclusiveServiceType represents an exclusive service type.
	ExclusiveServiceType = 0x80
)

const (
	// CoderDocumentService represents a serializer service for document.
	CoderDocumentService ServiceType = 0x01 | ExclusiveServiceType
	// CoderKeyService represents a serializer service for key.
	CoderKeyService ServiceType = 0x02 | ExclusiveServiceType
	// QueryService represents a query service.
	QueryService ServiceType = 0x03
	// StoreDocumentService represents a document store service.
	StoreDocumentService ServiceType = 0x04 | ExclusiveServiceType
	// StoreKvService represents a key-value store service.
	StoreKvService ServiceType = 0x05 | ExclusiveServiceType
	// CoordinatorService represents a coordinator service.
	CoordinatorService ServiceType = 0x06 | ExclusiveServiceType
	// ExtendService represents an uncategorized service.
	ExtendService ServiceType = 0x0F
)

// ServiceTypes returns all service types.
func ServiceTypes() []ServiceType {
	return []ServiceType{
		CoderDocumentService,
		CoderKeyService,
		QueryService,
		StoreDocumentService,
		StoreKvService,
		CoordinatorService,
		ExtendService,
	}
}

// IsExclusive returns true if the service type is exclusive.
func (t ServiceType) IsExclusive() bool {
	return (t & ExclusiveServiceType) != 0
}

// String returns a string representation of the service type.
func (t ServiceType) String() string {
	switch t {
	case CoderDocumentService:
		return "corder.document"
	case CoderKeyService:
		return "corder.key"
	case QueryService:
		return "query"
	case StoreDocumentService:
		return "store.document"
	case StoreKvService:
		return "store.key-value"
	case CoordinatorService:
		return "coordinator"
	case ExtendService:
		return "extend"
	default:
		return "unknown"
	}
}
