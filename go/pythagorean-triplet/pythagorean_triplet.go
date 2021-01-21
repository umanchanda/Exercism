package pythagorean

// Triplet represents a pythagorean triplet
type Triplet [3]int

// Range returns a list of all Pythagorean triplets with sides in the range min to max inclusive
func Range(min, max int) []Triplet {
	t := []Triplet{}
	for a := min; a <= max; a++ {
		for b := min; b <= max; b++ {
			for c := min; c <= max; c++ {
				if a*a + b*b == c*c && a < b && b < c {
					t = append(t, Triplet{a,b,c})
				}
			}
		}
	}
	return t
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c=p 
func Sum(p int) []Triplet {
	t := []Triplet{}
	for _,x := range Range(1, p) {
		if x[0] + x[1] + x[2] == p {
			t = append(t, Triplet{x[0],x[1],x[2]})
		}
	}
	return t
}