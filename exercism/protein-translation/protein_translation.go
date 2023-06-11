package protein

import (
	"fmt"
)

func FromRNA(rna string) ([]string, error) {
	codon := ""
	codons := []string{} //rename this to codons

	for pos, char := range rna {
		codon = codon + string(char)
		if (pos+1)%3 == 0 {
			codons = append(codons, codon)
			codon = ""
		}
		//fmt.Printf("character %c starts at byte position %d\n", char, pos)
	}
	fmt.Printf("Codons: %v \n", codons)
	//a := []string{}

	lastProtein := ""
	result := []string{}
	for _, cdn := range codons {
		protein, _ := FromCodon(string(cdn))
		if protein == "" || protein == "STOP" {
			break
		}
		if protein != lastProtein {
			result = append(result, protein)
			lastProtein = protein
		}
	}
	fmt.Printf("Result: %v \n", result)
	return result, nil
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
