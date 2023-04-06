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
	// EncoderDocumentService represents a serializer service for document.
	EncoderDocumentService ServiceType = iota
	// EncoderKeyService represents a serializer service for key.
	EncoderKeyService
	// QueryService represents a query service.
	QueryService
	// StoreDocumentService represents a document store service.
	StoreDocumentService
	// StoreKvService represents a key-value store service.
	StoreKvService
	// CoordinatorService represents a coordinator service.
	CoordinatorService
	// ExtendService represents an uncategorized service.
	ExtendService
)

// ServiceTypes returns all service types.
func ServiceTypes() []ServiceType {
	return []ServiceType{
		EncoderDocumentService,
		EncoderKeyService,
		QueryService,
		StoreDocumentService,
		StoreKvService,
		CoordinatorService,
		ExtendService,
	}
}

// String returns a string representation of the service type.
func (t ServiceType) String() string {
	switch t {
	case EncoderDocumentService:
		return "encorder.document"
	case EncoderKeyService:
		return "encorder.key"
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
