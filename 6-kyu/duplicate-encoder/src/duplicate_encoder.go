package src

import (
	"strings"
)

func DuplicateEncode(word string) string {
	word = strings.ToLower(word)
	var result string
	for _, v := range word {
		if strings.Count(word, string(v)) == 1 {
			result += "("
		} else {
			result += ")"
		}
	}
	return result
}
