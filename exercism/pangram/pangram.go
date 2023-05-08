package pangram

import (
	"regexp"
	"strings"
)

func IsPangram(input string) bool {
	input = strings.ToLower(input)
	input = regexp.MustCompile(`[^a-z]+`).ReplaceAllString(input, "")
	//input = strings.Replace(input, " ", "", -1)
	input = strings.ToLower(input)
	if len(input) < 26 {
		return false
	}
	var data = make(map[rune]bool)
	for _, i := range input {
		data[i] = true
	}
	return len(data) == 26
}
