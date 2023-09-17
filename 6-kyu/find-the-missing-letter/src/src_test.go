package src

import "testing"

func TestFindMissingLetter(t *testing.T) {
	tests := []struct {
		name  string
		chars []rune
		want  rune
	}{
		{
			name:  "case 1",
			chars: []rune{'a', 'b', 'c', 'd', 'f'},
			want:  'e',
		},
		{
			name:  "case 2",
			chars: []rune{'O', 'Q', 'R', 'S'},
			want:  'P',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMissingLetter(tt.chars); got != tt.want {
				t.Errorf("FindMissingLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
