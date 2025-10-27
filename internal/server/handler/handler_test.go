package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	handlerConfig "github.com/ar4ie13/metrics/internal/server/handler/config"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type MetricsTest struct {
	ID    string  `json:"id"`
	MType string  `json:"type"`
	Delta int64   `json:"delta,omitempty"`
	Value float64 `json:"value,omitempty"`
}

type MockConfig struct {
	handlerConfig handlerConfig.Config
}

type MockService struct {
	allMetrics map[string]MetricsTest
	err        error
}

func (s *MockService) GetAllMetrics() (string, error) {
	result, err := json.Marshal(s.allMetrics)
	if err != nil {
		panic(err)
	}
	return string(result), nil
}
func (s *MockService) GetSpecificMetric(metricName string, metricType string) (string, error) {
	if _, ok := s.allMetrics[metricName]; ok {
		switch metricType {
		case "counter":
			return strconv.FormatInt(s.allMetrics[metricName].Delta, 10), nil
		case "gauge":
			return strconv.FormatFloat(s.allMetrics[metricName].Value, 'f', -1, 64), nil
		}

	}
	return "", s.err
}
func (s *MockService) SaveMetric(metricName string, metricType string, value string) error {
	return s.err
}

func testRequest(t *testing.T, ts *httptest.Server, method,
	path string) (*http.Response, string) {
	req, err := http.NewRequest(method, ts.URL+path, nil)
	require.NoError(t, err)

	resp, err := ts.Client().Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	return resp, string(respBody)
}

func TestHandler_GetAllMetrics(t *testing.T) {

	tests := []struct {
		name           string
		fields         MockService
		url            string
		expectedValue  string
		expectedStatus int
		expectedErr    error
	}{
		{
			name: "success",
			fields: MockService{
				allMetrics: map[string]MetricsTest{
					"PollCounter": {
						ID:    "PollCounter",
						MType: "counter",
						Delta: 56,
					},
				},
			},
			url:            "/",
			expectedStatus: 200,
			expectedValue:  `{"PollCounter":{"id":"PollCounter","type":"counter","delta":56}}`,
		},
	}
	for _, v := range tests {
		s := &MockService{
			allMetrics: v.fields.allMetrics,
		}
		c := MockConfig{handlerConfig: handlerConfig.Config{LocalServerAddr: "localhost:8080"}}
		h := NewHandler(s, c.handlerConfig)
		router := chi.NewRouter()
		router.Get("/", h.GetAllMetrics)
		ts := httptest.NewServer(router)

		resp, get := testRequest(t, ts, "GET", v.url)

		assert.Equal(t, v.expectedStatus, resp.StatusCode)
		assert.Equal(t, v.expectedValue, get)
		resp.Body.Close()
		ts.Close()
	}

}
