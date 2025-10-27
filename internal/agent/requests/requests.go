package requests

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/ar4ie13/metrics/internal/agent/config"
	"github.com/ar4ie13/metrics/internal/agent/service"
	"github.com/go-resty/resty/v2"
)

type MyAPIError struct {
	Code      int       `json:"code"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

type Service interface {
	GetMetricsStorage() *service.Service
}

type Requests struct {
	service Service
	config  *config.AgentConfig
}

func NewRequests(s Service, c config.AgentConfig) *Requests {
	return &Requests{
		service: s,
		config:  &c,
	}
}

func (h *Requests) PostMetrics() {
	client := resty.New()
	client.SetRetryCount(3).
		// длительность ожидания между попытками
		SetRetryWaitTime(5 * time.Second).
		// длительность максимального ожидания
		SetRetryMaxWaitTime(15 * time.Second)
	var responseErr MyAPIError

	metrics := h.service.GetMetricsStorage()
	for _, metric := range metrics.Metrics {
		var valueString string
		if metric.MType == "gauge" && metric.Value != nil {
			valueString = strconv.FormatFloat(*metric.Value, 'f', -1, 64)
			url := fmt.Sprintf("http://%s/update/%s/%s/%s", h.config.EndpointServerAddr, metric.MType, metric.ID, valueString)
			_, err := client.R().SetError(&responseErr).Post(url)
			if err != nil {
				log.Print(err)
			}
		} else if metric.MType == "counter" && metric.Delta != nil {
			valueString = strconv.FormatInt(*metric.Delta, 10)
			url := fmt.Sprintf("http://%s/update/%s/%s/%s", h.config.EndpointServerAddr, metric.MType, metric.ID, valueString)
			_, err := client.R().SetError(&responseErr).Post(url)
			if err != nil {
				log.Print(err)
			}
		}
	}

}
