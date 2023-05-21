package isbn

import (
	"fmt"
	"regexp"
	"strconv"
)

func IsValidISBN(isbn string) bool {
	//panic("Please implement the IsValidISBN function")
	flag := regexp.MustCompile(`[X0-9\-]+`).MatchString(isbn)
	if !flag {
		fmt.Println("Failed: ", isbn)
		return false
	}
	counter := 10
	total := 0
	//fmt.Println("Assignment: ", isbn)
	for _, c := range isbn {
		//fmt.Println(string(c))

		if s, err := strconv.Atoi(string(c)); err == nil || counter == 1 && string(c) == "X" {

			if err != nil && counter > 1 && string(c) != "X" {
				return false
			}
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
	//fmt.Println("Total: ", total, " Result: ", total%11 == 0)
	return counter == 0 && total%11 == 0
}
