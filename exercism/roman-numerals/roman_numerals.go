package romannumerals

import "errors"

func ToRomanNumeral(input int) (string, error) {
	//panic("Please implement the ToRomanNumeral function")
if input < 1 {
	return "", errors.New("input cannot be less than 1")
}

}
