package src

import "testing"

func TestTypist(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "case 1",
			s:    "a",
			want: 1,
		},
		{
			name: "case 2",
			s:    "aa",
			want: 2,
		},
		{
			name: "case 3",
			s:    "A",
			want: 2,
		},
		{
			name: "case 4",
			s:    "AA",
			want: 3,
		},
		{
			name: "case 5",
			s:    "aA",
			want: 3,
		},
		{
			name: "case 6",
			s:    "Aa",
			want: 4,
		},
		{
			name: "case 7",
			s:    "BeiJingDaXueDongMen",
			want: 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Typist(tt.s); got != tt.want {
				t.Errorf("Typist() = %v, want %v", got, tt.want)
			}
		})
	}
}
