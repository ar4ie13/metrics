package service

import (
	"reflect"
	"testing"
)

func TestMetrics_getPollCounter(t *testing.T) {
	var value1 int64 = 1
	var value2 int64 = 2
	tests := []struct {
		name string

		m    Metrics
		want Metrics
	}{
		{
			name: "Wrong result",
			m: Metrics{
				ID:    "PollCounter",
				MType: "counter",
				Delta: &value2,
			},
			want: Metrics{
				ID:    "PollCounter",
				MType: "counter",
				Delta: &value1,
			},
		},
		{
			name: "Correct result",
			m: Metrics{
				ID:    "PollCounter",
				MType: "counter",
				Delta: &value1,
			},
			want: Metrics{
				ID:    "PollCounter",
				MType: "counter",
				Delta: &value1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.getPollCounter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPollCounter() = %v, want %v", got, tt.want)
			}
		})
	}
}
