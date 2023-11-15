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

package puzzledb

import (
	"github.com/cybergarage/puzzledb-go/puzzledb/auth"
)

// Authenticator represents an authenticator.
type Authenticator struct {
	handlers []auth.AuthHandler
}

// NewAuthenticator returns a new authenticator.
func NewAuthenticator() *Authenticator {
	return &Authenticator{
		handlers: []auth.AuthHandler{},
	}
}

// AddHandler adds a handler.
func (authenticator *Authenticator) AddHandler(handler auth.AuthHandler) {
	authenticator.handlers = append(authenticator.handlers, handler)
}
