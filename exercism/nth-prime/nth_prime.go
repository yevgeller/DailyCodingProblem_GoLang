package prime

import (
	"errors"
	"fmt"
)

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	//panic("Please implement the Nth function")
	if n < 1 {
		return 0, errors.New("lol nice try")
	}
	arr := createStorageArray(n)
	fmt.Println("the size of array is now ", len(arr))
	return 1, nil
}

func createStorageArray(n int) []bool {
	ret := make([]bool, n*n)

	for i := range ret {
		ret[i] = false
	}

	return ret
}
