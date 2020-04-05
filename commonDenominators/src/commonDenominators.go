package src

import (
	"fmt"
)

func ConvertFracts(a [][]int) string {
	mult := a[0][1] / getFactor(a[0][0], a[0][1])
	for i := 1; i < len(a); i++ {
		n := a[i][1] / getFactor(a[i][0], a[i][1])
		mult = n * mult / getFactor(n, mult)
	}

	var res string
	for _, v := range a {
		res += fmt.Sprintf("(%0.f,%d)", float64(mult)/float64(v[1])*float64(v[0]), mult)
	}

	return res
}

func getFactor(a, b int) int {
	if b == 0 {
		return a
	}
	return getFactor(b, a%b)
}
