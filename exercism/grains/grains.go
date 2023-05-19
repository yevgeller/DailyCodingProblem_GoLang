package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	//panic("Please implement the Square function")
	if number <= 0 || number > 64 {
		return 0, errors.New("bleh")
	}
	if number == 1 {
		return 1, nil
	}
	return uint64(math.Pow(2, float64(number-1))), nil
}

func Total() uint64 {
	//panic("Please implement the Total function")
	return math.MaxUint64
	//return uint64(1 + math.Pow(2, 65))
}
