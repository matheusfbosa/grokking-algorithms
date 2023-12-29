package chapter3

// Factorial calculates the factorial of a given non-negative integer.
func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

// FactorialTailRecursion calculates the factorial of a given non-negative integer using a recursive helper function.
func FactorialTailRecursion(n int) int {
	return fact(n, 1)
}

// fact is a helper function for FactorialTailRecursion.
// It uses recursion to compute the factorial with an accumulator for optimization.
// Note: Go does not perform true tail call optimization.
func fact(n, acc int) int {
	if n <= 1 {
		return acc
	}
	return fact(n-1, acc*n)
}
