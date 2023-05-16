// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	isQuestion := strings.HasSuffix("?")
	isAllUpper := unicode.IsUpper(remark)

	if isQuestion && !isAllUpper {
		return "Sure."
	}

	if !isQuestion && isAllUpper {
		return "Whoa, chill out!"
	}

	if isQuestion && isAllUpper {
		return "Calm down, I know what I'm doing!"
	}

	if len(remark) == 0 {
		return "Fine. Be that way!"
	}

	return "Whatever."
}
