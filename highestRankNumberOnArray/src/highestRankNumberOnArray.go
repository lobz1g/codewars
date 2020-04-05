package src

func HighestRank(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	var frequency, n int
	for i, v := range m {
		if frequency < v {
			frequency = v
			n = i
		} else if frequency == v {
			if n < i {
				n = i
			}
		}
	}

	return n
}
