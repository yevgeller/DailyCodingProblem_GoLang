package atbash

import (
	"strings"
	"unicode"
)

func Atbash(s string) string {
	//panic("Please implement the Atbash function")
	normalized := strings.ToLower(s)
	m := loadMap()
	ret := ""
	counter := 1
	for _, w := range normalized {
		if isLetter(w) {
			counter += 1
			ret += m[string(w)]
		} else if unicode.IsDigit(w) {
			counter += 1
			ret += string(w)
		}
		if counter%5 == 0 {
			ret += " "

		}
	}
	return ret
}
func isLetter(c rune) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}

func loadMap() map[string]string {
	m := map[string]string{
		"a": "z",
		"b": "y",
		"c": "x",
		"d": "w",
		"e": "v",
		"f": "u",
		"g": "t",
		"h": "s",
		"i": "r",
		"j": "q",
		"k": "p",
		"l": "o",
		"m": "n",
		"n": "m",
		"o": "l",
		"p": "k",
		"q": "j",
		"r": "i",
		"s": "h",
		"t": "g",
		"u": "f",
		"v": "e",
		"w": "d",
		"x": "c",
		"y": "b",
		"z": "a",
	}
	return m
}
