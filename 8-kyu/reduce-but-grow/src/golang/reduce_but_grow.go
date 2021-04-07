package src

func Grow(arr []int) int {
	res := 1
	for _, v := range arr {
		res *= v
	}
	return res
}
