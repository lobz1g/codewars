package src

import (
	"fmt"
	"testing"
)

func TestChooseBestSum(t *testing.T) {
	type args struct {
		t  int
		k  int
		ls []int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{163, 3, []int{50, 55, 56, 57, 58}},
			want: 163,
		},
		{
			args: args{163, 3, []int{50}},
			want: -1,
		},
		{
			args: args{230, 3, []int{91, 74, 73, 85, 73, 81, 87}},
			want: 228,
		},
		{
			args: args{331, 2, []int{91, 74, 73, 85, 73, 81, 87}},
			want: 178,
		},
		{
			args: args{331, 4, []int{91, 74, 73, 85, 73, 81, 87}},
			want: 331,
		},
		{
			args: args{331, 5, []int{91, 74, 73, 85, 73, 81, 87}},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			if got := ChooseBestSum(tt.args.t, tt.args.k, tt.args.ls); got != tt.want {
				t.Errorf("ChooseBestSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
