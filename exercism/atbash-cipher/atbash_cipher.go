package atbash

import (
	"fmt"
	"strings"
	"unicode"
)

func Atbash(s string) string {
	//panic("Please implement the Atbash function")
	normalized := strings.ToLower(s)
	m := loadMap()
	ret := ""
	counter := 0
	for _, w := range normalized {
		if isLetter(w) {
			counter += 1
			ret += m[string(w)]
			fmt.Println("'", string(w), "'->'", m[string(w)], "', result: ", ret, ", ret length", len(ret), " counter: ", counter)
		} else if unicode.IsDigit(w) {
			counter += 1
			ret += strings.TrimSpace(string(w))
			fmt.Println("attaching number '", strings.TrimSpace(string(w)), "' to ", ret, ", ret length", len(ret), " counter: ", counter)
		}
		if counter%5 == 0 {
			ret += " "

		}
		//fmt.Println("counter: ", counter, " ret:", ret, ", ret length", len(ret), ", w:", string(w))
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
