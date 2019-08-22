package src

import (
	"testing"
)

func TestFortune(t *testing.T) {
	type args struct {
		f0 int
		p  float64
		c0 int
		n  int
		i  float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{100000, 1.0, 2000, 15, 1.0}, true},
		{"", args{100000, 1.0, 9185, 12, 1.0}, false},
		{"", args{100000000, 1.0, 100000, 50, 1.0}, true},
		{"", args{100000000, 1.5, 10000000, 50, 1.0}, false},
		{"", args{100000000, 5.0, 1000000, 50, 1.0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fortune(tt.args.f0, tt.args.p, tt.args.c0, tt.args.n, tt.args.i); got != tt.want {
				t.Errorf("Fortune() = %v, want %v", got, tt.want)
			}
		})
	}
}
