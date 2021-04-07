package src

func DblLinear(n int) int {
	res := []int{1}
	x, y := 0, 0
	for i := 0; i < n; i++ {
		tmpX := res[x]*2 + 1
		tmpY := res[y]*3 + 1
		if tmpX <= tmpY {
			x++
			res = append(res, tmpX)
			if tmpY == tmpX {
				y++
			}
		} else {
			y++
			res = append(res, tmpY)

		}
	}
	return res[n]
}
