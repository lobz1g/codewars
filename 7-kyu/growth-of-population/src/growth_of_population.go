package src

func NbYear(p0 int, percent float64, aug int, p int) int {
	var count int
	for p0 < p {
		p0 = p0 + int(float64(p0)*percent/100) + aug
		count++
	}
	return count
}
