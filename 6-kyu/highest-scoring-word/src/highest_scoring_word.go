package src

import (
	"strings"
)

func High(s string) string {
	var max int32
	var word string
	for _, w := range strings.Split(s, " ") {
		var score int32
		for _, c := range w {
			score += c - 96
		}
		if max < score {
			max = score
			word = w
		}
	}
	return word
}
