package src

import (
	"testing"
)

func TestCountSheeps(t *testing.T) {
	tests := []struct {
		name    string
		numbers []bool
		want    int
	}{
		{
			name: "17",
			numbers: []bool{
				true, true, true, false,
				true, true, true, true,
				true, false, true, false,
				true, false, false, true,
				true, true, true, true,
				false, false, true, true,
			},
			want: 17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountSheeps(tt.numbers); got != tt.want {
				t.Errorf("CountSheeps() = %v, want %v", got, tt.want)
			}
		})
	}
}
