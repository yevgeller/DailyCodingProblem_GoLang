package luhn

import (
	"fmt"
)

func Valid(id string) bool {
	//panic("Please implement the Valid function")
	strLen := len(id)
	if strLen < 2 {
		return false
	}


	ongoingCtr := strLen
	if strLen%2 == 1 {
		ongoingCtr--
	}

	fmt.Println("Assignment: ", id, " len: ", strLen, ", working with: ongoingCtr: ", ongoingCtr)
	for ongoingCtr >= 0 {
		ongoingCtr -= 2
		fmt.Println(id[ongoingCtr])
	}

	fmt.Println("--------")
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
