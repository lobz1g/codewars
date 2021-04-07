package src

import (
	"testing"
)

func TestLastDigit(t *testing.T) {
	type args struct {
		as []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{}}, 1},
		{"2", args{[]int{0, 0}}, 1},
		{"3", args{[]int{0, 0, 0}}, 0},
		{"4", args{[]int{1, 2}}, 1},
		{"5", args{[]int{2013, 2013}}, 3},
		{"6", args{[]int{3, 4, 5}}, 1},
		{"7", args{[]int{4, 3, 6}}, 4},
		{"8", args{[]int{3, 3, 3}}, 7},
		{"9", args{[]int{7, 6, 21}}, 1},
		{"10", args{[]int{7, 7, 7, 7, 3, 3, 7}}, 3},
		{"11", args{[]int{12, 30, 21}}, 6},
		{"12", args{[]int{2, 0, 1}}, 1},
		{"13", args{[]int{2, 2, 2, 0}}, 4},
		{"14", args{[]int{2, 2, 2}}, 6},
		{"15", args{[]int{2, 2, 2, 0, 2}}, 4},
		{"16", args{[]int{937640, 767456, 981242}}, 0},
		{"17", args{[]int{123232, 694022, 140249}}, 6},
		{"18", args{[]int{499942, 898102, 846073}}, 6},
		{"19", args{[]int{499947, 898103, 846077}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LastDigit(tt.args.as); got != tt.want {
				t.Errorf("LastDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}
