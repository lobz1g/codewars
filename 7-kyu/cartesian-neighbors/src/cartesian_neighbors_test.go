package src

import (
	"reflect"
	"testing"
)

func TestCartesianNeighbor(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "",
			args: args{2, 2},
			want: [][]int{
				{1, 1},
				{1, 2},
				{1, 3},
				{2, 1},
				{2, 3},
				{3, 1},
				{3, 2},
				{3, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CartesianNeighbor(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartesianNeighbor() = %v, want %v", got, tt.want)
			}
		})
	}
}
