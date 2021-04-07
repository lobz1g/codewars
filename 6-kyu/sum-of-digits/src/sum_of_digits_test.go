package src

import "testing"

func TestDigitalRoot(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{16}, 7},
		{"", args{195}, 6},
		{"", args{992}, 2},
		{"", args{167346}, 9},
		{"", args{0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DigitalRoot(tt.args.n); got != tt.want {
				t.Errorf("DigitalRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}
