package service

import (
	"errors"
	model "github.com/ar4ie13/metrics/internal/model"
	"strconv"
)

var (
	ErrUnknownMetricType   = errors.New("unknown metric type")
	ErrIncorrectValueType  = errors.New("incorrect value type")
	ErrIncorrectMetricType = errors.New("incorrect metric type")
)

type Repository interface {
	SaveCounter(string, int64) error
	SaveGauge(string, float64) error
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
