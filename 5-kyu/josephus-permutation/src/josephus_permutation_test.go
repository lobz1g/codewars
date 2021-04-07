package src

import (
	"reflect"
	"testing"
)

func TestJosephus(t *testing.T) {
	type args struct {
		items []interface{}
		k     int
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{"", args{[]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1}, []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{"", args{[]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2}, []interface{}{2, 4, 6, 8, 10, 3, 7, 1, 9, 5}},
		{"", args{[]interface{}{1, 2, 3, 4, 5, 6, 7}, 3}, []interface{}{3, 6, 2, 7, 5, 1, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Josephus(tt.args.items, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Josephus() = %v, want %v", got, tt.want)
			}
		})
	}
}
