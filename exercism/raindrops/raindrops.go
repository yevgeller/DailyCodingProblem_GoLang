package raindrops

func Convert(number int) string {
	panic("Please implement the Convert function")
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
