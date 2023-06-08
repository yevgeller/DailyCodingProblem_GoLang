package protein

import "strings"

func FromRNA(rna string) ([]string, error) {
	panic("Please implement the FromRNA function")
}

func FromCodon(codon string) (string, error) {
	panic("Please implement the FromCodon function")
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) string {
	m := loadMap()

	return m[strings.ToLower(color)]
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
