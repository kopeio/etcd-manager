/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package openstack

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type OpenstackMetrics struct {
	Duration *prometheus.HistogramVec
	Total    *prometheus.CounterVec
	Errors   *prometheus.CounterVec
}

// MetricContext indicates the context for OpenStack metrics.
type MetricContext struct {
	Start      time.Time
	Attributes []string
	Metrics    *OpenstackMetrics
}

// NewMetricContext creates a new MetricContext.
func NewMetricContext(resource string, request string) *MetricContext {
	return &MetricContext{
		Start:      time.Now(),
		Attributes: []string{resource + "_" + request},
	}
}

// Observe records the request latency and counts the errors.
func (mc *MetricContext) Observe(om *OpenstackMetrics, err error) error {
	if om == nil {
		// mc.RequestMetrics not set, ignore this request
		return nil
	}

	om.Duration.WithLabelValues(mc.Attributes...).Observe(
		time.Since(mc.Start).Seconds())
	om.Total.WithLabelValues(mc.Attributes...).Inc()
	if err != nil {
		om.Errors.WithLabelValues(mc.Attributes...).Inc()
	}
	return err
}

func RegisterMetrics() {
	doRegisterAPIMetrics()
}
