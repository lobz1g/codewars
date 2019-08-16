package src

import (
	"unicode"
)

func wave(words string) []string {
	res := []string{}
	for i, c := range words {
		if c != ' ' {
			res = append(res, words[:i]+string(unicode.ToUpper(c))+words[i+1:])
		}
	}
	return res
}
