package isbn

import (
	"fmt"
	"strconv"
)

func IsValidISBN(isbn string) bool {
	//panic("Please implement the IsValidISBN function")
	// flag := regexp.MustCompile(`[X 0-9\-]+`).MatchString(isbn)
	// if !flag {
	// 	fmt.Println("Failed: ", isbn)
	// 	return false
	// }
	counter := 10
	total := 0
	calc := ""
	fmt.Println("Assignment: ", isbn)
	for _, c := range isbn {
		//fmt.Println(string(c))

		a, err := strconv.Atoi(string(c))
		if counter == 1 && string(c) == "X" && err != nil {
			a = 10
			fmt.Println("last position is X")
		}
		if err != nil && string(c) != "-" && counter > 1 {
			fmt.Println("error on ", string(c), " counter: ", counter)
			return false
		}
		if a > 0 || err == nil {
			calc += string(c) + "*" + strconv.Itoa(counter) + " "
			total += counter * a
			counter--
		}
		// if s, err := strconv.Atoi(string(c)); err == nil || counter == 1 && string(c) == "X" {

		// 	// if err != nil && counter > 1 && string(c) != "X" {
		// 	// 	return false
		// 	// }
		// 	//fmt.Printf("%T, %v", s, s)
		// 	if string(c) == "X" {
		// 		s = 10
		// 	}
		// 	total += counter * s
		// 	counter--
		// }
		//	i, err = strconv.ParseInt(string(c), 10, 32)
		// if err == nil {
		// 	total += i * counter
		// 	counter--
		// }
		// fmt.Println(int64(c))
	}
	fmt.Println("Total: ", total, " Result: ", total%11 == 0 && counter == 0, " calc: ", calc)
	return counter == 0 && total%11 == 0
}
