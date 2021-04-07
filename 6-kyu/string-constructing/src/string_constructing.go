package src

func StringConstructing(a, s string) int {
	count := 0
	result := ""
	for result != s {
		if len(s) != len(result) {
			result += a
			count++
		}

		for i := 0; i < len(result); i++ {
			if len(s) == i {
				count += len(result) - len(s)
				result = result[:len(s)]
				break
			} else if string(result[i]) != string(s[i]) {
				result = result[0:i] + result[i+1:]
				count++
				i--
			}
		}

	}
	return count
}
