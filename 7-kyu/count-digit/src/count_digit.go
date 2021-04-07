package src

import (
	"strconv"
	"strings"
)

func NbDig(n int, d int) int {
	count := 0
	for i := 0; i <= n; i++ {
		count += strings.Count(strconv.Itoa(i*i), strconv.Itoa(d))
	}

	return count
}
