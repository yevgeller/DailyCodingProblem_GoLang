package isogram

import (
	"fmt"
	"strings"
)

func isLetter(c rune) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}

func IsIsogram(word string) bool {
	//panic("Please implement the IsIsogram function")
	m := map[string]int{}
	fmt.Printf("Word: %s\n", word)
	for _, r := range word {
		//fmt.Printf("letter %s", string(r))
		//		if isLetter(r) {
		if ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z') {

			letter := strings.ToUpper(string(r))

			if m[letter] != 0 {
				fmt.Printf("duplicate letter %s\n", letter)
				return false
			}
			m[letter] = 1

		}
	}
	return true
}
