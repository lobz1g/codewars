package src

import (
	"reflect"
	"testing"
)

func TestPyramid(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "Testing for 0", args: args{0}, want: [][]int{}},
		{name: "Testing for 1", args: args{1}, want: [][]int{{1}}},
		{name: "Testing for 2", args: args{2}, want: [][]int{{1}, {1, 1}}},
		{name: "Testing for 3", args: args{3}, want: [][]int{{1}, {1, 1}, {1, 1, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Pyramid(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pyramid() = %v, want %v", got, tt.want)
			}
		})
	}
}
