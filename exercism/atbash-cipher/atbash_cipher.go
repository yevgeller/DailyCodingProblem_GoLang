package atbash

import (
	"strings"
	"unicode"
)

func Atbash(s string) string {
	normalized := strings.ToLower(s)
	m := loadMap()
	ret := ""
	word := ""
	for _, w := range normalized {
		if unicode.IsLetter(w) {
			word += m[string(w)]
		} else if unicode.IsDigit(w) {
			word += string(w)
		}
		if len(word) == 5 {
			ret += word + " "
			word = ""
		}
	}
	return strings.TrimSpace(ret + word)
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
