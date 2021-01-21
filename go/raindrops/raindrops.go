package raindrops

import (
	"strconv"
)

// Convert converts a number into a string
func Convert(n int) string {
	factors := Factors(n)
	s := strconv.Itoa(n)
	for _, f := range factors {
		if f == 3 {
			s = "Pling"
		} else if f == 5 {
			s = "Plang"
		} else if f == 7 {
			s = "Plong"
		} else if {
			s += ""
		}
	}
	return s
}

// Factors returns the factors ofa number
func Factors (n int) []int {
	factors := []int{}
	for i := 1; i <= n; i++ {
		if n % i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}