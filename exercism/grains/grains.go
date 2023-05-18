package grains

import "math"

func Square(number int) (uint64, error) {
	//panic("Please implement the Square function")
	return uint64(math.Pow(2, float64(number))), nil
}

func Total() uint64 {
	//panic("Please implement the Total function")
	return uint64(math.Pow(2, 64))
}
