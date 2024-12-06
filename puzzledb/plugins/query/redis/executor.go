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

type Conn = redis.Conn
type Message = redis.Message

func (service *Service) Del(conn *Conn, keys []string) (*Message, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("Del")
	defer ctx.FinishSpan()
	now := time.Now()

	txn, err := service.TransactDatabase(ctx, conn, true)
	if err != nil {
		return nil, err
	}

	removedCount := 0
	for _, key := range keys {
		docKey := NewDocumentKeyWith(txn.DatabaseID, key)
		err := txn.RemoveObject(ctx, docKey)
		if err == nil {
			removedCount++
		}
	}

	err = txn.Commit(ctx)
	if err != nil {
		return nil, err
	}

	mDelLatency.Observe(float64(time.Since(now).Milliseconds()))

	return redis.NewIntegerMessage(removedCount), nil
}

func (service *Service) Exists(conn *Conn, keys []string) (*Message, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("Exists")
	defer ctx.FinishSpan()

	txn, err := service.TransactDatabase(ctx, conn, false)
	if err != nil {
		return nil, err
	}

	existCount := 0
	for _, key := range keys {
		docKey := NewDocumentKeyWith(txn.DatabaseID, key)
		rs, err := txn.FindObjects(ctx, docKey)
		if err != nil {
			return nil, err
		}
		if rs.Next() {
			existCount++
		}
	}

	err = txn.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return redis.NewIntegerMessage(existCount), nil
}

func (service *Service) Expire(conn *Conn, key string, opt redis.ExpireOption) (*Message, error) {
	return nil, newErrNotSupported("Expire")
}

func (service *Service) Keys(conn *Conn, pattern string) (*Message, error) {
	return nil, newErrNotSupported("Keys")
}

func (service *Service) Scan(conn *redis.Conn, cursor int, opt redis.ScanOption) (*redis.Message, error) {
	return nil, newErrNotSupported("Scan")
}

func (service *Service) Rename(conn *Conn, key string, newkey string, opt redis.RenameOption) (*Message, error) {
	return nil, newErrNotSupported("Rename")
}

func (service *Service) Type(conn *Conn, key string) (*Message, error) {
	return nil, newErrNotSupported("Type")
}

func (service *Service) TTL(conn *Conn, key string) (*Message, error) {
	return nil, newErrNotSupported("TTL")
}

func (service *Service) LPush(conn *Conn, key string, elements []string, opt redis.PushOption) (*Message, error) {
	return nil, newErrNotSupported("LPush")
}

func (service *Service) RPush(conn *Conn, key string, elements []string, opt redis.PushOption) (*Message, error) {
	return nil, newErrNotSupported("RPush")
}

func (service *Service) LPop(conn *Conn, key string, count int) (*Message, error) {
	return nil, newErrNotSupported("LPop")
}

func (service *Service) RPop(conn *Conn, key string, count int) (*Message, error) {
	return nil, newErrNotSupported("RPop")
}

func (service *Service) LRange(conn *Conn, key string, start int, stop int) (*Message, error) {
	return nil, newErrNotSupported("LRange")
}

func (service *Service) LIndex(conn *Conn, key string, index int) (*Message, error) {
	return nil, newErrNotSupported("LIndex")
}

func (service *Service) LLen(conn *Conn, key string) (*Message, error) {
	return nil, newErrNotSupported("LLen")
}

func (service *Service) SAdd(conn *Conn, key string, members []string) (*Message, error) {
	return nil, newErrNotSupported("SAdd")
}

func (service *Service) SMembers(conn *Conn, key string) (*Message, error) {
	return nil, newErrNotSupported("SMembers")
}

func (service *Service) SRem(conn *Conn, key string, members []string) (*Message, error) {
	return nil, newErrNotSupported("SRem")
}

func (service *Service) ZAdd(conn *Conn, key string, members []*redis.ZSetMember, opt redis.ZAddOption) (*Message, error) {
	return nil, newErrNotSupported("ZAdd")
}

func (service *Service) ZRange(conn *Conn, key string, start int, stop int, opt redis.ZRangeOption) (*Message, error) {
	return nil, newErrNotSupported("ZRange")
}

func (service *Service) ZRangeByScore(conn *Conn, key string, zmin float64, zmax float64, opt redis.ZRangeOption) (*Message, error) {
	return nil, newErrNotSupported("ZRangeByScore")
}

func (service *Service) ZRem(conn *Conn, key string, members []string) (*Message, error) {
	return nil, newErrNotSupported("ZRem")
}

func (service *Service) ZScore(conn *Conn, key string, member string) (*Message, error) {
	return nil, newErrNotSupported("ZScore")
}

func (service *Service) ZIncBy(conn *Conn, key string, inc float64, member string) (*Message, error) {
	return nil, newErrNotSupported("ZIncBy")
}
