package src

import (
	"testing"
)

func TestMaximumSubarraySum(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "", args: args{[]int{}}, want: 0},
		{name: "", args: args{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}}, want: 6},
		{name: "", args: args{[]int{-2, -1, -3, -4, -1, -2, -1, -5, -4}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaximumSubarraySum(tt.args.numbers); got != tt.want {
				t.Errorf("MaximumSubarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
