package src

import (
	"reflect"
	"testing"
)

func TestBackwardsPrime(t *testing.T) {
	type args struct {
		start int
		stop  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "", args: args{1, 100}, want: []int{13, 17, 31, 37, 71, 73, 79, 97}},
		{name: "", args: args{501, 599}, want: nil},
		{name: "", args: args{1, 31}, want: []int{13, 17, 31}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BackwardsPrime(tt.args.start, tt.args.stop); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BackwardsPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}
