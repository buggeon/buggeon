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
