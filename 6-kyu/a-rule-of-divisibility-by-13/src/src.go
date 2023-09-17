package src

var sequence = map[int]int{0: 1, 1: 10, 2: 9, 3: 12, 4: 3, 5: 4}

func Thirt(n int) int {
	if n < 100 {
		return n
	}

	var res int
	for i := 0; n != 0; n, i = n/10, i+1 {
		if i == len(sequence) {
			i = 0
		}
		res += sequence[i] * (n % 10)
	}

	return Thirt(res)
}
