package anagram

import (
	"reflect"
	"strings"
)

func Detect(subject string, candidates []string) []string {
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
	return m
}
