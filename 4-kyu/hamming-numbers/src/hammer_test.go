package src

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		input    int
		expected uint
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 4},
		{5, 5},
		{6, 6},
		{7, 8},
		{8, 9},
		{9, 10},
		{10, 12},
		{11, 15},
		{12, 16},
		{13, 18},
		{14, 20},
		{15, 24},
		{16, 25},
		{17, 27},
		{18, 30},
		{19, 32},
	}

	for _, tt := range tests {
		if ret := Hammer(tt.input); ret != tt.expected {
			t.Errorf("Solution() = %v, want %v", ret, tt.expected)
		}
	}

}
