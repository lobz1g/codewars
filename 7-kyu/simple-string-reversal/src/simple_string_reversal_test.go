package src

import (
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		str  string
		want string
	}{
		{str: "codewars", want: "srawedoc"},
		{str: "your code", want: "edoc ruoy"},
		{str: "your code rocks", want: "skco redo cruoy"},
	}
	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			if got := Solve(tt.str); got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
