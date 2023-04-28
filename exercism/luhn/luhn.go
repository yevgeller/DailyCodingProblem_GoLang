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
	// if strLen%2 == 1 {
	// 	ongoingCtr--
	// }

	fmt.Println("Assignment: ", cleanId, " len: ", strLen, ", working with: ongoingCtr: ", ongoingCtr, " starting at", string(cleanId[ongoingCtr]))
	for ongoingCtr > 0 {
		val, _ := strconv.Atoi(string(cleanId[ongoingCtr]))
		if ongoingCtr%2 == 1 {
			fmt.Print("Working: ", val, " at pos ", ongoingCtr, ", ")
		} else {
			fmt.Print("Skipping ", val, " at pos ", ongoingCtr, ", ")
		}
		ongoingCtr--
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
