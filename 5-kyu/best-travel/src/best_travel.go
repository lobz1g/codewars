package src

func ChooseBestSum(t, k int, ls []int) int {
	arr := getAllPossibles(k, ls)
	result := -1
	for _, v := range arr {
		if result < v && v <= t {
			result = v
		}
	}
	return result
}

func getAllPossibles(k int, ls []int) []int {
	if k == 1 {
		return ls
	}

	result := []int{}
	for i := 0; i < len(ls)-k+1; i++ {
		arr := getAllPossibles(k-1, ls[i+1:])
		for _, v := range arr {
			result = append(result, ls[i]+v)
		}
	}
	return result
}
