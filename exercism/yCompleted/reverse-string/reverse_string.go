package reverse

func Reverse(input string) string {
	result := make([]byte, len(input))
	left, right := 0, len(input)
	for current := range input {
		right -= current - left
		copy(result[right:], input[left:current])

		left = current

	}
	copy(result[0:], input[left:])
	return string(result)

}
