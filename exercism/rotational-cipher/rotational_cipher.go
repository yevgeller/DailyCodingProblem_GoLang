package rotationalcipher

import (
	"fmt"
	"strings"
)

func RotationalCipher(plain string, shiftKey int) string {
	//panic("Please implement the RotationalCipher function")
	positionByLetter := mapPositionByLetter()
	letterByPosition := mapLetterByPosition()
	result := ""
	for _, ch := range plain {
		//fmt.Println("ch: ", string(ch), " isLetter: ", isLetter(ch))
		if isLetter(ch) {
			plainIndex := positionByLetter[strings.ToUpper(string(ch))]
			isUpperCase := isUpperCase(ch)                //if is upper case, add uppercased
			cipheredIndex := (plainIndex + shiftKey) % 26 //len(positionByLetter)
			candidate := letterByPosition[cipheredIndex]
			if !isUpperCase {
				candidate = strings.ToLower(candidate)
			}
			result += candidate
			fmt.Println("letter: ", string(ch), ", key: ", shiftKey, ", plainIndex: ", plainIndex, ", isUpperCase:", isUpperCase, ", cipheredIndex: ", cipheredIndex, " candidate: ", candidate)

		} else {
			result += string(ch)
		}
	}
	return result
}
func isLetter(c rune) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}

func isUpperCase(c rune) bool {
	return 'A' <= c && c <= 'Z'
}

func mapPositionByLetter() map[string]int {
	m := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
		"E": 5,
		"F": 6,
		"G": 7,
		"H": 8,
		"I": 9,
		"J": 10,
		"K": 11,
		"L": 12,
		"M": 13,
		"N": 14,
		"O": 15,
		"P": 16,
		"Q": 17,
		"R": 18,
		"S": 19,
		"T": 20,
		"U": 21,
		"V": 22,
		"W": 23,
		"X": 24,
		"Y": 25,
		"Z": 26,
	}
	return m
}

func mapLetterByPosition() map[int]string {
	m := map[int]string{
		0:  "Z",
		1:  "A",
		2:  "B",
		3:  "C",
		4:  "D",
		5:  "E",
		6:  "F",
		7:  "G",
		8:  "H",
		9:  "I",
		10: "J",
		11: "K",
		12: "L",
		13: "M",
		14: "N",
		15: "O",
		16: "P",
		17: "Q",
		18: "R",
		19: "S",
		20: "T",
		21: "U",
		22: "V",
		23: "W",
		24: "X",
		25: "Y",
		26: "Z",
	}
	return m
}
