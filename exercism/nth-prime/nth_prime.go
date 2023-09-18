package prime

import (
	"errors"
	"fmt"
	"math"
)

func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("lol nice try")
	}
	limit := uint64(n)
	if limit < 100 {
		limit = 100
	}
	arr := createStorageArray(int(limit))
	limit = uint64(len(arr)) - 1
	sqrt := uint64(math.Sqrt(float64(limit)))
	fmt.Println("the size of array is now ", len(arr), ", limit is ", limit, ", sqrt is ", sqrt)

	arr[2] = true
	arr[3] = true
	var x, y uint64
	
	for x = 1; x <= sqrt; x++ {
		for y = 1; y <= sqrt; y++ {
			a := 4*x*x + y*y
			if a <= limit && (a%12 == 1 || a%12 == 5) {
				arr[a] = !arr[a]
			}

			a = 3*x*x + y*y
			if a <= limit && a%12 == 7 {
				arr[a] = !arr[a]
			}

			a = 3*x*x - y*y
			if x > y && a <= limit && a%12 == 11 {
				arr[a] = !arr[a]
			}
		}

	}
	for x = 5; x <= sqrt; x++ {
		if arr[x] == true {
			xSquared := x * x
			for y = xSquared; y <= limit; y += xSquared {
				arr[y] = false
			}
		}
	}

	b := 0

	for x = 0; x < uint64(len(arr)); x++ {
		if arr[x] == true {
			b += 1
			fmt.Print(x, ", ")
			if b == n {
				return int(x), nil
			}
		}

	}
	fmt.Println("stopped at ", x)

	return 1, nil
}

func createStorageArray(n int) []bool {
	ret := make([]bool, n*n*2)

	for i := range ret {
		ret[i] = false
	}

	return ret
}
