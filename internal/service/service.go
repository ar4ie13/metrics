package service

import (
	"encoding/json"
	"errors"
	model "github.com/ar4ie13/metrics/internal/model"
	"github.com/ar4ie13/metrics/internal/repository"
	"strconv"
)

var (
	ErrUnknownMetricType  = errors.New("unknown metric type")
	ErrIncorrectValueType = errors.New("incorrect value type")
)

type Repository interface {
	SaveCounter(string, int64) error
	SaveGauge(string, float64) error
	GetAll() repository.MemStorage
	GetSpecific(string, string) (model.Metrics, error)
}

type Service struct {
	r Repository
}

func NewService(r Repository) *Service {
	return &Service{r}
}

func (s *Service) SaveMetric(metricName string, metricType string, value string) error {
	switch metricType {
	case model.Counter:
		counter, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return ErrIncorrectValueType
		}
		err = s.r.SaveCounter(metricName, counter)
		if err != nil {
			return err
		}
	case model.Gauge:
		gauge, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return ErrIncorrectValueType
		}
		err = s.r.SaveGauge(metricName, gauge)
		if err != nil {
			return err
		}
	default:
		return ErrUnknownMetricType
	}

	return nil
}

func (s *Service) GetAllMetrics() string {
	metrics := s.r.GetAll()
	result, _ := json.MarshalIndent(metrics, "", "\t")
	return string(result)
}

func (s *Service) GetSpecificMetric(metricName string, metricType string) (string, error) {
	metrics, err := s.r.GetSpecific(metricName, metricType)
	if err != nil {
		return "", err
	}
	result, _ := json.MarshalIndent(metrics, "", "\t")
	return string(result), nil
}
