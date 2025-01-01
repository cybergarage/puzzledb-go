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

package puzzledbtest

import (
	"testing"

	"github.com/cybergarage/puzzledb-go/puzzledb"
	"github.com/cybergarage/puzzledb-go/puzzledb/tls"
)

func TestConfigs(t *testing.T) {
	paths := []string{".", "../puzzledb/conf"}
	for _, path := range paths {
		t.Run(path, func(t *testing.T) {
			conf, err := puzzledb.NewConfigWithPath(path)
			if err != nil {
				t.Error(err)
				return
			}

			ports := []struct {
				name     string
				expected int
			}{
				{
					name:     "mysql",
					expected: 3306,
				},
			}
			for _, port := range ports {
				portNum, err := conf.GetConfigInt(puzzledb.ConfigPlugins, puzzledb.ConfigQuery, port.name, puzzledb.ConfigPort)
				if err != nil {
					t.Error(err)
					t.Log(conf.String())
					return
				}
				if portNum != port.expected {
					t.Errorf("expected port number is %d but got %d", port.expected, portNum)
					return
				}
			}
		})
	}
}

func TestDefaultTestConfig(t *testing.T) {
	configs := []Config{
		NewConfig(),
	}
	paths := []string{"."}
	for _, path := range paths {
		config, err := puzzledb.NewConfigWithPath(path)
		if err != nil {
			t.Error(err)
			return
		}
		configs = append(configs, config)
	}

	for _, config := range configs {
		t.Run(config.UsedConfigFile(), func(t *testing.T) {
			DefaultTestConfigTest(t, config)
		})
	}
}

func DefaultTestConfigTest(t *testing.T, config Config) {
	t.Helper()

	// Check tls config

	tlsConf, err := tls.NewConfigWith(config, puzzledb.ConfigTLS)
	if err != nil {
		t.Error(err)
		return
	}

	_, err = tlsConf.TLSConfig()
	if err != nil {
		t.Error(err)
		return
	}
}
