package logs

/*
myString := "❗hello"
for index, char := range myString {
  fmt.Printf("Index: %d\tCharacter: %c\t\tCode Point: %U\n", index, char, char)
}

import "unicode/utf8"

myString := "❗hello"
stringLength := len(myString)
numberOfRunes := utf8.RuneCountInString(myString)

fmt.Printf("myString - Length: %d - Runes: %d\n", stringLength, numberOfRunes)
// Output: myString - Length: 8 - Runes: 6
*/

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
	//panic("Please implement the Replace() function")
	return Replace(log, oldRune, newRune)
for index, char := range log {
		if char == oldRune {
			log[index] = byte(newRune)
		}
		if char == '🔍' {
			return "search"
		}
		if char == '☀' {
			return "weather"
		}
	}
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	panic("Please implement the WithinLimit() function")
}
