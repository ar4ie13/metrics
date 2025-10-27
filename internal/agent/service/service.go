package service

import (
	"math/rand/v2"
	"runtime"

	"github.com/ar4ie13/metrics/internal/model"
)

type Service struct {
	Metrics map[string]model.Metrics
}

func NewService() *Service {
	return &Service{
		Metrics: make(map[string]model.Metrics),
	}
}

func (s *Service) GetMetricsStorage() *Service {
	return s
}

func (s *Service) ResetPollCount() {
	if s.Metrics["PollCount"].Delta != nil {
		*s.Metrics["PollCount"].Delta = 0
	}
}

func (s *Service) UpdateMetrics() Service {
	metricsSlice := getMetrics()
	if len(s.Metrics) == 0 {
		for _, metric := range metricsSlice {
			s.Metrics[metric.ID] = metric
		}
		return *s

	}
	for _, metric := range metricsSlice {
		if _, ok := s.Metrics[metric.ID]; ok {
			switch metric.MType {
			case "gauge":
				*s.Metrics[metric.ID].Value = *metric.Value
			case "counter":
				*s.Metrics[metric.ID].Delta += *metric.Delta
			}

		}
	}
	return *s
}

func getRandomValue() model.Metrics {
	value := rand.Float64()
	return model.Metrics{
		ID:    "RandomValue",
		MType: "gauge",
		Value: &value,
	}
}

func getPollCount() model.Metrics {
	var delta int64 = 1
	return model.Metrics{
		ID:    "PollCount",
		MType: "counter",
		Delta: &delta,
	}
}

func getGauge(id string, value float64) model.Metrics {
	return model.Metrics{
		ID:    id,
		MType: "gauge",
		Value: &value,
	}
}

func getCounter(id string, delta int64) model.Metrics {
	return model.Metrics{
		ID:    id,
		MType: "counter",
		Delta: &delta,
	}
}

func getMetrics() []model.Metrics {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	metrics := []model.Metrics{
		getGauge("Alloc", float64(memStats.Alloc)),
		getGauge("BuckHashSys", float64(memStats.BuckHashSys)),
		getCounter("Frees", int64(memStats.Frees)),
		getGauge("GCCPUFraction", float64(uint64(memStats.GCCPUFraction))),
		getGauge("GCSys", float64(memStats.GCSys)),
		getGauge("HeapAlloc", float64(memStats.HeapAlloc)),
		getGauge("HeapIdle", float64(memStats.HeapIdle)),
		getGauge("HeapInuse", float64(memStats.HeapInuse)),
		getGauge("HeapObjects", float64(memStats.HeapObjects)),
		getGauge("HeapReleased", float64(memStats.HeapReleased)),
		getGauge("HeapSys", float64(memStats.HeapSys)),
		getGauge("LastGC", float64(memStats.LastGC)),
		getCounter("Lookups", int64(memStats.Lookups)),
		getGauge("MCacheInuse", float64(memStats.MCacheInuse)),
		getGauge("MCacheSys", float64(memStats.MCacheSys)),
		getGauge("MSpanInuse", float64(memStats.MSpanInuse)),
		getGauge("MSpanSys", float64(memStats.MSpanSys)),
		getCounter("Mallocs", int64(memStats.Mallocs)),
		getGauge("NextGC", float64(memStats.NextGC)),
		getCounter("NumForcedGC", int64(memStats.NumForcedGC)),
		getCounter("NumGC", int64(memStats.NumGC)),
		getGauge("OtherSys", float64(memStats.OtherSys)),
		getCounter("PauseTotalNs", int64(memStats.PauseTotalNs)),
		getGauge("StackInuse", float64(memStats.StackInuse)),
		getGauge("StackSys", float64(memStats.StackSys)),
		getGauge("Sys", float64(memStats.Sys)),
		getCounter("TotalAlloc", int64(memStats.TotalAlloc)),
		getPollCount(),
		getRandomValue(),
	}
	return metrics
}
