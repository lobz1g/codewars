package src

func Hammer(n int) uint {
	dp := make([]uint, n)
	for i := range dp {
		dp[i] = 1
	}

	x2, x3, x5 := uint(2), uint(3), uint(5)
	var i, j, k int

	for cnt := 1; cnt < n; cnt++ {
		dp[cnt] = min(x2, min(x3, x5))
		if x2 == dp[cnt] {
			i++
			x2 = 2 * dp[i]
		}
		if x3 == dp[cnt] {
			j++
			x3 = 3 * dp[j]
		}
		if x5 == dp[cnt] {
			k++
			x5 = 5 * dp[k]
		}
	}

	return dp[n-1]
}

func min(a, b uint) uint {
	if a <= b {
		return a
	}
	return b
}
