package src

import (
	"testing"
)

func TestTwiceAsOld(t *testing.T) {
	tests := []struct {
		name        string
		dadYearsOld int
		sonYearsOld int
		want        int
	}{
		{name: "22", dadYearsOld: 36, sonYearsOld: 7, want: 22},
		{name: "55", dadYearsOld: 55, sonYearsOld: 30, want: 5},
		{name: "42", dadYearsOld: 42, sonYearsOld: 21, want: 0},
		{name: "22", dadYearsOld: 22, sonYearsOld: 1, want: 20},
		{name: "29", dadYearsOld: 29, sonYearsOld: 0, want: 29},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwiceAsOld(tt.dadYearsOld, tt.sonYearsOld); got != tt.want {
				t.Errorf("TwiceAsOld() = %v, want %v", got, tt.want)
			}
		})
	}
}
