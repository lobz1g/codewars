package src

import (
	"testing"
)

func TestQuarterOf(t *testing.T) {
	tests := []struct {
		name  string
		month int
		want  int
	}{
		{name: "March", month: 3, want: 1},
		{name: "August", month: 8, want: 3},
		{name: "November", month: 11, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuarterOf(tt.month); got != tt.want {
				t.Errorf("QuarterOf() = %v, want %v", got, tt.want)
			}
		})
	}
}
