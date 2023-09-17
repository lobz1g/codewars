package src

import "testing"

func TestZeros(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "case 1",
			n:    6,
			want: 1,
		},
		{
			name: "case 2",
			n:    0,
			want: 0,
		},
		{
			name: "case 3",
			n:    30,
			want: 7,
		},
		{
			name: "case 4",
			n:    12,
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Zeros(tt.n); got != tt.want {
				t.Errorf("Zeros() = %v, want %v", got, tt.want)
			}
		})
	}
}
