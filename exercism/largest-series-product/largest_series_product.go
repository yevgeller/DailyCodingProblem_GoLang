package lsproduct

import (
	"errors"
	"fmt"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	//panic("Please implement the LargestSeriesProduct function")
	if span < len(digits) {
		return 0, errors.New("span must be greater or equal to string length")
	}
	//fmt.Println("digits: ", digits, " span: ", span)
	converted := make([]int, 0) //, len(digits))
	largestProduct := -1
	for _, c := range digits {
		digit, _ := strconv.Atoi(string(c))

		//fmt.Println(i, " => ", string(c))
		converted = append(converted, digit)
	}
	fmt.Println("converted: ", converted)

	for i, c := range converted {
		if i+span > len(converted) {
			fmt.Println("i: ", i, ", span: ", span, "len(converted):", len(converted), ", i+span > len(converted)")
			return int64(largestProduct), nil
		}
		cnt := span - 1
		interimProduct := c
		for cnt >= 0 {
			interimProduct *= converted[i+cnt]
			fmt.Println("new interimProduct", interimProduct, ", converted[i+cnt]", converted[i+cnt])
			cnt -= 1
		}
		if largestProduct < interimProduct {
			largestProduct = interimProduct
			fmt.Println("largestProduct is now ", largestProduct)
		}
	}

	return int64(largestProduct), nil
}
