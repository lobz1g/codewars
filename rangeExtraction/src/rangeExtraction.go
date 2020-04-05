package src

import (
	"fmt"
)

func Solution(list []int) string {
	var res string
	for i := 0; i < len(list); i++ {
		if res != "" {
			res += ","
		}

		j := i + 1
		for j < len(list) && list[j]-list[j-1] == 1 {
			j++
		}

		if j-i > 2 {
			res += fmt.Sprintf("%d-%d", list[i], list[j-1])
		} else if j-i == 2 {
			res += fmt.Sprintf("%d,%d", list[i], list[j-1])
		} else {
			res += fmt.Sprintf("%d", list[i])
		}
		i = j - 1
	}
	return res
}
