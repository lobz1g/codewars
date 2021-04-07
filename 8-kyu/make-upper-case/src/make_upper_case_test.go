package src

import (
	"testing"
)

func TestMakeUpperCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "", args: args{"hello"}, want: "HELLO"},
		{name: "", args: args{"hello world"}, want: "HELLO WORLD"},
		{name: "", args: args{"hello world !"}, want: "HELLO WORLD !"},
		{name: "", args: args{"heLlO wORLd !"}, want: "HELLO WORLD !"},
		{name: "", args: args{"1,2,3 hello world!"}, want: "1,2,3 HELLO WORLD!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeUpperCase(tt.args.str); got != tt.want {
				t.Errorf("MakeUpperCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
