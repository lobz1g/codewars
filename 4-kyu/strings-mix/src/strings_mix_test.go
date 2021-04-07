package src

import "testing"

func TestMix(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"Are they here", "yes, they are here"}, "2:eeeee/2:yy/=:hh/=:rr"},
		{"2", args{"uuuuuu", "uuuuuu"}, "=:uuuuuu"},
		{"3", args{"looping is fun but dangerous", "less dangerous than coding"}, "1:ooo/1:uuu/2:sss/=:nnn/1:ii/2:aa/2:dd/2:ee/=:gg"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mix(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("Mix() = %v, want %v", got, tt.want)
			}
		})
	}
}
