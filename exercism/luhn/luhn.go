package luhn

import (
	"fmt"
	"regexp"
)

func Valid(id string) bool {
	//panic("Please implement the Valid function")
	
	cleanId := regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(id, "")

	
	strLen := len(cleanId)
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
		fmt.Println(cleanId[ongoingCtr])
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
