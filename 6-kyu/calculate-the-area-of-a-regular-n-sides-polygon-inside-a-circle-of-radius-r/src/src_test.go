package src

import "testing"

func TestAreaOfPolygonInsideCircle(t *testing.T) {
	tests := []struct {
		name          string
		circleRadius  float64
		numberOfSides int
		want          float64
	}{
		{
			name:          "case 1",
			circleRadius:  3,
			numberOfSides: 3,
			want:          11.691,
		},
		{
			name:          "case 2",
			circleRadius:  5.8,
			numberOfSides: 7,
			want:          92.053,
		},
		{
			name:          "case 3",
			circleRadius:  4,
			numberOfSides: 5,
			want:          38.042,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AreaOfPolygonInsideCircle(tt.circleRadius, tt.numberOfSides); got != tt.want {
				t.Errorf("AreaOfPolygonInsideCircle() = %v, want %v", got, tt.want)
			}
		})
	}
}
