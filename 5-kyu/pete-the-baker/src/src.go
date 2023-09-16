package src

func Cakes(recipe, available map[string]int) int {
	res := -1
	for key, rec := range recipe {
		av, ok := available[key]
		if !ok {
			return 0
		}

		tmpRes := av / rec
		if res > tmpRes || res == -1 {
			res = tmpRes
		}
	}

	if res == -1 {
		return 0
	}
	return res
}
