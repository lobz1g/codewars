package src

import (
	"strconv"
	"strings"
)

var fruits = []string{"kiwi", "pear", "kiwi", "banana", "melon", "banana", "melon",
	"pineapple", "apple", "pineapple", "cucumber", "pineapple", "cucumber", "orange",
	"grape", "orange", "grape", "apple", "grape", "cherry", "pear", "cherry", "pear",
	"kiwi", "banana", "kiwi", "apple", "melon", "banana", "melon", "pineapple", "melon",
	"pineapple", "cucumber", "orange", "apple", "orange", "grape", "orange", "grape",
	"cherry", "pear", "cherry", "pear", "apple", "pear", "kiwi", "banana", "kiwi", "banana",
	"melon", "pineapple", "melon", "apple", "cucumber", "pineapple", "cucumber", "orange",
	"cucumber", "orange", "grape", "cherry", "apple", "cherry", "pear", "cherry", "pear",
	"kiwi", "pear", "kiwi", "banana", "apple", "banana", "melon", "pineapple", "melon",
	"pineapple", "cucumber", "pineapple", "cucumber", "apple", "grape", "orange", "grape",
	"cherry", "grape", "cherry", "pear", "cherry", "apple", "kiwi", "banana", "kiwi", "banana",
	"melon", "banana", "melon", "pineapple", "apple", "pineapple"}

func CheckIt(n int) string {
	var sum int
	for _, v := range strings.Split(strconv.Itoa(n), "") {
		num, _ := strconv.Atoi(v)
		sum += num
	}

	n -= sum
	if n < len(fruits) {
		return fruits[n-1]
	}

	return CheckIt(n)
}

func SubtractSum(n int) string {
	return "apple"
}
