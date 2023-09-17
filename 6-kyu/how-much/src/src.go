package src

import "fmt"

func HowMuch(m int, n int) [][3]string {
	if m > n {
		m, n = n, m
	}

	var res [][3]string
	for i := 0; ; i++ {
		sum := 63*i + 37
		if sum < m {
			continue
		}
		if sum > n {
			break
		}

		res = append(res, [3]string{
			fmt.Sprintf("M: %v", sum),
			fmt.Sprintf("B: %v", 9*i+5),
			fmt.Sprintf("C: %v", 7*i+4),
		})
	}

	return res
}
