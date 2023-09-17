package src

func Zeros(n int) int {
	var res int
	for i := 5; i <= n; i *= 5 {
		res += n / i
	}
	return res
}
