package src

import (
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "", args: args{"a"}, want: 0},
		{name: "", args: args{"bcd"}, want: 9},
		{name: "", args: args{"zea"}, want: 26},
		{name: "", args: args{"az"}, want: 26},
		{name: "", args: args{"baz"}, want: 26},
		{name: "", args: args{"aeiou"}, want: 0},
		{name: "", args: args{"uaoczei"}, want: 29},
		{name: "", args: args{"abababababfapeifapefijaefaepfjavnefjnfbhwyfnjsifjapnes"}, want: 143},
		{name: "", args: args{"codewars"}, want: 37},
		{name: "", args: args{"bup"}, want: 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.str); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
