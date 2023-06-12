package protein

func FromRNA(rna string) ([]string, error) {
	codon := ""
	codons := []string{}

	for pos, char := range rna {
		codon = codon + string(char)
		if (pos+1)%3 == 0 {
			codons = append(codons, codon)
			codon = ""
		}
	}
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
