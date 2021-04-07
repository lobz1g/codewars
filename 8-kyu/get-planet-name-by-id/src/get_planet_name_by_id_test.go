package src

import (
	"testing"
)

func TestGetPlanetName(t *testing.T) {
	tests := []struct {
		ID   int
		want string
	}{
		{ID: 1, want: "Mercury"},
		{ID: 2, want: "Venus"},
		{ID: 3, want: "Earth"},
		{ID: 4, want: "Mars"},
		{ID: 5, want: "Jupiter"},
		{ID: 8, want: "Neptune"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := GetPlanetName(tt.ID); got != tt.want {
				t.Errorf("GetPlanetName() = %v, want %v", got, tt.want)
			}
		})
	}
}
