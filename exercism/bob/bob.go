// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	if len(strings.TrimSpace(remark)) == 0 {
		return "Fine. Be that way!"
	}

	isQuestion := strings.HasSuffix(strings.TrimSpace(remark), "?")
	isAllUpper := strings.ToUpper(remark) == remark && strings.ToLower(remark) != remark

	if isQuestion && !isAllUpper {
		return "Sure."
	}

	if !isQuestion && isAllUpper {
		return "Whoa, chill out!"
	}

	if isQuestion && isAllUpper {
		return "Calm down, I know what I'm doing!"
	}

	return "Whatever."
}
