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

package memdb

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	//nolint:exhaustruct
	mReadLatency = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "memdb_read_latency",
		Help: "memdb read latency in milliseconds",
	})
	//nolint:exhaustruct
	mRangeReadLatency = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "memdb_range_read_latency",
		Help: "memdb range read latency in milliseconds",
	})
	//nolint:exhaustruct
	mWriteLatency = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "memdb_write_latency",
		Help: "memdb write latency in milliseconds",
	})
)

func init() { //nolint:gochecknoinits
	prometheus.MustRegister(mReadLatency)
	prometheus.MustRegister(mRangeReadLatency)
	prometheus.MustRegister(mWriteLatency)
}
