package diffsquares

// SquareOfSum returns the square of the sum of the first N numbers
func SquareOfSum(n int) int {
	sum := n * (n+1) / 2
	return sum * sum
}

// SumOfSquares returns the sum of the squares of the first N numbers
func SumOfSquares(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += (i*i)
	}
	return sum
}

// Difference returns the difference between the two methods
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}