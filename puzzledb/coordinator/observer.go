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

// EventType represents a coordinator event type.
type EventType uint8

// Event represents a  coordinator event.
type Event interface {
	// Type returns the event type.
	Type() EventType
	// Object returns the object of the event.
	Object() Object
}

// Observer represents a coordinator observer.
type Observer interface {
	// ProcessEvent processes the event.
	ProcessEvent(evt Event)
}