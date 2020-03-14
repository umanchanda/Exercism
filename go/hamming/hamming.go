// Package hamming contains a single function Distance
package hamming

import (
	"errors"
)

// Distance returns the Hamming distance between two strings
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Error")
	}
	score := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			score++
		}
	}
	return score, nil
}
