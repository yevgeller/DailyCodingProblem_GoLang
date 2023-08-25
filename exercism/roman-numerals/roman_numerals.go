package romannumerals

import "errors"

func ToRomanNumeral(input int) (string, error) {
	if input < 1 {
		return "", errors.New("out of range")
	}
	if input > 3999 {
		return "", errors.New("out of range")
	}
	return numberToString(input), nil
}

func numberToString(input int) string {
	if input >= 1000 {
		return "M" + numberToString(input-1000)
	}
	if input >= 900 {
		return "CM" + numberToString(input-900)
	}
	if input > 500 {
		return "D" + numberToString(input-500)
	}

	if input >= 400 {
		return "CD" + numberToString(input-400)
	}

	if input >= 100 {
		return "C" + numberToString(input-100)
	}

	if input >= 90 {
		return "XC" + numberToString(input-90)
	}

	if input >= 50 {
		return "L" + numberToString(input-50)
	}

	if input >= 40 {
		return "XL" + numberToString(input-40)
	}

	if input >= 10 {
		return "X" + numberToString(input-10)
	}

	if input >= 9 {
		return "IX" + numberToString(input-9)
	}

	if input >= 5 {
		return "V" + numberToString(input-5)
	}

	if input >= 4 {
		return "IV" + numberToString(input-4)
	}

	if input >= 1 {
		return "I" + numberToString(input-1)
	}

	return ""
}
