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

package prometheus

import (
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/cybergarage/go-logger/log"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	DefaultPrometheusPort                            = 9181
	DefaultPrometheusConnectionTimeout time.Duration = time.Second * 60
)

// PrometheusExporter is a prometheus exporter service.
type PrometheusExporter struct {
	*metrics.BaseService
	httpServer *http.Server
	Addr       string
	Port       int
}

// NewService returns a new prometheus exporter service.
func NewService() *PrometheusExporter {
	return &PrometheusExporter{
		BaseService: metrics.NewBaseService(),
		httpServer:  nil,
		Addr:        "",
		Port:        DefaultPrometheusPort,
	}
}

// ServiceName returns the plug-in service name.
func (exp *PrometheusExporter) ServiceName() string {
	return "prometheus"
}

// SetPort sets a port number of the exp.
func (exp *PrometheusExporter) SetPort(port int) {
	exp.Port = port
}

// Start starts the exp.
func (exp *PrometheusExporter) Start() error {
	err := exp.Stop()
	if err != nil {
		return err
	}

	port, err := exp.LookupServiceConfigPort(exp)
	if err == nil {
		exp.SetPort(port)
	}

	addr := net.JoinHostPort(exp.Addr, strconv.Itoa(exp.Port))
	exp.httpServer = &http.Server{ // nolint:exhaustruct
		Addr:        addr,
		ReadTimeout: DefaultPrometheusConnectionTimeout,
		Handler:     promhttp.Handler(),
	}

	c := make(chan error)
	go func() {
		c <- exp.httpServer.ListenAndServe()
	}()

	select {
	case err = <-c:
	case <-time.After(time.Millisecond * 500):
		err = nil
	}

	log.Infof("prometheus exporter (%s) started", addr)

	return err
}

// Stop stops the Grpc exp.
func (exp *PrometheusExporter) Stop() error {
	if exp.httpServer == nil {
		return nil
	}

	err := exp.httpServer.Close()
	if err != nil {
		return err
	}

	addr := net.JoinHostPort(exp.Addr, strconv.Itoa(exp.Port))
	log.Infof("prometheus exporter (%s) terminated", addr)

	return nil
}
