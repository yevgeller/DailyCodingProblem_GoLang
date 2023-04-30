package luhn

import (
	"fmt"
	"regexp"
	"strconv"
)

func Valid(id string) bool {
	//panic("Please implement the Valid function")

	cleanId := regexp.MustCompile(`[^0-9]+`).ReplaceAllString(id, "")

	fmt.Println("Before: ", id, " After: ", cleanId)

	strLen := len(cleanId)
	if strLen < 2 {
		return false
	}

	ongoingCtr := strLen - 1

	fmt.Println("Assignment: ", cleanId, " len: ", strLen, ", working with: ongoingCtr: ", ongoingCtr, " starting at", string(cleanId[ongoingCtr]))
	sum := 0

	for ongoingCtr >= 0 {
		val, _ := strconv.Atoi(string(cleanId[ongoingCtr]))
		if ongoingCtr%2 == 1 {
			digit := processDigit(val)
			fmt.Print(" P: ", digit)
			sum += digit
			//fmt.Print("Working: ", val, " at pos ", ongoingCtr, ", ")
		} else {
			fmt.Print(" U: ", val)
			sum += val
			//fmt.Print("Skipping ", val, " at pos ", ongoingCtr, ", ")
		}

		fmt.Print(" Sum: ", sum, "\n")

		ongoingCtr--
	}

	return sum%10 == 0
}

func processDigit(digit int) int {
	digit = digit * 2
	if digit > 9 {
		digit = digit - 9
	}
	return digit
}
