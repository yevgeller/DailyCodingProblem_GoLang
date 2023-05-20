package isbn

import "fmt"

func IsValidISBN(isbn string) bool {
	//panic("Please implement the IsValidISBN function")
	counter := 10
	for _, c := range isbn {
		
		fmt.Println(int64(c))
	}

	return true
}
