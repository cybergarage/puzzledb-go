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

package puzzledb

import (
	_ "embed"

	"github.com/cybergarage/puzzledb-go/puzzledb/config"
)

const (
	ConfigAPI     = "api"
	ConfigPlugins = "plugins"
	ConfigGrpc    = "grpc"
	ConfigQuery   = "query"
	ConfigPort    = "port"
	ConfigEnabled = "enabled"
	ConfigTracer  = "tracer"
	ConfigPprof   = "pprof"
	ConfigLogger  = "logger"
	ConfigDefault = "default"
	ConfigLevel   = "level"
	ConfigAuth    = "auth"
	ConfigTLS     = "tls"
)

// Config represents a configuration interface for PuzzleDB.
type Config interface {
	config.Config
}
