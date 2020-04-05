package src

import (
	"testing"
)

func TestCamelCase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "", args: args{"test case"}, want: "TestCase"},
		{name: "", args: args{"camel case method"}, want: "CamelCaseMethod"},
		{name: "", args: args{"say hello "}, want: "SayHello"},
		{name: "", args: args{" camel case word"}, want: "CamelCaseWord"},
		{name: "", args: args{""}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CamelCase(tt.args.s); got != tt.want {
				t.Errorf("CamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
