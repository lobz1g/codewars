package src

func Zeros(n int) int {
	var res int
	for n > 0 {
		n /= 5
		res += n
	}
	return res
}
