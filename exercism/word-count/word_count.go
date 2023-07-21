package wordcount

import (
	"fmt"
	"regexp"
	"strings"
)

type Frequency map[string]int

/* shame shame
func WordCount(input string) Frequency {
	result := map[string]int{}
	words := regexp.MustCompile(`\b[\p{L}\d']+\b`).FindAllString(input, -1)
	for _, w := range words {
		result[strings.ToLower(w)]++
	}
	return result
}
*/

func WordCount(phrase string) Frequency {
	words := strings.Fields(phrase)
	cleanWords := []string{}
	for _, w := range words {
		cleanWords = append(cleanWords, stripSpecialCharacters(w))
	}
	fmt.Println(phrase, ", arr: ", words, ", clean:", cleanWords)

	m := make(map[string]int)
	for _, w := range cleanWords {
		currentCount := m[w]
		m[w] = currentCount + 1

	}
	return m
}

func stripSpecialCharacters(word string) string {
	cleanWord := ""
	for i, runeValue := range word {
		lowerCase := strings.ToLower(string(runeValue))
		isChar, _ := regexp.MatchString(`[a-z\d]`, lowerCase)
		if isChar || (lowerCase == "'" && i > 0) {
			cleanWord += lowerCase
		}
	}
	if strings.HasSuffix(cleanWord, "'") {
		cleanWord = strings.TrimRight(cleanWord, "'")
	}
	return cleanWord
}

func WordCount2(phrase string) Frequency {
	words := []string{}
	word := ""
	for i, runeValue := range phrase {

		symbol := string(runeValue)
		if symbol == "'" {
			fmt.Println("Is contraction ", isContraction(phrase, i))
		}

		if symbol == "'" && isContraction(phrase, i) {
			word += symbol
		} else if isPunctuation(symbol) {
			if len(word) > 0 {
				words = append(words, strings.ToLower(word))
				word = ""
			}
		} else {
			word += symbol
		}
		fmt.Println("Word: ", word)
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
	fmt.Println("Symbol: ", b, ", isPunctuation: ", result)
	return result
}

func isContraction(phrase string, position int) bool {

	if position < 2 || position > len(phrase)-2 {
		return false
	}

	beforeChar := strings.ToLower(string(phrase[position-1]))
	afterChar := strings.ToLower(string(phrase[position+1]))
	beforeLiteral, _ := regexp.MatchString(`[a-z]`, beforeChar)
	afterLiteral, _ := regexp.MatchString(`[a-z]`, afterChar)
	//fmt.Println("inside isContraction, beforeLiteral: ", beforeLiteral, " beforeChar: ", beforeChar, ", afterLiteral: ", afterLiteral, ", afterChar: ", afterChar)
	// strings.ToLower(string([position-1]))
	return beforeLiteral && afterLiteral
}

//uncomment a test with an apostrophe
//test isContraction
