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
)

// MessageQueue is a queue for coordinator message.
type MessageQueue struct {
	messages []coordinator.Message
}

// NewMessageQueue returns a new messageQueue.
func NewMessageQueue() *MessageQueue {
	return &MessageQueue{
		messages: make([]coordinator.Message, 0),
	}
}

// EnqueueMessage appends a message to the queue.
func (q *MessageQueue) EnqueueMessage(msg coordinator.Message) {
	q.messages = append(q.messages, msg)
}

// PushMessage pushes a message to the queue.
func (q *MessageQueue) PushMessage(msg coordinator.Message) {
	q.messages = append([]coordinator.Message{msg}, q.messages...)
}

// PopMessage removes a message from the queue.
func (q *MessageQueue) PopMessage() (coordinator.Message, error) {
	if len(q.messages) == 0 {
		return nil, coordinator.ErrNoMessage
	}
	msg := q.messages[0]
	q.messages = q.messages[1:]
	return msg, nil
}
