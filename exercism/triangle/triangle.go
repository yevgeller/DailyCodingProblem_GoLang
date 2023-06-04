package triangle

type Kind int

const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

func KindFromSides(a, b, c float64) Kind {
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
