package src

func MxDifLg(a1 []string, a2 []string) int {
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}

	maxDiff := 0
	for _, i := range a1 {
		for _, j := range a2 {
			diff := abs(len(i) - len(j))
			if diff > maxDiff {
				maxDiff = diff
			}
		}
	}

	return int(maxDiff)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
