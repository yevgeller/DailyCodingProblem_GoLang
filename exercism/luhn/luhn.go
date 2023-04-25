package luhn

import "strings"

func Valid(id string) bool {
	//panic("Please implement the Valid function")
	for i, r := range id {
		if i%2 == 0 {
			if ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z') {

				letter := strings.ToUpper(string(r))

				if m[letter] != 0 {
					return false
				}
				m[letter] = 1

			}
		}
	}
}
