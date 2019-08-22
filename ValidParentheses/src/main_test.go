package src

import (
	"testing"
)

func TestValidParentheses(t *testing.T) {
	type args struct {
		parens string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{"()"}, true},
		{"", args{")"}, false},
		{"", args{")(()))"}, false},
		{"", args{"("}, false},
		{"", args{"(())((()())())"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidParentheses(tt.args.parens); got != tt.want {
				t.Errorf("ValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
