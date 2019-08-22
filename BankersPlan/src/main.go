package src

import (
	"math"
)

func Fortune(f0 int, p float64, c0 int, n int, i float64) bool {
	if f0 < 0 {
		return false
	}

	if n == 1 {
		return true
	}

	f := math.Floor(float64(f0) + p/100*float64(f0) - float64(c0))
	c := math.Floor(float64(c0) + float64(c0)*i/100)

	return Fortune(int(f), p, int(c), n-1, i)
}
