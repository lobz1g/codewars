package src

const g = 9.81

func MaxBall(v int) int {
	for i := 0.0; ; i++ {
		if formula(float64(v), i) > formula(float64(v), i+1) {
			return int(i)
		}
	}
}

func formula(v, t float64) float64 {
	return 5*v/18*t/10 - 0.5*g*t*t/100
}
