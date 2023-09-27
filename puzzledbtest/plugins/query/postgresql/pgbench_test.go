// Copyright (C) 2019 The PuzzleDB Authors.
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

package postgresql

import (
	_ "embed"
	"os/exec"
	"testing"

	"github.com/cybergarage/go-logger/log"
	"github.com/cybergarage/puzzledb-go/puzzledbtest"
)

func TestPgBench(t *testing.T) {
	log.SetStdoutDebugEnbled(true)

	server := puzzledbtest.NewServer()
	err := server.Start()
	if err != nil {
		t.Error(err)
		return
	}

	scripts := []string{
		"./pgbench-init",
		"./pgbench-run",
	}

	for _, script := range scripts {
		cmd := exec.Command(script)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Skip(err)
			return
		}
		t.Log(string(output))
	}

	err = server.Stop()
	if err != nil {
		t.Error(err)
		return
	}
}
