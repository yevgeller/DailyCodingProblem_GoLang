package logs

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, char := range log {
		if char == '❗' {
			return "recommendation"
		}
		if char == '🔍' {
			return "search"
		}
		if char == '☀' {
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	panic("Please implement the Replace() function")
	runes := []rune(log)

	for index, char := range runes {
		if char == oldRune {
			runes[index] = newRune
		}
	}
	return string(runes)

}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	panic("Please implement the WithinLimit() function")
	runes := []rune(log)
	return len(runes) <= limit

}
