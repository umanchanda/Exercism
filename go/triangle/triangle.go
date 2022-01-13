// Package triangle should have a package comment that summarizes what it's about.
package triangle

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind string

const (
	// Pick values for the following identifiers used by the test program.
	NaT = "NaT" // not a triangle
	Equ = "Equ" // equilateral
	Iso = "Iso" // isosceles
	Sca = "Sca" // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	// Write some code here to pass the test suite.
	var k Kind
	if (a+b < c) || (b+c < a) || (a+c < b) || a == 0 || b == 0 || c == 0 {
		k = NaT
	} else if a == b && b == c && a == c {
		k = Equ
	} else if a != b && b != c && a != c {
		k = Sca
	} else if (a == b && b != c) || (b == c && c != a) || (a == c && c != b) {
		k = Iso
	} else {
		k = NaT
	}

	return k
}
