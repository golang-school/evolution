package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type HTTPServer struct {
	total    *prometheus.CounterVec
	current  *prometheus.GaugeVec
	duration *prometheus.HistogramVec
}

func NewHTTPServer() *HTTPServer {
	m := &HTTPServer{}

	// Настройка метрик

	return m
}

func (m *HTTPServer) TotalInc(method string, code int) {}

func (m *HTTPServer) CurrentSet(process string, value float64) {}

func (m *HTTPServer) Duration(method string, startTime time.Time) {}
