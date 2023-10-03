// Copyright (C) 2019 The PuzzleDB Authors.
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

package sql

import (
	"time"

	"github.com/cybergarage/go-tracing/tracer"
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
	"github.com/google/uuid"
)

// Conn represents a SQL connection.
type Conn interface {
	// Database returns the database.
	Database() string
	// Timestamp returns the timestamp.
	Timestamp() time.Time
	// UUID returns the UUID.
	UUID() uuid.UUID
	// SpanContext returns the span context.
	SpanContext() tracer.Context
}

// Database represents a database.
type Database struct {
	store.Transaction
}

// DatabaseMap represents a database map.
type DatabaseMap map[string]*Database

// NewDatabase returns a new database.
func NewDatabaseMap() DatabaseMap {
	return make(DatabaseMap)
}

// ConnectionMap represents a connection map.
type ConnectionMap map[string]DatabaseMap

// NewConnection returns a new connection.
func NewConnectionMap() ConnectionMap {
	return make(ConnectionMap)
}
