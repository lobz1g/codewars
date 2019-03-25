package src

import (
	"math"
	"sort"
	"strconv"
)

func SumOfDivided(lst []int) string {
	tmpRes := map[int]int{}
	allPrimeNums := findAllPrimeNumbers(getMax(lst))
	for _, i := range lst {
		for _, value := range allPrimeNums {
			if i%value == 0 {
				tmpRes[value] += i
			}
		}
	}
	return getResult(tmpRes)
}

func getResult(arr map[int]int) string {
	res := ""
	for _, value := range getSortedArr(arr) {
		res += "(" + strconv.Itoa(value) + " " + strconv.Itoa(arr[value]) + ")"
	}
	return res
}

func getSortedArr(arr map[int]int) []int {
	var res []int
	for value := range arr {
		res = append(res, value)
	}
	sort.Ints(res)
	return res
}

func findAllPrimeNumbers(max int) []int {
	var res []int
	for i := 1; i <= max; i++ {

		if isPrime(i) {
			res = append(res, i)
		}
	}
	return res
}

func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func getMax(v []int) int {
	sort.Ints(v)
	a := int(math.Abs(float64(v[0])))
	b := int(math.Abs(float64(v[len(v)-1])))
	if a >= b {
		return a
	} else {
		return b
	}
}
