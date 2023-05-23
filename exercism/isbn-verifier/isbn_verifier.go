package isbn

import (
	"strconv"
)

func IsValidISBN(isbn string) bool {
	counter := 10
	total := 0
	calc := ""
	for _, ch := range isbn {
		numberCandidate, err := strconv.Atoi(string(ch))
		if counter == 1 && string(ch) == "X" && err != nil {
			numberCandidate = 10
		}
		if err != nil && string(ch) != "-" && counter > 1 {
			return false
		}
		if numberCandidate > 0 || err == nil {
			calc += string(ch) + "*" + strconv.Itoa(counter) + " "
			total += counter * numberCandidate
			counter--
		}
	}
	return counter == 0 && total%11 == 0
}
