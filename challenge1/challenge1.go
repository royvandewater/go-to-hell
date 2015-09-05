package challenge1

// Run calculates the nth value of the Fibonacci Sequence
func Run(n int) uint64 {
	sequence := []uint64{0, 1}

	if n < 2 {
		return sequence[n]
	}

	for i := 2; i <= n; i++ {
		lastValue := sequence[len(sequence)-1]
		secondToLastValue := sequence[len(sequence)-2]

		sequence = append(sequence, lastValue+secondToLastValue)
	}

	return sequence[len(sequence)-1]
}
