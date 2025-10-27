package repository

import (
	model "github.com/ar4ie13/metrics/internal/model"
	"github.com/ar4ie13/metrics/internal/server/service"
)

// MemStorage is used to store Metrics struct in the map
type MemStorage struct {
	Metrics map[string]model.Metrics
}

// NewMemStorage constructs MemStorage repository object
func NewMemStorage() *MemStorage {
	return &MemStorage{
		Metrics: make(map[string]model.Metrics),
	}
}

// SaveCounter saves counter metric value to mem storage
func (m *MemStorage) SaveCounter(metricName string, counter int64) error {
	if _, ok := m.Metrics[metricName]; ok {
		if m.Metrics[metricName].MType != model.Counter {
			return service.ErrIncorrectMetricType
		}
		*m.Metrics[metricName].Delta += counter
		return nil
	}
	m.Metrics[metricName] = model.Metrics{ID: metricName, MType: model.Counter, Delta: &counter}

	return nil
}

// SaveGauge saves gauge metric value to mem storage
func (m *MemStorage) SaveGauge(metricName string, gauge float64) error {
	if _, ok := m.Metrics[metricName]; ok {
		if m.Metrics[metricName].MType != model.Gauge {
			return service.ErrIncorrectMetricType
		}
		*m.Metrics[metricName].Value = gauge
		return nil
	}
	m.Metrics[metricName] = model.Metrics{ID: metricName, MType: model.Gauge, Value: &gauge}

	return nil
}

// GetAll return all metrics with values stored in mem storage
func (m *MemStorage) GetAll() map[string]model.Metrics {
	return m.Metrics
}

// GetSpecific return requested metric with value from mem storage
func (m *MemStorage) GetSpecific(metricName string, metricType string) (model.Metrics, error) {
	if _, ok := m.Metrics[metricName]; ok {
		if m.Metrics[metricName].MType == metricType {
			return m.Metrics[metricName], nil
		}
	}
	return model.Metrics{}, service.ErrIncorrectMetricName
}
