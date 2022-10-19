package src

import (
	"math"
)

func NextPrime(n int) int {
	if n < 2 {
		return 2
	}

	for i := n + 1; ; i++ {
		if checkPrime(i) {
			return i
		}
	}
}

func checkPrime(n int) bool {
	for i := int(math.Sqrt(float64(n))); i >= 2; i-- {
		if n%i == 0 {
			return false
		}
	}
	return true
}
