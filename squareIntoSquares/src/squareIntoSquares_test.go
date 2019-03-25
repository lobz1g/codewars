package src

import (
	"reflect"
	"testing"
)

func TestDecompose(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{"", args{11}, []int64{1, 2, 4, 10}},
		{"", args{50}, []int64{1, 3, 5, 8, 49}},
		{"", args{5}, []int64{3, 4}},
		{"", args{2}, []int64{}},
		{"", args{7}, []int64{2, 3, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decompose(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decompose() = %v, want %v", got, tt.want)
			}
		})
	}
}
