package src

import (
	"math"
)

func Cycle(n int) int {
	if n%5 == 0 || n%2 == 0 || n%10 == 0 {
		return -1
	}

	rem := 10
	arr := map[int]bool{}

	for i := 0; i < n; i++ {
		t, _ := math.Modf(float64(rem / n))
		if t == 0 {
			rem *= 10
		} else {
			rem = (rem - int(t)*n) * 10
		}
		if _, ok := arr[rem]; ok {
			return len(arr)
		}
		arr[rem] = true
	}

	return -1
}
