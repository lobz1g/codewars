package src

import "testing"

func TestSumOfDivided(t *testing.T) {
	type args struct {
		lst []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{[]int{12, 15}}, "(2 12)(3 27)(5 15)"},
		{"", args{[]int{15, 21, 24, 30, 45}}, "(2 54)(3 135)(5 90)(7 21)"},
		{"", args{[]int{107, 158, 204, 100, 118, 123, 126, 110, 116, 100}}, "(2 1032)(3 453)(5 310)(7 126)(11 110)(17 204)(29 116)(41 123)(59 118)(79 158)(107 107)"},
		{"", args{[]int{-4847, 1364, -781, 950, -2583, 3280, -365, -219, -2163, -413, 2790, -3767}},
			"(2 8384)(3 -2175)(5 6655)(7 -5159)(11 583)(19 950)(31 4154)(37 -4847)(41 697)(59 -413)(71 -781)(73 -584)(103 -2163)(131 -4847)(3767 -3767)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfDivided(tt.args.lst); got != tt.want {
				t.Errorf("SumOfDivided() = %v, want %v", got, tt.want)
			}
		})
	}
}
