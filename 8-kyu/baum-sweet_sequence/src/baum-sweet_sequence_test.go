package src

import (
	"reflect"
	"testing"
)

func TestBaumSweet(t *testing.T) {
	want := []int{1, 1, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 0, 1, 1, 0, 0, 1}
	got := make([]int, len(want))
	p := make(chan int, len(want))
	go BaumSweet(p)

	for i := 0; i < len(want); i++ {
		got[i] = <-p
	}

	t.Run("", func(t *testing.T) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("BaumSweet() = %v, want %v", got, want)
		}
	})
}
