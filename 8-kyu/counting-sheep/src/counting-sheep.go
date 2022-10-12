package src

import (
	"fmt"
	"strings"
)

func CountSheeps(numbers []bool) int {
	return strings.Count(fmt.Sprintf("%v", numbers), "true")
}
