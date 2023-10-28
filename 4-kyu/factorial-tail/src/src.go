package src

import (
	"math"
	"sort"
)

func Zeroes(base, number int) int {
	var result []int
	delimiters := findDelimiters(base)
	for del, cnt := range delimiters {
		var res int
		for i := number / del; i > 0; i /= del {
			res += i
		}
		result = append(result, res/cnt)
	}
	sort.Ints(result)
	return result[0]
}

func findDelimiters(num int) map[int]int {
	delimiters := map[int]int{}
	for i := int(math.Sqrt(float64(num))); i > 1; i-- {
		if num == i {
			delimiters[i]++
			break
		}

		if num%i == 0 {
			for k, v := range findDelimiters(num / i) {
				delimiters[k] += v
			}
			for k, v := range findDelimiters(i) {
				delimiters[k] += v
			}
			break
		}
	}

	if len(delimiters) == 0 {
		delimiters[num]++
	}
	return delimiters
}
