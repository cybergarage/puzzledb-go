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

package sql

import (
	"github.com/cybergarage/puzzledb-go/puzzledb/context"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/query"
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
)

// Service represents a new MySQL service instance.
type Service struct {
	*query.BaseService
}

// NewService returns a new MySQL service.
func NewService() *Service {
	service := &Service{
		BaseService: query.NewBaseService(),
	}
	return service
}

// CancelTransactionWithError cancels the specified transaction with the specified error.
func (service *Service) CancelTransactionWithError(ctx context.Context, txn store.Transaction, err error) error {
	if txErr := txn.Cancel(ctx); txErr != nil {
		return txErr
	}
	return err
}
