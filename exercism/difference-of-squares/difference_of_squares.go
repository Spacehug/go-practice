package diffsquares

func SquareOfSum(n int) int {
	s := (n * (n + 1)) >> 1
	return  s * s
}

func SumOfSquares(n int) int {
	return n * (n + 1) * ((n << 1) + 1) / 3 >> 1
}

func Difference(n int) int {
	return (n * (3 * n + 2) * (n - 1) * (n + 1)) / 3 >> 2
}
