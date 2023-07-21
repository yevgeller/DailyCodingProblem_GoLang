package lsproduct

import (
	"errors"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	//panic("Please implement the LargestSeriesProduct function")
	if span < len(digits) {
		return 0, errors.New("span must be greater or equal to string length")
	}

	converted := make([]int, len(digits))
largestProduct := 0
for _, c := range digits {
digit, _ := strconv.Atoi(string(c))

		//fmt.Println(i, " => ", string(c))
		converted := append(converted, digit)
	}


	return 0, nil
}
