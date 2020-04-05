package src

func AddLetters(letters []rune) rune {
	if len(letters) == 0 {
		return 'z'
	}

	var sum rune = 0
	for _, v := range letters {
		sum += v - 'a' + 1
		if sum > 'z'-'a'+1 {
			sum -= 'z' - 'a' + 1
		}
	}

	return sum + 'a' - 1
}
