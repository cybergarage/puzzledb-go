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
	"github.com/cybergarage/go-sqlparser/sql/net"
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
)

// Conn represents a SQL connection.
type Conn = net.Conn

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

// SetTransaction sets the transaction.
func (connMap ConnectionMap) SetTransaction(conn Conn, db store.Database, txn store.Transaction) error {
	dbName := db.Name()
	if _, hasDB := connMap[dbName]; !hasDB {
		connMap[dbName] = NewDatabaseMap()
	}
	connMap[dbName][conn.UUID().String()] = &Database{
		Transaction: txn,
	}
	return nil
}

// GetTransaction returns the transaction.
func (connMap ConnectionMap) GetTransaction(conn Conn, db store.Database) (store.Transaction, error) {
	dbName := db.Name()
	dbMap, hasTxn := connMap[dbName]
	if !hasTxn {
		return nil, newErrDatabaseNotExist(dbName)
	}
	dbTxn, hasTxn := dbMap[conn.UUID().String()]
	if !hasTxn {
		return nil, newErrConnectionExist(conn.UUID().String())
	}
	return dbTxn.Transaction, nil
}

// RemoveTransaction removes the transaction.
func (connMap ConnectionMap) RemoveTransaction(conn Conn, db store.Database) error {
	dbName := db.Name()
	dbMap, hasTxn := connMap[dbName]
	if !hasTxn {
		return nil
	}
	delete(dbMap, conn.UUID().String())
	if len(dbMap) == 0 {
		delete(connMap, dbName)
	}
	return nil
}
