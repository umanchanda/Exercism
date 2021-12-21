package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct{}

func (e SillyNephewError) Error(i int) error {
	return fmt.Errorf("silly nephew, there cannot be %d cows", i)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	amt, err := weightFodder.FodderAmount()
	if amt < 0 {
		return 0.0, errors.New("Negative fodder")
	}

	if err != nil {
		if err == ErrScaleMalfunction && amt > 0 {
			amt *= 2
		} else {
			return 0.0, err
		}
	}

	if cows == 0 {
		return 0.0, errors.New("Division by zero")
	}

	if cows < 0 {
		e := SillyNephewError{}
		return 0.0, e.Error(cows)
	}

	return amt / float64(cows), nil

	// panic("Please implement DivideFood")
}
