package repository

import (
	model "github.com/ar4ie13/metrics/internal/model"
	"github.com/ar4ie13/metrics/internal/service"
)

// MemStorage is used to store metrics struct in the map
type MemStorage struct {
	metrics map[string]model.Metrics
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		metrics: make(map[string]model.Metrics),
	}
}

func (m *MemStorage) SaveCounter(metricName string, counter int64) error {
	if _, ok := m.metrics[metricName]; ok {
		if m.metrics[metricName].MType != model.Counter {
			return service.ErrIncorrectMetricType
		}
		*m.metrics[metricName].Delta += counter
		return nil
	}
	m.metrics[metricName] = model.Metrics{MType: model.Counter, Delta: &counter}

	return nil
}

func (m *MemStorage) SaveGauge(metricName string, gauge float64) error {
	if _, ok := m.metrics[metricName]; ok {
		if m.metrics[metricName].MType != model.Gauge {
			return service.ErrIncorrectMetricType
		}
		*m.metrics[metricName].Value = gauge
		return nil
	}
	m.metrics[metricName] = model.Metrics{MType: model.Gauge, Value: &gauge}

	return nil
}
