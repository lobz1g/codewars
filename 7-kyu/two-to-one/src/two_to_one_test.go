package src

import "testing"

func TestTwoToOne(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"aretheyhere", "yestheyarehere"}, "aehrsty"},
		{"2", args{"loopingisfunbutdangerous", "lessdangerousthancoding"}, "abcdefghilnoprstu"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoToOne(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("TwoToOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
