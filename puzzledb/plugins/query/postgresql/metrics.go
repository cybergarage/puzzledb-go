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

package postgresql

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	//nolint:exhaustruct
	mInsertLatency = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "postgresql_insert_latency",
		Help: "postgresql insert latency in milliseconds",
	})
	//nolint:exhaustruct
	mUpdateLatency = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "postgresql_update_latency",
		Help: "postgresql update latency in milliseconds",
	})
	//nolint:exhaustruct
	mDeleteLatency = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "postgresql_delete_latency",
		Help: "postgresql delete latency in milliseconds",
	})
	//nolint:exhaustruct
	mSelectLatency = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "postgresql_select_latency",
		Help: "postgresql select latency in milliseconds",
	})
)

func init() { //nolint:gochecknoinits
	prometheus.MustRegister(mInsertLatency)
	prometheus.MustRegister(mUpdateLatency)
	prometheus.MustRegister(mDeleteLatency)
	prometheus.MustRegister(mSelectLatency)
}
