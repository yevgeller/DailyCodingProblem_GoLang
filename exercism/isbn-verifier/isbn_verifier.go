package isbn

import (
	"fmt"
	"regexp"
	"strconv"
)

func IsValidISBN(isbn string) bool {
	//panic("Please implement the IsValidISBN function")
	flag := regexp.MustCompile(`[0-9\-X]+`).MatchString(isbn)
	if !flag {
		return false
	}
	counter := 10
	total := 0
	fmt.Println("Assignment: ", isbn)
	for _, c := range isbn {
		//fmt.Println(string(c))

		if s, err := strconv.Atoi(string(c)); err == nil || counter == 1 && string(c) == "X" {
			//fmt.Printf("%T, %v", s, s)
			if string(c) == "X" {
				s = 10
			}
			total += counter * s
			counter--
		}
		//	i, err = strconv.ParseInt(string(c), 10, 32)
		// if err == nil {
		// 	total += i * counter
		// 	counter--
		// }
		// fmt.Println(int64(c))
	}
	fmt.Println("Total: ", total, " Result: ", total%11 == 0)
	return counter == 0 && total%11 == 0
}
