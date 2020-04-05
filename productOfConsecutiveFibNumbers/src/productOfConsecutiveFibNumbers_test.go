package src

import (
	"reflect"
	"testing"
)

func TestProductFib(t *testing.T) {
	type args struct {
		prod uint64
	}
	tests := []struct {
		name string
		args args
		want [3]uint64
	}{
		{name: "", args: args{4895}, want: [3]uint64{55, 89, 1}},
		{name: "", args: args{5895}, want: [3]uint64{89, 144, 0}},
		{name: "", args: args{74049690}, want: [3]uint64{6765, 10946, 1}},
		{name: "", args: args{84049690}, want: [3]uint64{10946, 17711, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProductFib(tt.args.prod); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductFib() = %v, want %v", got, tt.want)
			}
		})
	}
}
