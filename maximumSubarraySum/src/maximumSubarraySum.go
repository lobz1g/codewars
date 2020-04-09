package src

func MaximumSubarraySum(numbers []int) int {
	allVariants := getAllVariants(numbers)

	var max int
	for _, arr := range allVariants {
		var sum int
		for _, v := range arr {
			sum += v
		}
		if sum > max {
			max = sum
		}
	}

	return max
}

func getAllVariants(arr []int) [][]int {
	result := [][]int{}
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			result = append(result, arr[i:j+1])
		}
	}
	return result
}
