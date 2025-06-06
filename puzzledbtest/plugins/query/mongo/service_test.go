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

package mongo

import (
	"testing"

	"github.com/cybergarage/go-mongo/mongo/shell"
	"github.com/cybergarage/go-mongo/mongotest"
	"github.com/cybergarage/puzzledb-go/puzzledbtest"
)

const (
	TestClientCertFile = "../../../certs/client.pem"
	TestClientCAFile   = "../../../certs/ca.pem"
)

func TestMongoService(t *testing.T) {
	server := puzzledbtest.NewServer()
	err := server.Start()
	if err != nil {
		t.Error(err)
		return
	}

	defer func() {
		err := server.Stop()
		if err != nil {
			t.Error(err)
			return
		}
	}()

	client := shell.NewClient()

	if server.IsTLSEnabled() {
		client.SetTLSEnabled(true)
		client.SetTLSCertificateKeyFile(TestClientCertFile)
		client.SetTLSCAFile(TestClientCAFile)
	}

	err = client.Open()
	if err != nil {
		t.Skip(err.Error())
		return
	}

	defer func() {
		err := client.Close()
		if err != nil {
			t.Error(err)
		}
	}()

	mongotest.RunEmbedSuite(t, client)
}
