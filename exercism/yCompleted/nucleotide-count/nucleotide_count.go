package dna

import (
	"fmt"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

func (d DNA) Counts() (Histogram, error) {

	results := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, nuc := range d {
		switch nuc {
		case 'A', 'C', 'G', 'T':
			results[nuc]++
		default:
			return nil, fmt.Errorf("invalid nucleotide '%c'", nuc)
		}
	}
	return results, nil

}
