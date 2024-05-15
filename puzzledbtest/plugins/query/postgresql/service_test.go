// Copyright (C) 2020 The PuzzleDB Authors.
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
	"fmt"
	"testing"
	"time"

	"github.com/cybergarage/go-postgresql/postgresql/auth"
	"github.com/cybergarage/go-postgresql/postgresqltest/client"
	"github.com/cybergarage/go-sqltest/sqltest"
	"github.com/cybergarage/puzzledb-go/puzzledbtest"
)

const testDBNamePrefix = "pgtest"

type ServerTestFunc = func(*testing.T, *puzzledbtest.Server, string)

func TestPostgreSQLServer(t *testing.T) {
	testFuncs := []struct {
		name string
		fn   ServerTestFunc
	}{
		{"authenticator", RunAuthenticatorTest},
		{"tls", RunTLSSessionTest},
		// {"copy", TestServerCopy},
	}

	server := puzzledbtest.NewServer()
	err := server.Start()
	if err != nil {
		t.Error(err)
		return
	}

	for _, testFunc := range testFuncs {
		testDBName := fmt.Sprintf("%s%d", testDBNamePrefix, time.Now().UnixNano())
		t.Run(testFunc.name, func(t *testing.T) {
			// Create a test database

			client := sqltest.NewPostgresClient()

			client.SetDatabase("postgres")
			err := client.Open()
			if err != nil {
				t.Error(err)
				return
			}

			err = client.CreateDatabase(testDBName)
			if err != nil {
				t.Error(err)
				return
			}

			err = client.Close()
			if err != nil {
				t.Error(err)
			}

			// Run tests

			testFunc.fn(t, server, testDBName)

			err = client.DropDatabase(testDBName)
			if err != nil {
				t.Error(err)
			}

			err = client.Close()
			if err != nil {
				t.Error(err)
			}
		})
	}

	err = server.Stop()
	if err != nil {
		t.Error(err)
		return
	}
}

// RunAuthenticatorTest tests the authenticators.
func RunAuthenticatorTest(t *testing.T, server *puzzledbtest.Server, testDBName string) {
	t.Helper()

	const (
		username = "testuser"
		password = "testpassword"
	)

	authenticators := []auth.Authenticator{
		auth.NewCleartextPasswordAuthenticatorWith(username, password),
	}

	for _, authenticator := range authenticators {
		server.AddAuthenticator(authenticator)

		client := client.NewDefaultClient()
		client.SetUser(username)
		client.SetPassword(password)
		client.SetDatabase(testDBName)
		err := client.Open()
		if err != nil {
			t.Error(err)
			return
		}

		err = client.Ping()
		if err != nil {
			t.Error(err)
		}

		err = client.Close()
		if err != nil {
			t.Error(err)
		}

		server.ClearAuthenticators()
	}
}

// RunTLSSessionTest tests the TLS session.
// PostgreSQL: Documentation: 16: 34.19. SSL Support
// https://www.postgresql.org/docs/current/libpq-ssl.html
// PostgreSQL: Documentation: 16: 19.9. Secure TCP/IP Connections with SSL
// https://www.postgresql.org/docs/current/ssl-tcp.html#SSL-CERTIFICATE-CREATION
func RunTLSSessionTest(t *testing.T, server *puzzledbtest.Server, testDBName string) {
	t.Helper()

	const (
		clientKey  = ".,/../certs/key.pem"
		clientCert = ".,/../certs/cert.pem"
		rootCert   = ".,/../certs/root_cert.pem"
	)

	client := client.NewDefaultClient()
	client.SetClientKeyFile(clientKey)
	client.SetClientCertFile(clientCert)
	client.SetRootCertFile(rootCert)

	err := client.Open()
	if err != nil {
		t.Error(err)
		return
	}

	err = client.Ping()
	if err != nil {
		t.Error(err)
	}

	err = client.Close()
	if err != nil {
		t.Error(err)
	}

	server.ClearAuthenticators()
}
