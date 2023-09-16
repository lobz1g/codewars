package src

import "testing"

func TestCakes(t *testing.T) {
	tests := []struct {
		name      string
		recipe    map[string]int
		available map[string]int
		want      int
	}{
		{
			name:      "case 1",
			recipe:    map[string]int{"flour": 500, "sugar": 200, "eggs": 1},
			available: map[string]int{"flour": 1200, "sugar": 1200, "eggs": 5, "milk": 200},
			want:      2,
		},
		{
			name:      "case 2",
			recipe:    map[string]int{"apples": 3, "flour": 300, "sugar": 150, "milk": 100, "oil": 100},
			available: map[string]int{"sugar": 500, "flour": 2000, "milk": 2000},
			want:      0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Cakes(tt.recipe, tt.available); got != tt.want {
				t.Errorf("Cakes() = %v, want %v", got, tt.want)
			}
		})
	}
}
