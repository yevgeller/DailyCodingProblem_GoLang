package resistorcolorduo

import "strings"

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	panic("Implement the Value function")

		m := loadMap()
	sum :=  m[strings.ToLower(colors[0])]*10
	    if(len(colors)>1) {
	        sum += m[strings.ToLower(colors[1])]
	    }

	    return sum
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
