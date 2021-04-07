package src

import (
	"fmt"
	"testing"
)

func TestSubtractSum(t *testing.T) {
	tests := []struct {
		num  int
		want string
	}{
		{num: 10, want: "apple"},
		{num: 11, want: "apple"},
		{num: 12, want: "apple"},
		{num: 13, want: "apple"},
		{num: 15, want: "apple"},
		{num: 20, want: "apple"},
		{num: 325, want: "apple"},
		{num: 330, want: "apple"},
		{num: 1000, want: "apple"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.num), func(t *testing.T) {
			if got := CheckIt(tt.num); got != tt.want {
				t.Errorf("CheckIt() = %v, want %v", got, tt.want)
			}
		})
	}
}
