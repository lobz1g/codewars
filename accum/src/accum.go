package src

import (
	"strings"
)

func Accum(s string) string {
	var result string
	for i := 0; i < len(s); i++ {
		result += strings.ToUpper(string(s[i]))
		for j := 0; j < i; j++ {
			result += strings.ToLower(string(s[i]))
		}
		if i != len(s)-1 {
			result += "-"
		}
	}

	return result
}
