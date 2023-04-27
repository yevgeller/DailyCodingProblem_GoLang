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

	ongoingCtr := strLen
	if strLen%2 == 1 {
		ongoingCtr--
	}

	fmt.Println("Assignment: ", cleanId, " len: ", strLen, ", working with: ongoingCtr: ", ongoingCtr)
	for ongoingCtr > 2 {
		val, _ := strconv.Atoi(string(cleanId[ongoingCtr]))
		fmt.Print(val, ":", ongoingCtr, " ")
		ongoingCtr -= 2
	}

	fmt.Println("\n--------")
	// for i, r := range id {
	// 	if i%2 == 0 {
	// 		if r {

	// 			letter := strings.ToUpper(string(r))

	// 			if m[letter] != 0 {
	// 				return false
	// 			}
	// 			m[letter] = 1

	// 		}
	// 	}
	// }

	return true
}
