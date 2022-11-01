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

package query

// ExecutionError represents a query execution error.
type ExecutionError interface {
	// GetErrorCode returns an error code if the execution is failed, otherwise zero.
	GetErrorCode() int
	// GetErrorMessage returns an error message if the execution is failed, otherwise a zero string.
	GetErrorMessage() string
}

// ExecutionResult represents a query execution result.
type ExecutionResult interface {
	// IsSuccess returns true when the execution is succeed, otherwise false.
	IsSuccess() bool
	// NModified returns an affected object count executing the query.
	NModified() int64
}

// Result represents a result set which includes query execution status.
type Result interface {
	ExecutionError
	ExecutionResult
	String() string
}
