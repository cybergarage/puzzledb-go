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

// ProcessStatus represents a coordinator process status.
type ProcessStatus int

const (
	// ProcessUnknown represents an unknown process status.
	ProcessUnknown ProcessStatus = iota
	// ProcessIdle represents an idle process status.
	ProcessIdle
	// ProcessStarting represents a starting process status.
	ProcessStarting
	// ProcessRunning represents a running process status.
	ProcessRunning
	// ProcessStopping represents a stopping process status.
	ProcessStopping
	// ProcessStopped represents a stopped process status.
	ProcessStopped
	// ProcessAborted represents an aborted process status.
	ProcessAborted
)

var processStatuses = map[ProcessStatus]string{
	ProcessUnknown:  "unknown",
	ProcessIdle:     "idle",
	ProcessStarting: "starting",
	ProcessRunning:  "running",
	ProcessStopping: "stopping",
	ProcessStopped:  "stopped",
	ProcessAborted:  "aborted",
}

// NewProcessStatusWith returns a new process status with the specified string.
func NewProcessStatusWith(s string) ProcessStatus {
	for status, statusString := range processStatuses {
		if statusString == s {
			return status
		}
	}
	return ProcessUnknown
}

// State represents a coordinator state.
func (t ProcessStatus) String() string {
	s, ok := processStatuses[t]
	if ok {
		return s
	}
	return processStatuses[ProcessUnknown]
}
