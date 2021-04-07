package src

import (
	"strings"
)

func toWeirdCase(str string) string {
	var result []string
	for _, w := range strings.Split(str, " ") {
		var weirdWord string
		for j := range w {
			if j%2 == 0 {
				weirdWord += strings.ToUpper(string(w[j]))
			} else {
				weirdWord += strings.ToLower(string(w[j]))
			}
		}
		result = append(result, weirdWord)
	}

	return strings.Join(result, " ")
}
