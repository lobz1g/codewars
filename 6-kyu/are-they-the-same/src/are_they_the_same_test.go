package src

import (
	"testing"
)

func TestComp(t *testing.T) {
	tests := []struct {
		name   string
		array1 []int
		array2 []int
		want   bool
	}{
		{
			name:   "",
			array1: []int{121, 144, 19, 161, 19, 144, 19, 11},
			array2: []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19},
			want:   true,
		},
		{
			name:   "",
			array1: []int{121, 144, 19, 161, 19, 144, 19, 11},
			array2: []int{11 * 21, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19},
			want:   false,
		},
		{
			name:   "",
			array1: nil,
			array2: []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Comp(tt.array1, tt.array2); got != tt.want {
				t.Errorf("Comp() = %v, want %v", got, tt.want)
			}
		})
	}
}
