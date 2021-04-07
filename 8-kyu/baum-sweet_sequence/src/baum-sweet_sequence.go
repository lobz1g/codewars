package src

import (
	"fmt"
	"strings"
)

func BaumSweet(ch chan<- int) {
	for i := 0; ; i++ {
		if i == 0 || !strings.Contains(strings.ReplaceAll(fmt.Sprintf("%b", i), "00", "1"), "0") {
			ch <- 1
		} else {
			ch <- 0
		}
	}
}
