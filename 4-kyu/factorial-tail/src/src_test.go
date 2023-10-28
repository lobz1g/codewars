package src

import "testing"

func TestZeroes(t *testing.T) {
	tests := []struct {
		name   string
		base   int
		number int
		want   int
	}{
		{
			name:   "10-10",
			base:   10,
			number: 10,
			want:   2,
		},
		{
			name:   "16-16",
			base:   16,
			number: 16,
			want:   3,
		},
		{
			name:   "12-5",
			base:   12,
			number: 5,
			want:   1,
		},
		{
			name:   "40-10",
			base:   40,
			number: 10,
			want:   2,
		},
		{
			name:   "17-16",
			base:   17,
			number: 16,
			want:   0,
		},
		{
			name:   "12-26",
			base:   12,
			number: 26,
			want:   10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Zeroes(tt.base, tt.number); got != tt.want {
				t.Errorf("Zeroes() = %v, want %v", got, tt.want)
			}
		})
	}
}
