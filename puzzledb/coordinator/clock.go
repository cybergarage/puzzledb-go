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

import "math"

// Clock represents a logical clock.
type Clock uint64

// CompareClocks compares two clocks. If c1 and c2 are equal, returns 0.
// If c1 is greater than c2, returns 1. // If c2 is greater than c1, returns -1.
// However, if the difference between c1 and c2 is greater than half of the maximum UNIT64 value,
// the clock is considered to be back to zero and the comparison results are reversed.
func (c1 Clock) Compare(c2 Clock) int {
	if c1 == c2 {
		return 0
	}

	if c1 > c2 {
		if c1-c2 > math.MaxUint64/2 {
			return -1
		} else {
			return 1
		}
	}

	if c2-c1 > math.MaxUint64/2 {
		return 1
	} else {
		return -1
	}
}