package src

import (
	"testing"
)

func TestHighestRank(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "", args: args{[]int{12, 10, 8, 12, 7, 6, 4, 10, 12}}, want: 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HighestRank(tt.args.nums); got != tt.want {
				t.Errorf("HighestRank() = %v, want %v", got, tt.want)
			}
		})
	}
}
