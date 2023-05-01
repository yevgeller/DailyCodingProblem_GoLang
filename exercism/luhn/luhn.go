package luhn

import (
	"regexp"
	"strconv"
	"strings"
)

func Valid(id string) bool {
	cleanId := strings.ReplaceAll(id, " ", "")
	flag := regexp.MustCompile(`[^0-9]+`).MatchString(cleanId)

	if flag {
		return false
	}

	strLen := len(cleanId)
	if strLen < 2 {
		return false
	}

	ongoingCtr := strLen - 1
	posCtr := 0
	sum := 0

	for ongoingCtr >= 0 {
		val, _ := strconv.Atoi(string(cleanId[ongoingCtr]))
		if posCtr%2 == 1 {
			digit := processDigit(val)
			sum += digit
		} else {
			sum += val
		}

		//fmt.Print(" Sum: ", sum, "\n")

		ongoingCtr--
		posCtr++
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
