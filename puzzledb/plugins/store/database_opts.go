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

package store

import (
	"fmt"

	"github.com/cybergarage/puzzledb-go/puzzledb/store"
)

func newDatabaeOptionsFrom(obj any) (store.DatabaseOptions, error) {
	opts, ok := obj.(store.DatabaseOptions)
	if ok {
		return opts, nil
	}

	optMap, ok := obj.(map[any]any)
	if ok {
		opts := store.DatabaseOptions{}
		for k, v := range optMap {
			opts[fmt.Sprintf("%v", k)] = v
		}
		return opts, nil
	}

	return nil, store.NewDatabaseOptionsInvalidError(obj)
}
