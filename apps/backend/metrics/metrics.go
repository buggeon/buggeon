// Buggeon - SelfHosted service for bug and task tracking
// Copyright (C) 2026 DEVE corp.

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics struct {
	RequestsTotal      *prometheus.CounterVec
	RequestDuration    *prometheus.HistogramVec
	RequestsInProgress *prometheus.GaugeVec
	RequestSize        *prometheus.SummaryVec
	ResponseSize       *prometheus.SummaryVec

	UsersTotal    prometheus.Counter
	ProjectsTotal prometheus.Counter
	CardsTotal    prometheus.Counter
	BoardsTotal   prometheus.Counter
	MessagesTotal prometheus.Counter

	MembersByProject prometheus.Counter
	BoardsByProject  prometheus.Counter
	CardsByProject   prometheus.Counter
	CardsByBoard     prometheus.Counter
}

var metrics *Metrics

func InitMetrics() *Metrics {
	if metrics != nil {
		return metrics
	}

	m := &Metrics{}

	m.RequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "requests_total",
		},
		[]string{""},
	)

	return m
}
