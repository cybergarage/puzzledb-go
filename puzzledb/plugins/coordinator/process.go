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
	"github.com/cybergarage/puzzledb-go/puzzledb/coordinator"
	"github.com/google/uuid"
)

// ProcessObject represents a store process state object.
type ProcessObject struct {
	ID    uuid.UUID
	Host  string
	Clock uint64
}

// NewProcessWith returns a new process with the specified process object.
func NewProcessWith(obj *ProcessObject) coordinator.Process {
	process := coordinator.NewProcess()
	process.SetID(obj.ID)
	process.SetHost(obj.Host)
	process.SetClock(obj.Clock)
	return process
}

func NewProcessScanKey() coordinator.Key {
	return coordinator.NewKeyWith(coordinator.StateObjectKeyHeader[:])
}

// NewProcessObjectWith returns a new process object with the specified process.
func NewProcessObjectWith(process coordinator.Process) (*ProcessObject, error) {
	return &ProcessObject{
		ID:    process.ID(),
		Host:  process.Host(),
		Clock: uint64(process.Clock()),
	}, nil
}
