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

package actor

import (
	"github.com/cybergarage/puzzledb-go/puzzledb/coordinator"
)

type MessageBox struct {
	coordinator.Coordinator
}

// NewMessageBox returns a new actor service.
func NewMessageBox() *MessageBox {
	return NewMessageBoxWith(nil)
}

// NewServiceWith returns a new actor service with the specified coordinator.
func NewMessageBoxWith(coordinator coordinator.Coordinator) *MessageBox {
	return &MessageBox{
		Coordinator: coordinator,
	}
}

// SetCoordinator sets a coordinator.
func (mbox *MessageBox) SetCoordinator(c coordinator.Coordinator) { // nolint: stylecheck
	mbox.Coordinator = c
}
