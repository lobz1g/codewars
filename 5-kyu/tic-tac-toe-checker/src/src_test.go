package src

import "testing"

func TestIsSolved(t *testing.T) {
	tests := []struct {
		name  string
		board [3][3]int
		want  int
	}{
		{
			name: "case 1",
			board: [3][3]int{
				{0, 0, 1},
				{0, 1, 2},
				{2, 1, 0},
			},
			want: -1,
		},
		{
			name: "case 2",
			board: [3][3]int{
				{1, 1, 1},
				{0, 2, 2},
				{0, 0, 0},
			},
			want: 1,
		},
		{
			name: "case 3",
			board: [3][3]int{
				{2, 1, 2},
				{2, 1, 1},
				{1, 1, 2},
			},
			want: 1,
		},
		{
			name: "case 4",
			board: [3][3]int{
				{2, 0, 2},
				{1, 0, 1},
				{1, 0, 2},
			},
			want: -1,
		},
		{
			name: "case 5",
			board: [3][3]int{
				{0, 2, 1},
				{1, 0, 2},
				{2, 1, 0},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSolved(tt.board); got != tt.want {
				t.Errorf("IsSolved() = %v, want %v", got, tt.want)
			}
		})
	}
}
