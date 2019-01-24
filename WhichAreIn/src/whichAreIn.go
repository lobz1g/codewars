package src

import (
	"sort"
	"strings"
)

func InArray(array1 []string, array2 []string) []string {
	result := []string{}

	for _, i := range array1 {
		for _, j := range array2 {
			if !stringInSlice(i, result) && strings.Contains(j, i) {
				result = append(result, i)
			}
		}
	}

	sort.Strings(result)
	return result
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
