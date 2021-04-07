package src

import (
	"fmt"
	"math/big"
)

func Decomp(n int) string {
	fact := factorial(int64(n))
	var result string

	var i int64 = 2
	for fact.Int64() != 1 {
		var j int64 = 0
		mod := getMod(fact, i)
		if mod != 0 {
			i++
			continue
		}
		for mod == 0 {
			fact.Div(fact, big.NewInt(i))
			mod = getMod(fact, i)
			j++
		}

		if result != "" {
			result += " * "
		}

		if j == 1 {
			result += fmt.Sprintf("%d", i)
		} else {
			result += fmt.Sprintf("%d^%d", i, j)
		}
	}
	return result
}

func getMod(a *big.Int, del int64) int64 {
	tmp := new(big.Int).Set(a)
	tmp.Mod(tmp, big.NewInt(del))
	return tmp.Int64()
}

func factorial(n int64) *big.Int {
	result := big.NewInt(1)
	for i := n; i > 0; i-- {
		result.Mul(result, big.NewInt(i))
	}
	return result
}
