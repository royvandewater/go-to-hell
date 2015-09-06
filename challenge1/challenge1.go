package challenge1

var sequence []uint64

// Run calculates the nth value of the Fibonacci Sequence
func Run(n int) uint64 {
	sequence := generate()
	return sequence[n]
}

func generate() []uint64 {
	if sequence != nil {
		return sequence
	}
	sequence = []uint64{0, 1}

	for i := 2; i <= 100; i++ {
		lastValue := sequence[len(sequence)-1]
		secondToLastValue := sequence[len(sequence)-2]

		sequence = append(sequence, lastValue+secondToLastValue)
	}

	return sequence
}
