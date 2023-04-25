package luhn

import (
	"fmt"
	"strings"
)

func Valid(id string) bool {
	//panic("Please implement the Valid function")
strLen := len(id)

ongoingCtr := strLen 
if strLen %2 == 1 {
	ongoingCtr--
}

for ongoingCtr >= 0 {
	ongoingCtr -= 2
	fmt.Println(id[ongoingCtr])
}

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
}
