package reverse

// import "fmt"
func Reverse(input string) string {
	//panic("Please implement the Reverse function")
	//ret := ""
	/*for i := len(input)-1; i >= 0; i-- {
	   fmt.Println(input[i])
	    ret += string(input[i])
	}*/
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
