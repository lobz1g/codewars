package src

import (
	"testing"
)

func TestCycle(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{33}, 2},
		{"", args{18118}, -1},
		{"", args{69}, 22},
		{"", args{197}, 98},
		{"", args{65}, -1},
		{"", args{1234567}, 34020},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Cycle(tt.args.n); got != tt.want {
				t.Errorf("Cycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
