package src

import (
	"testing"
)

func TestHist(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"tpwaemuqxdmwqbqrjbeosjnejqorxdozsxnrgpgqkeihqwybzyymqeazfkyiucesxwutgszbenzvgxibxrlvmzihcb"}, "u  3     ***\rw  4     ****\rx  6     ******\rz  6     ******"},
		{"", args{"aaifzlnderpeurcuqjqeywdq"}, "u  2     **\rw  1     *\rz  1     *"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hist(tt.args.s); got != tt.want {
				t.Errorf("Hist() = %v, want %v", got, tt.want)
			}
		})
	}
}
