package lsproduct

import (
	"errors"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span > len(digits) {
		return 0, errors.New("span must be smaller than string length")
	}
	if span < 1 {
		return 0, errors.New("span must be greater than 0")
	}
	converted := make([]int, 0)
	largestProduct := -1
	for _, c := range digits {
		digit, err := strconv.Atoi(string(c))

		if err != nil {
			return 0, err
		}
		converted = append(converted, digit)
	}

	for i, c := range converted {
		if i+span > len(converted) {
			return int64(largestProduct), nil
		}
		cnt := span - 1
		interimProduct := c
		for cnt > 0 {
			interimProduct *= converted[i+cnt]
			cnt -= 1
		}
		if largestProduct < interimProduct {
			largestProduct = interimProduct
		}
	}

	return int64(largestProduct), nil
}
