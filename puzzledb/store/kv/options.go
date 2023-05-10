// DocumentOperation represents a document operation. type DocumentOperation interface { // InsertDocument puts a document object with the specified primary key. InsertDocument(ctx context.Context, docKey Key, obj Object) error // FindDocuments returns a result set matching the specified key. FindDocuments(ctx context.Context, docKey Key) (ResultSet, error) // UpdateDocument updates a document object with the specified primary key. UpdateDocument(ctx context.Context, docKey Key, obj Object) error // RemoveDocument removes a document object with the specified primary key. RemoveDocument(ctx context.Context, docKey Key) error // RemoveDocuments removes document objects with the specified primary key. RemoveDocuments(ctx context.Context, docKey Key) error // TruncateDocuments removes all document objects. TruncateDocuments(ctx context.Context) error }// Copyright (C) 2022 The PuzzleDB Authors.
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

package kv

// Option represents a option.
type Option = any
