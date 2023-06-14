package anagram

import (
	"fmt"
	"reflect"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	//panic("Please implement the Detect function")
	subjectMap := toMap(strings.ToLower(subject))
	result := []string{}
	for _, word := range candidates {
		if !(strings.EqualFold(subject, word)) {
			candidateMap := toMap(strings.ToLower(word))
			eq := reflect.DeepEqual(subjectMap, candidateMap)
			if eq {
				result = append(result, word)
			}
		}
	}
	return result
}

func toMap(subject string) map[rune]int {
	m := make(map[rune]int)

	for _, runeValue := range subject {
		current := m[runeValue]
		m[runeValue] = current + 1
	}
	fmt.Println("subject: ", subject, ", map: ", m)
	return m
}

// func FromCodon(codon string) (string, error) {
// 	m := loadMap()
// 	return m[codon], nil
// }

// func loadMap() map[string]string {
// 	m := map[string]string{
// 		"AUG": "Methionine",
// 		"UAG": "STOP",
// 		"UGA": "STOP",
// 	}
// 	return m
// }
