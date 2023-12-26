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
	"time"

	"github.com/cybergarage/go-redis/redis"
	"github.com/cybergarage/puzzledb-go/puzzledb/context"
)

func (service *Service) HDel(conn *Conn, key string, fields []string) (*Message, error) {
	return nil, newErrNotSupported("HDel")
}

func (service *Service) HSet(conn *Conn, key string, field string, val string, opt redis.HSetOption) (*Message, error) {
	return nil, newErrNotSupported("HSet")
}

func (service *Service) HGet(conn *Conn, key string, field string) (*Message, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("HGet")
	defer ctx.FinishSpan()
	now := time.Now()

	txn, err := service.TransactDatabase(ctx, conn, false)
	if err != nil {
		return nil, err
	}

	obj, err := txn.GetKeyHashObject(ctx, key)
	if err != nil {
		return nil, txn.CancelWithError(ctx, err)
	}

	err = txn.Commit(ctx)
	if err != nil {
		return nil, err
	}

	if obj == nil {
		return redis.NewNilMessage(), nil
	}

	val, ok := obj[field]
	if !ok {
		return redis.NewNilMessage(), nil
	}

	mHGetLatency.Observe(float64(time.Since(now).Milliseconds()))

	return redis.NewStringMessage(val), nil
}

func (service *Service) HGetAll(conn *Conn, key string) (*Message, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("HGetAll")
	defer ctx.FinishSpan()
	now := time.Now()

	txn, err := service.TransactDatabase(ctx, conn, false)
	if err != nil {
		return nil, err
	}

	obj, err := txn.GetKeyHashObject(ctx, key)
	if err != nil {
		return nil, txn.CancelWithError(ctx, err)
	}

	err = txn.Commit(ctx)
	if err != nil {
		return nil, err
	}

	arrayMsg := redis.NewArrayMessage()
	array, _ := arrayMsg.Array()
	for key, val := range obj {
		array.Append(redis.NewBulkMessage(key))
		array.Append(redis.NewBulkMessage(val))
	}

	mHGetAllLatency.Observe(float64(time.Since(now).Milliseconds()))

	return arrayMsg, nil
}
