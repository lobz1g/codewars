package src

import (
	"sort"
	"strings"
)

func TwoToOne(s1 string, s2 string) string {
	res := ""
	for _, value := range s1 + s2 {
		if !strings.Contains(res, string(value)) {
			res += string(value)
		}
	}
	resArr := strings.Split(res, "")
	sort.Strings(resArr)

	return strings.Join(resArr, "")
}
