package src

import (
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "", args: args{[]int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}}, want: "-6,-3-1,3-5,7-11,14,15,17-20"},
		{name: "", args: args{[]int{-3, -2, -1, 2, 10, 15, 16, 18, 19, 20}}, want: "-3--1,2,10,15,16,18-20"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.list); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
