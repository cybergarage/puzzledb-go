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

package coordinator

import (
	"errors"
	"fmt"
)

var (
	ErrInvalid = errors.New("invalid")
)

func newKeyInvalidError(v any) error {
	return fmt.Errorf("key type (%s:%T) is %w", v, v, ErrInvalid)
}

func newValueInvalidError(v any) error {
	return fmt.Errorf("value type (%s:%T) is %w", v, v, ErrInvalid)
}
