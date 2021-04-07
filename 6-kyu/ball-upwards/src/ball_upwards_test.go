package src

import (
	"testing"
)

func TestMaxBall(t *testing.T) {
	tests := []struct {
		name  string
		speed int
		want  int
	}{
		{"", 15, 4},
		{"", 37, 10},
		{"", 25, 7},
		{"", 45, 13},
		{"", 99, 28},
		{"", 85, 24},
		{"", 136, 39},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxBall(tt.speed); got != tt.want {
				t.Errorf("MaxBall() = %v, want %v", got, tt.want)
			}
		})
	}
}
