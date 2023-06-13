package anagram

import "strings"

func Detect(subject string, candidates []string) []string {
	//panic("Please implement the Detect function")
subjectMap := toMap(subject)
for _, word := range candidates {
	if !(strings.EqualFold(subject, word)) {
candidateMap := toMap(word)

	}
}
	a := []string{}
	return a
}

func toMap(subject string) map[rune]int {
	m := make(map[rune]int)

	for _, runeValue := range subject {
		current := m[runeValue]
		m[runeValue] = current + 1
	}

	return m
}

func FromCodon(codon string) (string, error) {
	m := loadMap()
	return m[codon], nil
}

func loadMap() map[string]string {
	m := map[string]string{
		"AUG": "Methionine",
		"UUU": "Phenylalanine",
		"UUC": "Phenylalanine",
		"UUA": "Leucine",
		"UUG": "Leucine",
		"UCU": "Serine",
		"UCC": "Serine",
		"UCA": "Serine",
		"UCG": "Serine",
		"UAU": "Tyrosine",
		"UAC": "Tyrosine",
		"UGU": "Cysteine",
		"UGC": "Cysteine",
		"UGG": "Tryptophan",
		"UAA": "STOP",
		"UAG": "STOP",
		"UGA": "STOP",
	}
	return m
}
