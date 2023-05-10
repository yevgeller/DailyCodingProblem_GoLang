package reverse

import "fmt"

func Reverse(input string) string {
	fmt.Println("assignment: ", input)
	result := make([]byte, len(input))
	left, right := 0, len(input)
	for current := range input {
		fmt.Println("left: ", left, ", current: ", current, " right: ", right, ", result[right:]", result[right:], ", input[left:current]: ", input[left:current])
		right -= current - left
		copy(result[right:], input[left:current])

		left = current

	}
	fmt.Println("left: ", left, ", current: N/A", " right: ", right, ", result[right:]", result[right:], ", input[left:len(input)-1]: ", input[left:len(input)-1])
	fmt.Println("result[0:]", result[0:])
	copy(result[0:], input[left:])
	fmt.Println()
	return string(result)

}
