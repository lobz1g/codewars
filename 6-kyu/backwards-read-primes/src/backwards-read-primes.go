package src

import (
	"math/big"
)

func BackwardsPrime(start int, stop int) (result []int) {
	for i := start; i <= stop; i++ {
		num, revNum := int64(i), reverseInt(i)
		if num != revNum && big.NewInt(num).ProbablyPrime(0) && big.NewInt(revNum).ProbablyPrime(0) {
			result = append(result, i)
		}
	}

	if len(result) == 0 {
		return
	}
	return
}

func reverseInt(number int) int64 {
	var newNumber int
	for number > 0 {
		newNumber *= 10
		newNumber += number % 10
		number /= 10
	}
	return int64(newNumber)
}
