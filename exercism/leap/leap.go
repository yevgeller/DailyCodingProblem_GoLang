// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package leap should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package leap

// IsLeapYear should have a comment documenting it.
func IsLeapYear(year int) bool {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	return year % 4 == 0 && year % 100 != 0 || year % 100 == 0 && year % 400 == 0
	
	//panic("Please implement the IsLeapYear function")


}

/*
on every year that is evenly divisible by 4
  except every year that is evenly divisible by 100
    unless the year is also evenly divisible by 400
	*/