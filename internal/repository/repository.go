package repository

import (
	"errors"
	model "github.com/ar4ie13/metrics/internal/model"
)

var (
	ErrIncorrectMetricType = errors.New("incorrect metric type")
	ErrIncorrectMetricName = errors.New("unknown metric name")
)

// MemStorage is used to store Metrics struct in the map
type MemStorage struct {
	Metrics map[string]model.Metrics
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		Metrics: make(map[string]model.Metrics),
	}
}

func (m *MemStorage) SaveCounter(metricName string, counter int64) error {
	if _, ok := m.Metrics[metricName]; ok {
		if m.Metrics[metricName].MType != model.Counter {
			return ErrIncorrectMetricType
		}
		*m.Metrics[metricName].Delta += counter
		return nil
	}
	m.Metrics[metricName] = model.Metrics{ID: metricName, MType: model.Counter, Delta: &counter}

	return nil
}

func (m *MemStorage) SaveGauge(metricName string, gauge float64) error {
	if _, ok := m.Metrics[metricName]; ok {
		if m.Metrics[metricName].MType != model.Gauge {
			return ErrIncorrectMetricType
		}
		*m.Metrics[metricName].Value = gauge
		return nil
	}
	m.Metrics[metricName] = model.Metrics{ID: metricName, MType: model.Gauge, Value: &gauge}

	return nil
}

func (m *MemStorage) GetAll() map[string]model.Metrics {
	return m.Metrics
}

func (m *MemStorage) GetSpecific(metricName string, metricType string) (model.Metrics, error) {
	if _, ok := m.Metrics[metricName]; ok {
		if m.Metrics[metricName].MType == metricType {
			return m.Metrics[metricName], nil
		}
	}
	return model.Metrics{}, ErrIncorrectMetricName
}
