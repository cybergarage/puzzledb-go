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
	"github.com/cybergarage/puzzledb-go/puzzledb/cluster"
)

// NewMessageScanKey returns a new scan message key to get the latest message clock.
func NewMessageScanKey() Key {
	return NewKeyWith(MessageObjectKeyHeader[:])
}

// NewMessageKeyWith returns a new message key with the specified message.
func NewMessageKeyWith(msg Message, clock cluster.Clock) Key {
	return NewKeyWith(MessageObjectKeyHeader[:], clock)
}
