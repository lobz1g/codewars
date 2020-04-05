package src

import (
	"testing"
)

func TestAddLetters(t *testing.T) {
	type args struct {
		letters []rune
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{name: "Testing for []rune{'a', 'b', 'c'}", args: args{[]rune{'a', 'b', 'c'}}, want: 'f'},
		{name: "Testing for []rune{'z'}", args: args{[]rune{'z'}}, want: 'z'},
		{name: "Testing for []rune{'a', 'b'}", args: args{[]rune{'a', 'b'}}, want: 'c'},
		{name: "Testing for []rune{'c'}", args: args{[]rune{'c'}}, want: 'c'},
		{name: "Testing for []rune{'z', 'a'}", args: args{[]rune{'z', 'a'}}, want: 'a'},
		{name: "Testing for []rune{'y', 'c', 'b'}", args: args{[]rune{'y', 'c', 'b'}}, want: 'd'},
		{name: "Testing for []rune{}", args: args{[]rune{}}, want: 'z'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddLetters(tt.args.letters); got != tt.want {
				t.Errorf("AddLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
