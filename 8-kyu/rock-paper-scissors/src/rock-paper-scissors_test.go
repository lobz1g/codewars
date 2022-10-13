package src

import (
	"testing"
)

func TestRps(t *testing.T) {
	tests := []struct {
		name string
		p1   string
		p2   string
		want string
	}{
		{
			name: "p1",
			p1:   "rock",
			p2:   "scissors",
			want: "Player 1 won!",
		},
		{
			name: "p2",
			p1:   "scissors",
			p2:   "rock",
			want: "Player 2 won!",
		},
		{
			name: "draw",
			p1:   "rock",
			p2:   "rock",
			want: "Draw!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Rps(tt.p1, tt.p2); got != tt.want {
				t.Errorf("Rps() = %v, want %v", got, tt.want)
			}
		})
	}
}
