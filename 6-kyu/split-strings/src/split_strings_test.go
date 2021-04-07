package src

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"",
			args{"abc"},
			[]string{"ab", "c_"},
		},
		{
			"",
			args{"dawsd"},
			[]string{"da", "ws", "d_"},
		},
		{
			"",
			args{"awsaws"},
			[]string{"aw", "sa", "ws"},
		},
		{
			"",
			args{""},
			[]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
