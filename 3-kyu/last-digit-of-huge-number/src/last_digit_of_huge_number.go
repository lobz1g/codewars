package src

func LastDigit(as []int) int {
	if len(as) == 0 {
		return 1
	}

	for key := len(as) - 1; key > 0; key-- {
		if as[key] == 0 {
			as[key-1] = 1
			as = as[:key]
			continue
		} else if as[key] == 1 {
			as = as[:key]
			continue
		}

		if as[key-1]%10 == 2 || as[key-1]%10 == 3 || as[key-1]%10 == 7 || as[key-1]%10 == 8 {
			switch as[key] % 4 {
			case 0:
				as[key-1] = pow(as[key-1], 4)
			case 1:
				as[key-1] = pow(as[key-1], 5)
			case 2:
				as[key-1] = pow(as[key-1], 6)
			case 3:
				as[key-1] = pow(as[key-1], 7)
			}
		} else if as[key-1]%10 == 4 || as[key-1]%10 == 9 || as[key-1]%10 == 1 {
			switch as[key] % 2 {
			case 0:
				as[key-1] = pow(as[key-1], 2)
			case 1:
				as[key-1] = pow(as[key-1], 3)
			}
		} else if as[key-1]%10 == 5 {
			as[key-1] = 25
		} else if as[key-1]%10 == 6 {
			as[key-1] = 36
		} else if as[key-1]%10 == 0 {
			as[key-1] *= as[key-1]
		}
		as = as[:key]
	}

	return as[0] % 10
}

func pow(a, b int) int {
	result := getDigs(a)
	for i := 1; i < b; i++ {
		result = getDigs(result * a)
	}

	return result
}

func getDigs(a int) int {
	if a > 1000 {
		return a % 1000
	}

	return a
}
