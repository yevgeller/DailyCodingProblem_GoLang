package raindrops

import "fmt"

func Convert(number int) string {
    
	var ret = ""
	if number%3 == 0 {
		ret += "Pling"
	}
	if number%5 == 0 {
		ret += "Plang"
	}
	if number%7 == 0 {
		ret += "Plong"
	}

	if len(ret) == 0 {
		return fmt.Sprintf("%d", number)
	}

	return ret
}
