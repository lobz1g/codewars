package src

import (
	"testing"
)

func Test_alphanumeric(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "", args: args{".*?"}, want: false},
		{name: "", args: args{"a"}, want: true},
		{name: "", args: args{"Mazinkaiser"}, want: true},
		{name: "", args: args{"hello world_"}, want: false},
		{name: "", args: args{"PassW0rd"}, want: true},
		{name: "", args: args{"     "}, want: false},
		{name: "", args: args{""}, want: false},
		{name: "", args: args{"\n\t\n"}, want: false},
		{name: "", args: args{"ciao\n$$_"}, want: false},
		{name: "", args: args{"__ * __"}, want: false},
		{name: "", args: args{"&)))((("}, want: false},
		{name: "", args: args{"43534h56jmTHHF3k"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alphanumeric(tt.args.str); got != tt.want {
				t.Errorf("alphanumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}
