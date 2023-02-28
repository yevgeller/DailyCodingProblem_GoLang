package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	return stringBanner(numStarsPerLine) + "\n" + welcomeMsg + "\n" + stringBanner(numStarsPerLine)
}

func stringBanner(numStarsPerLine int) string {
	return strings.Repeat("*", numStarsPerLine)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	panic("Please implement the CleanupMessage() function")
}
