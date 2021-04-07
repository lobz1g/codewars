package src

import "testing"

func TestDuplicate_count(t *testing.T) {
	type args struct {
		s1 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{""}, 0},
		{"", args{"abcde"}, 0},
		{"", args{"abcdea"}, 1},
		{"", args{"abcdeaB"}, 2},
		{"", args{"indivisibility"}, 1},
		{"", args{"Indivisibilities"}, 2},
		{"", args{"AABB"}, 2},
		{"", args{"aA11"}, 2},
		{"", args{"A11122222BB"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Duplicate_count(tt.args.s1); got != tt.want {
				t.Errorf("Duplicate_count() = %v, want %v", got, tt.want)
			}
		})
	}
}
