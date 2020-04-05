package BaumSweetSequence

import (
	"reflect"
	"testing"
)

func TestBaumSweet(t *testing.T) {
	arr := make([]int, 20)
	p := make(chan int, 100)
	go BaumSweet(p)

	for i := 0; i < 20; i++ {
		arr[i] = <-p
	}

	t.Run("", func(t *testing.T) {
		want := []int{1, 1, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 0, 1, 1, 0, 0, 1}
		if !reflect.DeepEqual(arr, want) {
			t.Errorf("BaumSweet() = %v, want %v", arr, want)
		}
	})
}
