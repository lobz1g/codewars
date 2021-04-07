package src

import (
	"regexp"
	"strings"
)

func solve(str string) int {
	var result int
	r := regexp.MustCompile(`[aeiou]`)
	for _, v := range strings.Split(r.ReplaceAllString(str, " "), " ") {
		var sum int
		for _, c := range v {
			sum += int(c) % 96
		}
		if sum > result {
			result = sum
		}
	}
	return result
}
