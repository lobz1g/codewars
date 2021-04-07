package src

import (
	"math"
)

func TwiceAsOld(dadYearsOld, sonYearsOld int) int {
	res := dadYearsOld - sonYearsOld*2
	return int(math.Abs(float64(res)))
}
