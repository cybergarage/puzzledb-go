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

type Conn = redis.Conn
type Message = redis.Message

func (service *Service) Del(conn *Conn, keys []string) (*Message, error) {
	now := time.Now()
	mDelLatency.Observe(float64(time.Since(now).Milliseconds()))
	return nil, newErrNotSupported("Del")
}

func (service *Service) Exists(conn *Conn, keys []string) (*Message, error) {
	return nil, newErrNotSupported("Exists")
}

func (service *Service) Expire(conn *Conn, key string, opt redis.ExpireOption) (*Message, error) {
	return nil, newErrNotSupported("Expire")
}

func (service *Service) Keys(conn *Conn, pattern string) (*Message, error) {
	return nil, newErrNotSupported("Keys")
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

func (service *Service) Set(conn *Conn, key string, val string, opt redis.SetOption) (*Message, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("Set")
	defer ctx.FinishSpan()
	now := time.Now()

	db, err := service.GetDatabase(ctx, conn.Database())
	if err != nil {
		return nil, err
	}

	txn, err := db.Transact(true)
	if err != nil {
		return nil, err
	}
	err = txn.InsertDocument(ctx, []any{key}, val)
	if err != nil {
		err = txn.Cancel(ctx)
		if err != nil {
			return nil, err
		}
		return nil, err
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

	db, err := service.GetDatabase(ctx, conn.Database())
	if err != nil {
		return nil, err
	}

	txn, err := db.Transact(false)
	if err != nil {
		return nil, err
	}
	rs, err := txn.FindDocuments(ctx, []any{key})
	if err != nil {
		err = txn.Cancel(ctx)
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	objs := rs.Objects()
	if err != nil || len(objs) != 1 {
		err = txn.Cancel(ctx)
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	err = txn.Commit(ctx)
	if err != nil {
		return nil, err
	}

	mGetLatency.Observe(float64(time.Since(now).Milliseconds()))

	obj := objs[0]
	switch v := obj.(type) {
	case string:
		return redis.NewBulkMessage(v), nil
	case []byte:
		return redis.NewBulkMessage(string(v)), nil
	default:
		return redis.NewBulkMessage(fmt.Sprintf("%v", v)), nil
	}
}

func (service *Service) MSet(conn *Conn, dict map[string]string, opt redis.MSetOption) (*Message, error) {
	return nil, newErrNotSupported("MSet")
}

func (service *Service) MGet(conn *Conn, keys []string) (*Message, error) {
	return nil, newErrNotSupported("MGet")
}

func (service *Service) HDel(conn *Conn, key string, fields []string) (*Message, error) {
	return nil, newErrNotSupported("HDel")
}

func (service *Service) HSet(conn *Conn, key string, field string, val string, opt redis.HSetOption) (*Message, error) {
	return nil, newErrNotSupported("HSet")
}

func (service *Service) HGet(conn *Conn, key string, field string) (*Message, error) {
	return nil, newErrNotSupported("HGet")
}

func (service *Service) HGetAll(conn *Conn, key string) (*Message, error) {
	return nil, newErrNotSupported("HGetAll")
}

func (service *Service) HMSet(conn *Conn, key string, dict map[string]string) (*Message, error) {
	return nil, newErrNotSupported("HMSet")
}

func (service *Service) HMGet(conn *Conn, key string, fields []string) (*Message, error) {
	return nil, newErrNotSupported("HMGet")
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

func (service *Service) ZRangeByScore(conn *Conn, key string, min float64, max float64, opt redis.ZRangeOption) (*Message, error) {
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
