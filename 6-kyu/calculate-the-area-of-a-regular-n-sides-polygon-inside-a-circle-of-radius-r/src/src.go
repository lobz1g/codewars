package src

import "math"

const PI = 3.141592653589793 // Use this value as pi in your calculations if necessary

func AreaOfPolygonInsideCircle(circleRadius float64, numberOfSides int) float64 {
	n := float64(numberOfSides)
	area := circleRadius * circleRadius * n / 2 * math.Sin(2*PI/n)
	return math.Round(area*1000) / 1000
}
