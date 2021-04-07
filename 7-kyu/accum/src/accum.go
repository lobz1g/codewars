package src

import (
	"strings"
)

func Accum(s string) (result string) {
	for i, v := range s {
		result += strings.ToUpper(string(v))
		for j := 0; j < i; j++ {
			result += strings.ToLower(string(v))
		}
		if i != len(s)-1 {
			result += "-"
		}
	}
	return
}
