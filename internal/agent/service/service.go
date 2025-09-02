package service

import (
	models "github.com/ar4ie13/metrics/internal/model"
	"math/rand/v2"
	"runtime"
)

type MetricsStorage struct {
	Metrics map[string]Metrics
}

func NewMetricsStorage() *MetricsStorage {
	return &MetricsStorage{
		Metrics: make(map[string]Metrics),
	}
}

type Metrics models.Metrics

func (m *Metrics) getRandomValue() Metrics {
	value := rand.Float64()
	return Metrics{
		ID:    "RandomValue",
		MType: "gauge",
		Value: &value,
	}
}

func (m *Metrics) getPollCounter() Metrics {
	var delta int64 = 1
	return Metrics{
		ID:    "PollCounter",
		MType: "counter",
		Delta: &delta,
	}
}

func (m *Metrics) getGauge(id string, value float64) Metrics {
	return Metrics{
		ID:    id,
		MType: "gauge",
		Value: &value,
	}
}

func (m *Metrics) getCounter(id string, delta int64) Metrics {
	return Metrics{
		ID:    id,
		MType: "counter",
		Delta: &delta,
	}
}

func (m *Metrics) getMetrics() []Metrics {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	metrics := []Metrics{
		m.getGauge("Alloc", float64(memStats.Alloc)),
		m.getGauge("BuckHashSys", float64(memStats.BuckHashSys)),
		m.getCounter("Frees", int64(memStats.Frees)),
		m.getGauge("GCCPUFraction", float64(uint64(memStats.GCCPUFraction))),
		m.getGauge("GCSys", float64(memStats.GCSys)),
		m.getGauge("HeapAlloc", float64(memStats.HeapAlloc)),
		m.getGauge("HeapIdle", float64(memStats.HeapIdle)),
		m.getGauge("HeapInuse", float64(memStats.HeapInuse)),
		m.getGauge("HeapObjects", float64(memStats.HeapObjects)),
		m.getGauge("HeapReleased", float64(memStats.HeapReleased)),
		m.getGauge("HeapSys", float64(memStats.HeapSys)),
		m.getGauge("LastGC", float64(memStats.LastGC)),
		m.getCounter("Lookups", int64(memStats.Lookups)),
		m.getGauge("MCacheInuse", float64(memStats.MCacheInuse)),
		m.getGauge("MCacheSys", float64(memStats.MCacheSys)),
		m.getGauge("MSpanInuse", float64(memStats.MSpanInuse)),
		m.getGauge("MSpanSys", float64(memStats.MSpanSys)),
		m.getCounter("Mallocs", int64(memStats.Mallocs)),
		m.getGauge("NextGC", float64(memStats.NextGC)),
		m.getCounter("NumForcedGC", int64(memStats.NumForcedGC)),
		m.getCounter("NumGC", int64(memStats.NumGC)),
		m.getGauge("OtherSys", float64(memStats.OtherSys)),
		m.getCounter("PauseTotalNs", int64(memStats.PauseTotalNs)),
		m.getGauge("StackInuse", float64(memStats.StackInuse)),
		m.getGauge("StackSys", float64(memStats.StackSys)),
		m.getGauge("Sys", float64(memStats.Sys)),
		m.getCounter("TotalAlloc", int64(memStats.TotalAlloc)),
		m.getPollCounter(),
		m.getRandomValue(),
	}
	return metrics
}

func (ms *MetricsStorage) GetMetricsStorage() *MetricsStorage {
	return ms
}

func (ms *MetricsStorage) UpdateMetrics() MetricsStorage {
	var metrics Metrics
	metricsSlice := metrics.getMetrics()
	if len(ms.Metrics) == 0 {
		for _, metric := range metricsSlice {
			ms.Metrics[metric.ID] = metric
		}
		return *ms

	}
	for _, metric := range metricsSlice {
		if _, ok := ms.Metrics[metric.ID]; ok {
			switch metric.MType {
			case "gauge":
				*ms.Metrics[metric.ID].Value = *metric.Value
			case "counter":
				*ms.Metrics[metric.ID].Delta += *metric.Delta
			}

		}
	}
	return *ms
}
