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
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/cybergarage/go-logger/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	DefaultPrometheusPort                            = 9181
	DefaultPrometheusConnectionTimeout time.Duration = time.Second * 60
)

type PrometheusExporter struct {
	*Server
	httpServer *http.Server
	Addr       string
	Port       int
}

// NewPrometheusExporterWith returns a new PrometheusExporter.
func NewPrometheusExporterWith(server *Server) *PrometheusExporter {
	return &PrometheusExporter{
		Server:     server,
		httpServer: nil,
		Addr:       "",
		Port:       DefaultPrometheusPort,
	}
}

// SetPort sets a port number of the server.
func (server *PrometheusExporter) SetPort(port int) {
	server.Port = port
}

// EnabledConfig returns a port number for the specified query service name.
func (server *PrometheusExporter) EnabledConfig() (bool, error) {
	return server.Config.GetBool(metricsConfig, prometheusConfig, enabledConfig)
}

// PortConfig returns a port number for the specified query service name.
func (server *PrometheusExporter) PortConfig() (int, error) {
	return server.Config.GetInt(metricsConfig, prometheusConfig, portConfig)
}

// Start starts the server.
func (server *PrometheusExporter) Start() error {
	err := server.Stop()
	if err != nil {
		return err
	}

	addr := net.JoinHostPort(server.Addr, strconv.Itoa(server.Port))
	server.httpServer = &http.Server{ // nolint:exhaustruct
		Addr:        addr,
		ReadTimeout: DefaultPrometheusConnectionTimeout,
		Handler:     promhttp.Handler(),
	}

	c := make(chan error)
	go func() {
		c <- server.httpServer.ListenAndServe()
	}()

	select {
	case err = <-c:
	case <-time.After(time.Millisecond * 500):
		err = nil
	}

	log.Infof("prometheus exporter (%s) started", addr)

	return err
}

// Stop stops the Grpc server.
func (server *PrometheusExporter) Stop() error {
	if server.httpServer == nil {
		return nil
	}

	err := server.httpServer.Close()
	if err != nil {
		return err
	}

	addr := net.JoinHostPort(server.Addr, strconv.Itoa(server.Port))
	log.Infof("prometheus exporter (%s) terminated", addr)

	return nil
}
