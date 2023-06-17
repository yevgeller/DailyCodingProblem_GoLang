package wordcount

import (
	"fmt"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	//panic("Please implement the WordCount function")
	fmt.Println("Assignment: ", phrase)
	words := []string{}
	word := ""
	for _, runeValue := range phrase {
		symbol := string(runeValue)
		if isPunctuation(symbol) {
			if len(word) > 0 {
				words = append(words, strings.ToLower(word))
			}
			fmt.Println("Word: ", word)
			word = ""
		} else {
			word += symbol
		}

	}
	fmt.Println("Remaining word: ", word, ", len: ", len(word))
	if len(word) > 0 {
		words = append(words, strings.ToLower(word))
	}

	m := make(map[string]int)
	for _, w := range words {
		currentCount := m[w]
		m[w] = currentCount + 1

	}

	return m
}

func isPunctuation(b string) bool {
	//b := string(a)
	result := b == "," || b == "." || b == "!" || b == ":" || b == ";" || b == "\n" || b == " " || b == "$" || b == "&" || b == "@" || b == "%" || b == "^"

result2 := strings.Index(",.!:;$&@%^", b) > -1 || b == " " || b == "\n"

if result != result2 {
	panic("incorrectly determining punctuation")
}

	fmt.Println("isPunctuation: ", b, " result: ")
	return result
}
