package src

func Perimeter(n int) int {
	result := 0
	fib := fibonacciSequence(n)
	for i := 0; i < len(fib); i++ {
		result += fib[i]
	}
	result *= 4
	return result
}

func fibonacciSequence(n int) []int {
	result := []int{1, 1}
	for i := 1; i < n; i++ {
		result = append(result, result[i]+result[i-1])
	}
	return result
}
