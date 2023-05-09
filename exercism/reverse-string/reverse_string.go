package reverse

import "fmt"

func Reverse(input string) string {
	res := make([]byte, len(input))
	prevPos, resPos := 0, len(input)
	for pos := range input {
		resPos -= pos - prevPos
		copy(res[resPos:], input[prevPos:pos])
		prevPos = pos
		fmt.Println("pos: ", pos, " resPos: ", resPos, ", res[resPos:]", res[resPos:], ", input[prevPos:pos]", input[prevPos:pos],
			", prevPos: ", prevPos)
	}
	copy(res[0:], input[prevPos:])
	return string(res)

	// fmt.Println("return: ", ret)
	// return ret
}
