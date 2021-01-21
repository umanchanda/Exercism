package perfect

import "errors"

// Classification classifies a number
type Classification string

const (
	// ClassificationDeficient sum < n
	ClassificationDeficient Classification = "deficient"
	// ClassificationPerfect sum == n
	ClassificationPerfect Classification = "perfect"
	// ClassificationAbundant sum > n
	ClassificationAbundant Classification = "abundunt"
)

// ErrOnlyPositive is called when n < 0
var ErrOnlyPositive = errors.New("Not a positive number")

// Classify determines what classification a number is
func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return "", ErrOnlyPositive
	}
	sum := AliquotSum(int(n))
	if sum == n {
		return ClassificationPerfect, nil
	} else if sum > n {
		return ClassificationAbundant, nil
	}
	return ClassificationDeficient, nil
}

// AliquotSum returns the sum of factors
func AliquotSum(n int) int64 {
	sum := 0
	for i := 1; i <= n/2; i++ {
		if n % i == 0 {
			sum += i
		}
	}
	return int64(sum)
}