package src

import (
	"fmt"
	"strings"
)

var alph = []string{"u", "w", "x", "z"}

func Hist(s string) string {
	res := ""
	for _, v := range alph {
		count := strings.Count(s, v)
		if count > 0 {
			if res != "" {
				res += "\r"
			}
			res += fmt.Sprintf("%s  %d     %s", v, count, strings.Repeat("*", count))
		}
	}

	return res
}
