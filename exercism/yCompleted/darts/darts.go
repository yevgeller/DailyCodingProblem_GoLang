package darts

import "math"

func Score(x, y float64) int {

	len := math.Sqrt(x*x + y*y)

	if len <= 1 {
		return 10
	} else if len <= 5 {
		return 5
	} else if len <= 10.0 {
		return 1
	}
	return 0
}
