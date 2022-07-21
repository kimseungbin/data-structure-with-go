package ch01

func Factorial(i int) int {
	// Terminal condition
	if i <= 1 {
		return 1
	}

	// Body, recursive expansion
	return i * Factorial(i-1)
}
