package src

import (
	"strconv"
	"testing"
)

func TestNextPrime(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{
			n:    0,
			want: 2,
		},
		{
			n:    2,
			want: 3,
		},
		{
			n:    3,
			want: 5,
		},
		{
			n:    13,
			want: 17,
		},
		{
			n:    181,
			want: 191,
		},
		{
			n:    911,
			want: 919,
		},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.n), func(t *testing.T) {
			if got := NextPrime(tt.n); got != tt.want {
				t.Errorf("NextPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}
