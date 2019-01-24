package main

import (
	"fmt"
	"strings"
)

func duplicate_count(s1 string) int {
	s1 = strings.ToLower(s1)
	result := make(map[string]int)
	for _, value := range s1 {
		if _, ok := result[string(value)]; ok {
			result[string(value)]++
		} else {
			result[string(value)] = 1

		}
	}
	count := 0
	for _, value := range result {
		if value > 1 {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(duplicate_count("absaS1133a"))
}
