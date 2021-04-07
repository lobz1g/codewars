package src

import (
	"testing"
)

func TestAddLetters(t *testing.T) {
	tests := []struct {
		name    string
		letters []rune
		want    rune
	}{
		{
			name:    "Testing for []rune{'a', 'b', 'c'}",
			letters: []rune{'a', 'b', 'c'},
			want:    'f',
		},
		{
			name:    "Testing for []rune{'z'}",
			letters: []rune{'z'},
			want:    'z',
		},
		{
			name:    "Testing for []rune{'a', 'b'}",
			letters: []rune{'a', 'b'},
			want:    'c',
		},
		{
			name:    "Testing for []rune{'c'}",
			letters: []rune{'c'},
			want:    'c',
		},
		{
			name:    "Testing for []rune{'z', 'a'}",
			letters: []rune{'z', 'a'},
			want:    'a',
		},
		{
			name:    "Testing for []rune{'y', 'c', 'b'}",
			letters: []rune{'y', 'c', 'b'},
			want:    'd',
		},
		{
			name:    "Testing for []rune{}",
			letters: []rune{},
			want:    'z',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddLetters(tt.letters); got != tt.want {
				t.Errorf("AddLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
