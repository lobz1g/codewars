package main

import (
	"fmt"
)

func MxDifLg(a1 []string, a2 []string) int {
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}

	maxDiff := 0
	for _, i := range a1 {
		for _, j := range a2 {
			diff := Abs(len(i) - len(j))
			if diff > maxDiff {
				maxDiff = diff
			}
		}
	}

	return int(maxDiff)
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	var s1 = []string{"hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"}
	var s2 = []string{"cccooommaaqqoxiie", "gggqaffhhh", "tttoowwwmmww"}

	fmt.Println(MxDifLg(s1, s2))
}
