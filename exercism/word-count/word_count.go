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
	for i, runeValue := range phrase {
		//if i > 1 {
		fmt.Println("Next: ", string(phrase[i]), ", Prev: ", string(phrase[i-1]))
		//}
		symbol := string(runeValue)
		if (symbol == "'" && i > 2 && i < len(phrase)-2) || isPunctuation(symbol) {
			if len(word) > 0 {
				fmt.Println("Word: ", word)
				words = append(words, strings.ToLower(word))
			}

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
	//result_old := b == "," || b == "." || b == "!" || b == ":" || b == ";" || b == "\n" || b == " " || b == "$" || b == "&" || b == "@" || b == "%" || b == "^"
	result := strings.Index(",.!:;$&@%^'", b) > -1 || b == " " || b == "\n"
	return result
}

func isContraction(b, phrase string, position int) bool {
	if b != "'" {
		return false
	}
	if position < 2 || position > len(phrase)-1 {
		return false
	}
return input = regexp.MustCompile(`[^a-z]+`) strings.ToLower(string([position-1]))
}
