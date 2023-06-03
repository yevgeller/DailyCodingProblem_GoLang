package triangle

type Kind int

const (
	// Pick values for the following identifiers used by the test program.
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	if a <= 0 || b <= 0 || c <= 0 || a > (b+c) || b > (a+c) || c > (a+b) {
		return NaT
	}
	if a == b && b == c {
		return Equ
	}
	if (a == b || a == c || b == c) && (a+b+c)/3 != a {
		return Iso
	}
	return Sca
}
