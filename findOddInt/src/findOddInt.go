package src

import (
	"sort"
)

func FindOdd(seq []int) int {
	sort.Ints(seq)
	for i := len(seq) - 1; i > 0; i-- {
		if seq[i] == seq[i-1] {
			seq = seq[:len(seq)-2]
			i--
		} else {
			break
		}
	}
	return seq[len(seq)-1]
}
