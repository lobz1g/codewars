package src

import (
	"math/big"
)

func BackwardsPrime(start int, stop int) []int {
	result := []int{}
	for i := start; i <= stop; i++ {
		n, revN := int64(i), int64(reverseInt(i))
		if n != revN && big.NewInt(n).ProbablyPrime(0) && big.NewInt(revN).ProbablyPrime(0) {
			result = append(result, i)
		}
	}

	if len(result) == 0 {
		return nil
	}
	return result
}

func reverseInt(n int) int {
	newInt := 0
	for n > 0 {
		remainder := n % 10
		newInt *= 10
		newInt += remainder
		n /= 10
	}
	return newInt
}
