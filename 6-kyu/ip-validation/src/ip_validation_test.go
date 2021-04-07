package src

import "testing"

func Test_is_valid_ip(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"12.255.56.1"}, true},
		{"2", args{""}, false},
		{"3", args{"abc.def.ghi.jkl"}, false},
		{"4", args{"123.456.789.0"}, false},
		{"5", args{"12.34.56"}, false},
		{"6", args{"12.34.56 .1"}, false},
		{"7", args{"12.34.56.-1"}, false},
		{"8", args{"123.045.067.089"}, false},
		{"9", args{"127.1.1.0"}, true},
		{"10", args{"0.0.0.0"}, true},
		{"11", args{"0.34.82.53"}, true},
		{"12", args{"192.168.1.300"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Is_valid_ip(tt.args.ip); got != tt.want {
				t.Errorf("Is_valid_ip() = %v, want %v", got, tt.want)
			}
		})
	}
}
