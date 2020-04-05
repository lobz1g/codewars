package src

import (
	"regexp"
)

func alphanumeric(str string) bool {
	ok, _ := regexp.Match(`\W`, []byte(str))
	return !ok && len(str) != 0
}
