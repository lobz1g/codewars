package src

import "strings"

func FirstNonRepeating(str string) string {
	allChars := getAllChars(str)
	for _, char := range str {
		if allChars[strings.ToLower(string(char))] == 1 {
			return string(char)
		}
	}
	return ""
}

func getAllChars(str string) map[string]int {
	allChars := map[string]int{}
	for _, char := range str {
		allChars[strings.ToLower(string(char))]++
	}
	return allChars
}
