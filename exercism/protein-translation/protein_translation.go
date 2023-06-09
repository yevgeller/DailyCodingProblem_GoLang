package protein

import (
	"fmt"
)

func FromRNA(rna string) ([]string, error) {
	fmt.Printf("Assignment: %v ", rna)
	//panic("Please implement the FromRNA function")
	//ctr := 0
	codon := ""
	codons := []string{} //rename this to codons

	for pos, char := range rna {
		fmt.Printf("%s:%d ", string(char), pos)
		codon = codon + string(char)
		if (pos+1)%3 == 0 {
			//ctr = 0

			codons = append(codons, codon)
			codon = ""
		}
		//fmt.Printf("character %c starts at byte position %d\n", char, pos)
	}
	fmt.Printf("Codons: %v \n", codons)
	//a := []string{}

	lastCodon := ""
	result := []string{}
	for _, cdn := range codons {
		if string(cdn) != lastCodon {
			result = append(result, cdn)
			lastCodon = cdn
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
