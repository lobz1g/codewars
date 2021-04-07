package src

import (
	"testing"
)

func TestNbYear(t *testing.T) {
	type args struct {
		p0      int
		percent float64
		aug     int
		p       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "", args: args{1500, 5, 100, 5000}, want: 15},
		{name: "", args: args{1500000, 2.5, 10000, 2000000}, want: 10},
		{name: "", args: args{1500000, 0.25, 1000, 2000000}, want: 94},
		{name: "", args: args{1500000, 0.25, -1000, 2000000}, want: 151},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NbYear(tt.args.p0, tt.args.percent, tt.args.aug, tt.args.p); got != tt.want {
				t.Errorf("NbYear() = %v, want %v", got, tt.want)
			}
		})
	}
}
