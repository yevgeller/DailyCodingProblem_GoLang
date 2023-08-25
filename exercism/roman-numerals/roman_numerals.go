package romannumerals

import "errors"

func ToRomanNumeral(input int) (string, error) {
	//panic("Please implement the ToRomanNumeral function")
	if input < 1 {
		return "", errors.New("input cannot be less than 1")
	}

	return numberToString(input), nil
}

func numberToString(input int) string {
	if input > 1000 {
		return "M" + numberToString(input-1000)
	}
}
