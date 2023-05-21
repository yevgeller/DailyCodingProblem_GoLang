package isbn

import (
	"fmt"
	"strconv"
)

func IsValidISBN(isbn string) bool {
	//panic("Please implement the IsValidISBN function")
	counter := 10
	total := 0
	for _, c := range isbn {
		n, err = strconv.Atoi(string(c))
		if err == nil {
			total += n * counter
			counter--
		}
		fmt.Println(int64(c))
	}

	return true
}
