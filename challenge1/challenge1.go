package challenge1

import "math/big"

// Run calculates the nth value of the Fibonacci Sequence
func Run(n int) uint64 {
	// Initialize two big ints with the first two numbers in the sequence.
	a := big.NewInt(0)
	b := big.NewInt(1)

	// Loop while a is smaller than 1e100.
	for i := 0; i < n; i++ {
		// Compute the next Fibonacci number, storing it in a.
		a.Add(a, b)
		// Swap a and b so that b is the next number in the sequence.
		a, b = b, a
	}

	return a.Uint64()
}
