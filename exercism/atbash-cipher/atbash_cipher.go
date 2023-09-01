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

	for _, w := range normalized {
		if isLetter(w) {
			ret += m[string(w)]
		} else if unicode.IsDigit(w)  {
			ret += string(w)
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
