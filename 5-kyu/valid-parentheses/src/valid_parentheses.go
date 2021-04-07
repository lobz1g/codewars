package src

func ValidParentheses(parens string) bool {
	count := 0
	for _, c := range parens {
		if c == '(' {
			count++
		} else {
			count--
		}
		if count < 0 {
			return false
		}
	}
	return count == 0
}
