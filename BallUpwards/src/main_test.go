package src

import (
	"testing"
)

func Test_max_ball(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{15}, 4},
		{"", args{37}, 10},
		{"", args{25}, 7},
		{"", args{45}, 13},
		{"", args{99}, 28},
		{"", args{85}, 24},
		{"", args{136}, 39},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxBall(tt.args.v); got != tt.want {
				t.Errorf("MaxBall() = %v, want %v", got, tt.want)
			}
		})
	}
}
