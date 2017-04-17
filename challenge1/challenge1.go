package challenge1

// Run calculates the nth value of the Fibonacci Sequence
func Run(n int) uint64 {
	n0 := uint64(0)
	n1 := uint64(1)

	if n == 0 {
		n1 = n0
	}

	for i := 0; i < n-1; i++ {
		n1 = n0 + n1
		n0 = n1 - n0
	}

	return n1
}
