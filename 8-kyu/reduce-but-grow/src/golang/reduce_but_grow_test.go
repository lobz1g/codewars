package src

import (
	"testing"
)

func TestGrow(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "", args: args{arr: []int{1, 2, 3}}, want: 6},
		{name: "", args: args{arr: []int{4, 1, 1, 1, 4}}, want: 16},
		{name: "", args: args{arr: []int{2, 2, 2, 2, 2, 2}}, want: 64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Grow(tt.args.arr); got != tt.want {
				t.Errorf("Grow() = %v, want %v", got, tt.want)
			}
		})
	}
}
