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
	"fmt"
	"time"

	"github.com/cybergarage/go-redis/redis"
	"github.com/cybergarage/puzzledb-go/puzzledb/context"
)

func (service *Service) Set(conn *Conn, key string, val string, opt redis.SetOption) (*Message, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("Set")
	defer ctx.FinishSpan()
	now := time.Now()

	txn, err := service.TransactDatabase(ctx, conn, true)
	if err != nil {
		return nil, err
	}

	err = txn.SetObject(ctx, key, val)
	if err != nil {
		return nil, txn.CancelWithError(ctx, err)
	}

	err = txn.Commit(ctx)
	if err != nil {
		return nil, err
	}

	mSetLatency.Observe(float64(time.Since(now).Milliseconds()))

	return redis.NewOKMessage(), nil
}

func (service *Service) Get(conn *Conn, key string) (*Message, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("Get")
	defer ctx.FinishSpan()
	now := time.Now()

	txn, err := service.TransactDatabase(ctx, conn, false)
	if err != nil {
		return nil, err
	}

	obj, err := txn.GetObject(ctx, key)
	if err != nil {
		return nil, txn.CancelWithError(ctx, err)
	}

	err = txn.Commit(ctx)
	if err != nil {
		return nil, err
	}

	mGetLatency.Observe(float64(time.Since(now).Milliseconds()))

	switch v := obj.(type) {
	case string:
		return redis.NewBulkMessage(v), nil
	case []byte:
		return redis.NewBulkMessage(string(v)), nil
	default:
		return redis.NewBulkMessage(fmt.Sprintf("%v", v)), nil
	}
}
