package service

import (
	"reflect"
	"testing"

	"github.com/ar4ie13/metrics/internal/model"
)

func TestMetrics_getPollCounter(t *testing.T) {
	var value1 int64 = 1
	tests := []struct {
		name    string
		m       model.Metrics
		want    model.Metrics
		wantErr bool
	}{
		{
			name: "Correct result",
			m: model.Metrics{
				ID:    "PollCount",
				MType: "counter",
				Delta: &value1,
			},
			want: model.Metrics{
				ID:    "PollCount",
				MType: "counter",
				Delta: &value1,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPollCount(); !reflect.DeepEqual(got, tt.want) {
				if tt.wantErr {
					t.Errorf("getPollCount() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
