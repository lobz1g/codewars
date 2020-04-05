package src

import (
	"testing"
)

func TestConvertFracts(t *testing.T) {
	type args struct {
		a [][]int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "", args: args{[][]int{{1, 2}, {1, 3}, {1, 4}}}, want: "(6,12)(4,12)(3,12)"},
		{name: "", args: args{[][]int{{69, 130}, {87, 1310}, {30, 40}}}, want: "(18078,34060)(2262,34060)(25545,34060)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertFracts(tt.args.a); got != tt.want {
				t.Errorf("ConvertFracts() = %v, want %v", got, tt.want)
			}
		})
	}
}
