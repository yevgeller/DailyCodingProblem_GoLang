package prime

import (
	"errors"
	"fmt"
	"math"
)

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	//panic("Please implement the Nth function")
	if n < 1 {
		return 0, errors.New("lol nice try")
	}
	if n == 1 {
		return 2, nil
	}
	if n == 2 {
		return 3, nil
	}
	arr := createStorageArray(n)
	fmt.Println("the size of array is now ", len(arr))
	limit := uint64(n)
	sqrt := uint64(math.Floor(math.Sqrt(float64(n))))
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

		for x = 5; x <= sqrt; x++ {
			if arr[x] == true {
				for y = x * x; y <= limit; y *= y {
					arr[y] = false
				}
			}
		}

		arr[2] = true
		arr[3] = true
	}

	for x = 0; x < limit; x++ {
		if arr[x] == true {
			fmt.Print(x, ", ")
		}
	}

	// for (ulong x = 1; x <= sqrt; x++)
	//     for (ulong y = 1; y <= sqrt; y++)
	//     {
	//         var n = 4*x*x + y*y;
	//         if (n <= limit && (n % 12 == 1 || n % 12 == 5))
	//             isPrime[n] ^= true;

	//         n = 3*x*x + y*y;
	//         if (n <= limit && n % 12 == 7)
	//             isPrime[n] ^= true;

	//         n = 3*x*x - y*y;
	//         if (x > y && n <= limit && n % 12 == 11)
	//             isPrime[n] ^= true;
	//     }

	// for (ulong n = 5; n <= sqrt; n++)
	//     if (isPrime[n])
	//         for (ulong k = n*n; k <= limit; k *= k)
	//             isPrime[k] = false;

	// primes.Add(2);
	// primes.Add(3);
	// for (ulong n = 5; n <= limit; n++)
	//     if (isPrime[n])
	//         primes.Add(n);

	return 1, nil
}

func createStorageArray(n int) []bool {
	ret := make([]bool, n*n)

	for i := range ret {
		ret[i] = false
	}

	return ret
}
