package isogram

import (
	"strings"
)

func isLetter(c rune) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}

func IsIsogram(word string) bool {
	m := map[string]int{}
	for _, r := range word {
		if ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z') {

			letter := strings.ToUpper(string(r))

			if m[letter] != 0 {
				return false
			}
			m[letter] = 1

		}
	}
	return true
}
