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

package ristretto

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	mRequestCount = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "ristretto_request_count",
		Help: "Ristretto request count",
	})
	mHitCount = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "ristretto_hit_count",
		Help: "Ristretto hit count",
	})
	mHitRate = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "ristretto_hit_rate",
		Help: "Ristretto hit rate",
	})
)

func init() { //nolint:gochecknoinits
	prometheus.MustRegister(mRequestCount)
	prometheus.MustRegister(mHitCount)
	prometheus.MustRegister(mHitRate)
}
