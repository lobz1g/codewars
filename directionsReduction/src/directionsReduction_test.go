package src

import (
	"reflect"
	"testing"
)

func TestDirReduc(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"", args{[]string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "SOUTH", "NORTH"}}, []string{"SOUTH"}},
		{"", args{[]string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST"}}, []string{"WEST"}},
		{"", args{[]string{"NORTH", "WEST", "SOUTH", "EAST"}}, []string{"NORTH", "WEST", "SOUTH", "EAST"}},
		{"", args{[]string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "NORTH"}}, []string{"NORTH"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DirReduc(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DirReduc() = %v, want %v", got, tt.want)
			}
		})
	}
}
