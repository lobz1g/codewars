package src

import "regexp"

func Typist(s string) int {
	var res int
	for i, v := range regexp.MustCompile(`[A-Z]+|[a-z]+`).FindAllString(s, -1) {
		if (i == 0 && s[0] >= 'A' && s[0] <= 'Z') || i > 0 {
			res++
		}
		res += len(v)
	}
	return res
}
