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

package redis

import (
	"strconv"

	"github.com/cybergarage/go-redis/redis"
	"github.com/cybergarage/puzzledb-go/puzzledb/context"
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
)

// GetDatabase returns the database with the specified ID.
func (service *Service) LookupDatabase(ctx context.Context, dbid redis.DatabaseID) (store.Database, error) {
	store := service.Store()
	name := strconv.Itoa(dbid)
	db, err := store.LookupDatabase(ctx, name)
	if err == nil {
		return db, nil
	}
	err = store.CreateDatabase(ctx, name)
	if err != nil {
		return nil, err
	}
	return store.LookupDatabase(ctx, name)
}

// TransactDatabase returns a transaction for the database with the specified ID.
func (service *Service) TransactDatabase(ctx context.Context, conn *Conn, write bool) (*Transaction, error) {
	dbid := conn.Database()
	db, err := service.LookupDatabase(ctx, dbid)
	if err != nil {
		return nil, err
	}

	txn, err := db.Transact(write)
	if err != nil {
		return nil, err
	}

	return &Transaction{
		Transaction: txn,
		DatabaseID:  dbid,
	}, nil
}
