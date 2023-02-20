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

	"github.com/cybergarage/go-redis/redis"
)

type DBContext = redis.DBContext
type Message = redis.Message

func (service *Service) Del(ctx *DBContext, keys []string) (*Message, error) {
	return nil, nil
}

func (service *Service) Exists(ctx *DBContext, keys []string) (*Message, error) {
	return nil, nil
}

func (service *Service) Expire(ctx *DBContext, key string, opt redis.ExpireOption) (*Message, error) {
	return nil, nil
}

func (service *Service) Keys(ctx *DBContext, pattern string) (*Message, error) {
	return nil, nil
}

func (service *Service) Rename(ctx *DBContext, key string, newkey string, opt redis.RenameOption) (*Message, error) {
	return nil, nil
}

func (service *Service) Type(ctx *DBContext, key string) (*Message, error) {
	return nil, nil
}

func (service *Service) TTL(ctx *DBContext, key string) (*Message, error) {
	return nil, nil
}

func (service *Service) Set(ctx *DBContext, key string, val string, opt redis.SetOption) (*Message, error) {
	db, err := service.GetDatabase(ctx.ID())
	if err != nil {
		return nil, err
	}

	tx, err := db.Transact(true)
	if err != nil {
		return nil, err
	}
	err = tx.InsertDocument([]any{key}, val)
	if err != nil {
		err = tx.Cancel()
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return redis.NewOKMessage(), nil
}

func (service *Service) Get(ctx *DBContext, key string) (*Message, error) {
	db, err := service.GetDatabase(ctx.ID())
	if err != nil {
		return nil, err
	}

	tx, err := db.Transact(false)
	if err != nil {
		return nil, err
	}
	rs, err := tx.FindDocuments([]any{key})
	if err != nil {
		err = tx.Cancel()
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	objs := rs.Objects()
	if err != nil || len(objs) != 1 {
		err = tx.Cancel()
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

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

func (service *Service) MSet(ctx *DBContext, dict map[string]string, opt redis.MSetOption) (*Message, error) {
	return nil, nil
}

func (service *Service) MGet(ctx *DBContext, keys []string) (*Message, error) {
	return nil, nil
}

func (service *Service) HDel(ctx *DBContext, key string, fields []string) (*Message, error) {
	return nil, nil
}

func (service *Service) HSet(ctx *DBContext, key string, field string, val string, opt redis.HSetOption) (*Message, error) {
	return nil, nil
}

func (service *Service) HGet(ctx *DBContext, key string, field string) (*Message, error) {
	return nil, nil
}

func (service *Service) HGetAll(ctx *DBContext, key string) (*Message, error) {
	return nil, nil
}

func (service *Service) HMSet(ctx *DBContext, key string, dict map[string]string) (*Message, error) {
	return nil, nil
}

func (service *Service) HMGet(ctx *DBContext, key string, fields []string) (*Message, error) {
	return nil, nil
}

func (service *Service) LPush(ctx *DBContext, key string, elements []string, opt redis.PushOption) (*Message, error) {
	return nil, nil
}

func (service *Service) RPush(ctx *DBContext, key string, elements []string, opt redis.PushOption) (*Message, error) {
	return nil, nil
}

func (service *Service) LPop(ctx *DBContext, key string, count int) (*Message, error) {
	return nil, nil
}

func (service *Service) RPop(ctx *DBContext, key string, count int) (*Message, error) {
	return nil, nil
}

func (service *Service) LRange(ctx *DBContext, key string, start int, stop int) (*Message, error) {
	return nil, nil
}

func (service *Service) LIndex(ctx *DBContext, key string, index int) (*Message, error) {
	return nil, nil
}

func (service *Service) LLen(ctx *DBContext, key string) (*Message, error) {
	return nil, nil
}

func (service *Service) SAdd(ctx *DBContext, key string, members []string) (*Message, error) {
	return nil, nil
}

func (service *Service) SMembers(ctx *DBContext, key string) (*Message, error) {
	return nil, nil
}

func (service *Service) SRem(ctx *DBContext, key string, members []string) (*Message, error) {
	return nil, nil
}

func (service *Service) ZAdd(ctx *DBContext, key string, members []*redis.ZSetMember, opt redis.ZAddOption) (*Message, error) {
	return nil, nil
}

func (service *Service) ZRange(ctx *DBContext, key string, start int, stop int, opt redis.ZRangeOption) (*Message, error) {
	return nil, nil
}

func (service *Service) ZRangeByScore(ctx *DBContext, key string, min float64, max float64, opt redis.ZRangeOption) (*Message, error) {
	return nil, nil
}

func (service *Service) ZRem(ctx *DBContext, key string, members []string) (*Message, error) {
	return nil, nil
}

func (service *Service) ZScore(ctx *DBContext, key string, member string) (*Message, error) {
	return nil, nil
}

func (service *Service) ZIncBy(ctx *DBContext, key string, inc float64, member string) (*Message, error) {
	return nil, nil
}
