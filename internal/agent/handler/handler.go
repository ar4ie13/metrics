package handler

import (
	"fmt"
	"github.com/ar4ie13/metrics/internal/agent/service"
	"github.com/go-resty/resty/v2"
	"log"
	"strconv"
	"time"
)

type MyApiError struct {
	Code      int       `json:"code"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

type Service interface {
	GetMetricsStorage() *service.MetricsStorage
}

type AgentConfig interface {
	GetEndpointServerAddr() string
	GetPollInterval() int
	GetReportInterval() int
}

type Handler struct {
	s Service
	c AgentConfig
}

func NewHandler(s Service, c AgentConfig) *Handler {
	return &Handler{
		s: s,
		c: c,
	}
}

func (h *Handler) PostMetrics() {
	client := resty.New()
	client.SetRetryCount(3).
		// длительность ожидания между попытками
		SetRetryWaitTime(30 * time.Second).
		// длительность максимального ожидания
		SetRetryMaxWaitTime(90 * time.Second)
	var responseErr MyApiError

	metrics := h.s.GetMetricsStorage()
	for _, metric := range metrics.Metrics {
		var valueString string
		if metric.MType == "gauge" && metric.Value != nil {
			valueString = strconv.FormatFloat(*metric.Value, 'f', -1, 64)
			url := fmt.Sprintf("http://%s/update/%s/%s/%s", h.c.GetEndpointServerAddr(), metric.MType, metric.ID, valueString)
			_, err := client.R().SetError(&responseErr).Post(url)
			if err != nil {
				log.Print(err)
			}
		} else if metric.MType == "counter" && metric.Delta != nil {
			valueString = strconv.FormatInt(*metric.Delta, 10)
			url := fmt.Sprintf("http://%s/update/%s/%s/%s", h.c.GetEndpointServerAddr(), metric.MType, metric.ID, valueString)
			_, err := client.R().SetError(&responseErr).Post(url)
			if err != nil {
				log.Print(err)
			}
		}
	}

}
