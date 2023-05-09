package reverse

func Reverse(input string) string {
	res := make([]byte, len(input))
	prevPos, resPos := 0, len(input)
	for pos := range input {
		resPos -= pos - prevPos
		copy(res[resPos:], input[prevPos:pos])
		prevPos = pos
	}
	copy(res[0:], input[prevPos:])
	return string(res)

	// fmt.Println("return: ", ret)
	// return ret
}
