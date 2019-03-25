package src

func Decompose(n int64) []int64 {
	var res int64 = 0
	arr := []int64{n - 1}
	for res != n*n {
		var lastNum int64
		if res > n*n {
			lastNum = arr[len(arr)-1]
			arr = arr[:len(arr)-1]
			res -= lastNum * lastNum
		}

		if lastNum == 0 {
			tmp := arr[len(arr)-1] - 1
			if tmp <= 0 {
				if len(arr) > 1 {
					arr = arr[:len(arr)-1]
					arr[len(arr)-1]--
				} else {
					return []int64{}
				}
			} else {
				arr = append(arr, arr[len(arr)-1]-1)
			}
		} else {
			arr = append(arr, lastNum-1)
		}
		res = getSum(arr)
	}

	var resArr []int64
	for i := len(arr) - 1; i >= 0; i-- {
		resArr = append(resArr, arr[i])
	}
	return resArr
}

func getSum(arr []int64) int64 {
	var res int64
	for _, value := range arr {
		res += value * value
	}
	return res
}
