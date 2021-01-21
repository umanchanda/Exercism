package grains

import (
	"errors"
)

// Square returns the number of grains on a given square
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("error")
	}
	return uint64(1 << (n-1)), nil
}

// Total returns the total number of grains on the chessboard
func Total() uint64 {
	return uint64((1 << 64) - 1)
}