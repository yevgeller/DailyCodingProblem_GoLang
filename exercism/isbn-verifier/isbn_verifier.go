package isbn

import "fmt"

func IsValidISBN(isbn string) bool {
	//panic("Please implement the IsValidISBN function")
	for i, c := range isbn {
		fmt.Println(i, " => ", string(c))
	}

	return true
}
