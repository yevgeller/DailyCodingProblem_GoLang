// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	var ret []string
	ln := len(rhyme)
	if ln == 0 {
		return ret
	}
	a := 0

	for a < (ln - 1) {
		//if i < len(rhyme) {
		ret = append(ret, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[a], rhyme[a+1]))
		//}
		a++
	}

	ret = append(ret, fmt.Sprintf("And all for the want of a %s", rhyme[0])) //ADD A PERIOD HERE
	return ret
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	//panic("Please implement the Proverb function")
}
