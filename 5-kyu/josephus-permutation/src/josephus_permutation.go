package src

func Josephus(items []interface{}, k int) []interface{} {
	result := []interface{}{}
	j := 0
	for i := 0; i < len(items); i++ {
		j++
		if k == j {
			result = append(result, items[i])
			items = append(items[:i], items[i+1:]...)
			i--
			j = 0
		}
		if i == len(items)-1 {
			i = -1
		}
	}
	return result
}
