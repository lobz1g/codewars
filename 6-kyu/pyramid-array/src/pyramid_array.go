package src

func Pyramid(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			res[i] = append(res[i], 1)
		}
	}

	return res
}
