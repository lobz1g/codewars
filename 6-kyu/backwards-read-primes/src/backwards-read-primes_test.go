package src

import (
	"reflect"
	"testing"
)

func TestBackwardsPrime(t *testing.T) {
	tests := []struct {
		name  string
		start int
		stop  int
		want  []int
	}{
		{
			name:  "",
			start: 1,
			stop:  100,
			want:  []int{13, 17, 31, 37, 71, 73, 79, 97},
		},
		{
			name:  "",
			start: 501,
			stop:  599,
			want:  nil,
		},
		{
			name:  "",
			start: 1,
			stop:  31,
			want:  []int{13, 17, 31},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BackwardsPrime(tt.start, tt.stop); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BackwardsPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}
