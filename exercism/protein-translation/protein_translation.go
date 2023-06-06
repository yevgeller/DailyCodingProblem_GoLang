package protein

func FromRNA(rna string) ([]string, error) {
	panic("Please implement the FromRNA function")
}

func FromCodon(codon string) (string, error) {
	panic("Please implement the FromCodon function")
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	m := loadMap()

	return m[strings.ToLower(color)]
}

func loadMap() map[string]int {
	m := map[string]int{
		"black":  0,
		"brown":  1,
		"red":    2,
		"orange": 3,
		"yellow": 4,
		"green":  5,
		"blue":   6,
		"violet": 7,
		"grey":   8,
		"white":  9,
	}
	return m
}