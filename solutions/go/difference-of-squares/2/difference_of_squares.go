package diffsquares

func SquareOfSum(n int) int {
	k := (n * (n + 1)) / 2
	return k * k
}

func SumOfSquares(n int) int {
	k := n * (n + 1) * (2*n + 1)
	return k / 6
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
