package src

import (
	"strings"
)

func Solve(s string) string {
	chars := strings.Split(s, "")
	for i, j := 0, len(chars)-1; i <= j; i, j = i+1, j-1 {
		if chars[i] == " " {
			j++
		} else if chars[j] == " " {
			i--
		} else {
			chars[i], chars[j] = chars[j], chars[i]
		}
	}
	return strings.Join(chars, "")
}
