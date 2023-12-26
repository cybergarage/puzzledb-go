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

package redis

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	//nolint:exhaustruct
	mSetLatency = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "redis_set_latency",
		Help: "redis SET latency in milliseconds",
	})
	//nolint:exhaustruct
	mGetLatency = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "redis_get_latency",
		Help: "redis GET latency in milliseconds",
	})
	//nolint:exhaustruct
	mDelLatency = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "redis_del_latency",
		Help: "redis DEL latency in milliseconds",
	})
	//nolint:exhaustruct
	mHSetLatency = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "redis_hset_latency",
		Help: "redis HSET latency in milliseconds",
	})
	//nolint:exhaustruct
	mHGetLatency = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "redis_hget_latency",
		Help: "redis HGET latency in milliseconds",
	})
	//nolint:exhaustruct
	mHGetAllLatency = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "redis_hgetall_latency",
		Help: "redis HGETALL latency in milliseconds",
	})
	//nolint:exhaustruct
	mHDelLatency = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "redis_hdel_latency",
		Help: "redis HDEL latency in milliseconds",
	})
)

func init() { //nolint:gochecknoinits
	prometheus.MustRegister(mSetLatency)
	prometheus.MustRegister(mGetLatency)
	prometheus.MustRegister(mDelLatency)
	prometheus.MustRegister(mHSetLatency)
	prometheus.MustRegister(mHGetLatency)
	prometheus.MustRegister(mHGetAllLatency)
	prometheus.MustRegister(mHDelLatency)
}
