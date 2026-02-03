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
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins"
)

// BaseStore represents a base store.
type BaseStore struct {
	plugins.Config
	document.KeyCoder
}

// NewBaseStoreWith returns a new base store instance.
func NewBaseStoreWith(coder document.KeyCoder) *BaseStore {
	return &BaseStore{
		Config:   plugins.NewConfig(),
		KeyCoder: coder,
	}
}

// NewBaseStore returns a new base store instance.
func NewBaseStore() *BaseStore {
	return NewBaseStoreWith(nil)
}

// ServiceType returns the plug-in service type.
func (store *BaseStore) ServiceType() plugins.ServiceType {
	return plugins.StoreKvService
}

// SetKeyCoder sets the key coder.
func (store *BaseStore) SetKeyCoder(coder document.KeyCoder) {
	store.KeyCoder = coder
}

// DecodeKey returns the decoded key from the specified bytes if available, otherwise returns an error.
func (store *BaseStore) DecodeKey(b []byte) (document.Key, error) {
	return store.KeyCoder.DecodeKey(b)
}

// EncodeKey returns the encoded bytes from the specified key if available, otherwise returns an error.
func (store *BaseStore) EncodeKey(key document.Key) ([]byte, error) {
	return store.KeyCoder.EncodeKey(key)
}
