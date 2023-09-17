package src

import (
	"reflect"
	"testing"
)

func TestHowMuch(t *testing.T) {
	tests := []struct {
		name string
		m    int
		n    int
		want [][3]string
	}{
		{
			name: "case 1",
			m:    1,
			n:    100,
			want: [][3]string{{"M: 37", "B: 5", "C: 4"}, {"M: 100", "B: 14", "C: 11"}},
		},
		{
			name: "case 2",
			m:    0,
			n:    200,
			want: [][3]string{{"M: 37", "B: 5", "C: 4"}, {"M: 100", "B: 14", "C: 11"}, {"M: 163", "B: 23", "C: 18"}},
		},
		{
			name: "case 3",
			m:    9990,
			n:    9992,
			want: [][3]string{{"M: 9991", "B: 1427", "C: 1110"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HowMuch(tt.m, tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HowMuch() = %v, want %v", got, tt.want)
			}
		})
	}
}
