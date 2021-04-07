package src

func DigitalRoot(n int) int {
	if n < 10 {
		return n
	}

	result := 0
	for n != 0 {
		result += n % 10
		n /= 10
	}
	return DigitalRoot(result)
}
