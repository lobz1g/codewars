package src

import "testing"

func TestFirstNonRepeating(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"a"}, "a"},
		{"", args{"stress"}, "t"},
		{"", args{"moonmen"}, "e"},
		{"", args{""}, ""},
		{"", args{"abba"}, ""},
		{"", args{"aa"}, ""},
		{"", args{"~><#~><"}, "#"},
		{"", args{"hello world, eh?"}, "w"},
		{"", args{"sTreSS"}, "T"},
		{"", args{"Go hang a salami, I'm a lasagna hog!"}, ","},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstNonRepeating(tt.args.str); got != tt.want {
				t.Errorf("FirstNonRepeating() = %v, want %v", got, tt.want)
			}
		})
	}
}
